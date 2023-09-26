# Gedis——基于RESP协议的Redis客户端demo
### 一、客户端与服务端的通信
在开始项目之前，我们首先需要了解redis客户端与服务端之间是利用RESP协议完成通信，在RESP中，一些数据类型通过他的第一个字节进行判断
如：
- 单行回复：回复的第一个字节是“+”
- 错误信息：回复的第一个字节是“-”
- 整形数字：回复的第一个字节是“:”
- 多行字符串：回复的第一个字节是“$”
- 数组：回复的第一个字节是“*”
### 二、客户端的作用
作为客户端，要在接受服务端的信息之后用一种可读性更强的方式反馈给用户，所以，在RESP中
相应服务端的客户端返回值是除了“+”和CRLF以外的内容


#### 单行回复
```127.0.0.1:6379> set name illumwang
 +OK\r\n  # 服务端实际返回
 ---
 OK   # redis-cli 客户端显示
```


#### 错误信息
错误信息与单行回复很相似，不过是将“-”替换为错误信息本身
```
127.0.0.1:6379> illumwang
-ERR unknown command 'illumwang'\r\n  # 服务端实际返回, 下同
---
(error) ERR unknown command 'illumwang'  # redis-cli 客户端显示, 下同

127.0.0.1:6379> set name illumwang konomo
-ERR syntax error\r\n
---
(error) ERR syntax error
```

#### 整数
使用以 ":" 作为前缀，以CRLF作为结尾的字符串来表示整数，很多命令都会返回整数，如`INCR``LLEN``LPUSH`
等命令
```
127.0.0.1:6379> LPUSH info JintaoZhang MoeLove
:2\r\n  # 服务端实际返回
---
(integer) 2  # redis-cli 客户端显示
```

#### 数组
数组类型可以用于客户端向服务端发送消息，当某些命令将元素结合返回给客户端的时候，也是用数组作为回复类型
，它以"*"开头，随后是CRLF，最典型的是`LRRANGE`命令
```
127.0.0.1:6379> LPUSH info illumwang konomo.info
:2\r\n   # 服务端实际返回, 下同
---
(integer) 2  # redis-cli 客户端显示, 下同

127.0.0.1:6379> LRANGE info 0 -1
*2\r\n$12\r\nkonomo.info\r\n$8\r\n  illumwang\r\n
---
1) "moelove.info"
2) "illumwang"

127.0.0.1:6379> LPOP info
$12\r\nkonomo.info\r\n
---
"konomo.info"

127.0.0.1:6379> LPOP info
$8\r\nillumwang\r\n
---
"illumwang"

127.0.0.1:6379> LRANGE info 0 -1
*0\r\n
---
(empty list or set)
```
#### 多行字符串
用于返回长度最大为512MB的二进制字符串，以"$"开头，后跟实际要发送的字节数，随后是 CRLF，
然后是实际的字符串数据，最后以 CRLF 结束。如果一个要发送一个空字符串，则会编码为 "$0\r\n\r\n" 。
某些情况下，当要表示不存在的值时候，则以 "$-1\r\n" 返回，这被叫做空多行字符串，当客户端库接收到这个
响应的时候，同样应该返回一个空值（例如 nil）而不是一个空字符串
### 三、客户端开发思路
经过对RESF的学习，Redis客户端需要满足最基础的两个功能:
1. 通信：通信采用TCP来完成通信，客户端与服务端需保持双工连接
2. 传参：将程序中输入的字符串用RESP协议转化后发送给服务端，再将服务端返回的数据转化为可读性较好的数据
### 四、设计
我将客户端分为了三大模块：

1. 通信模块：connection类，负责链接redis服务端，从服务端接受resp数据，然后将接受的数据发送给数据处理模块
2. 数据处理模块：接受通信传来的resp数据转化为数字
3. api模块：定义调用服务端数据的api