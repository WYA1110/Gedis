/*
*

	@author: illumwang
	@since:  2023/9/23
	@desc:  //负责完成对redis的网络连接，实现tcp通信，连接6379端口并从服务端获取信息

*
*/
package redisClient

import (
	"Gedis/protocol"
	"bufio"
	"errors"
	"fmt"
	"net"
	"strconv"
)

// 定义所有接受的数据是同一种类型
type RedisInfo struct {
	//存储redis服务端信息
	address string
	//存储连接
	Conn net.Conn
	//读取链接
	Reader  *bufio.Reader
	bufRead []string
}

// 常量定义RESP中固定的报文值（开头除了以下几种没别的可能了）
const (
	CRLF              = "\r\n"
	RedisArray        = '*'
	redisInt          = ':'
	RedisString       = '$'
	redisError        = '-'
	redisSampleString = '+'
)

func Client(address string, password string, index int) (client RedisInfo, err error) {
	//1. 链接redis
	fmt.Printf("connecting Redis server....\naddress:%s \npassword:%s \n index:%d \n", address, password, index)
	client.address = address
	client.Conn, err = net.Dial("tcp", client.address)
	if err != nil {
		return
	}
	client.Reader = bufio.NewReader(client.Conn)
	//2.redis设置
	if password != "" {
		//密码校验，不为空则err
	}
	if index != 0 {
		//索引校验，不为0则err
	}
	return
}

// 发送数据
func (client *RedisInfo) RespWriter(buf []byte) (err error) {
	writenLine := 0
	for writenLine < len(buf) {
		n, err := client.Conn.Write(buf[writenLine:])
		if err != nil {
			return err
		}
		writenLine += n
	}
	return nil
}

// 获取数据，返回获取的resp字符串,并且按照读取类型不同分派出去
func (client *RedisInfo) RespReader() (buf []string, err error) {
	//从输入流中读取数据
	newLine, _, _ := client.Reader.ReadLine()
	//通过读取数据的第一个字符来判断类型
	switch newLine[0] {
	//数组类型，读取到数组类型的时候，目的是转换阅读型更好的字符串，数组格式的字符串发送来格式为'*'
	case RedisArray:
		count, err := protocol.ByteToInt(newLine[1:])
		if err != nil {
			fmt.Println("数据初次转换错误")
			return nil, err
		}
		return client.redisArrayResp(count)

		//整数类型，返回值如：“:4\r\n”
	case redisInt:
		//去掉冒号和\r\n
		buf = append(buf, string(newLine[1:]))

		//多行字符串，"$6\r\nfoobar\r\n",即返回一个6字符的字符串foobar，然后返回string
	case RedisString:
		//先判断返回元素有多少个
		n, _ := protocol.ByteToInt(newLine[1:])
		//创建新的切片,有n个元素
		newSlice := make([]byte, n)
		//在client.reader中读取newSlice的内容，并转换为string
		client.Reader.Read(newSlice)
		buf = append(buf, string(newSlice))
		//读取结尾的回车换行符号，表示多行字符串结束
		crlf := make([]byte, 2)
		client.Reader.Read(crlf)
		return

		//错误信息，返回值"-ERR unknown command 'illumwang'\r\n",客户端显示“(error) ERR unknown command 'illumwang' ”
		//要将-替换为(error),因为错误返回值中不存在返回元素多少的问题，所以不需要byteToInt，直接返回错误就行
	case redisError:
		err = errors.New(string(newLine[1:]))

		//单行字符串,返回值"+OK\r\n"，去掉第一个字符+
	case redisSampleString:
		buf = append(buf, string(newLine[1:]))
	}
	return
}

// 数组格式数据处理，接收到元素个数为count，确定for循环次数为count
func (client *RedisInfo) redisArrayResp(count int) (buf []string, err error) {
	for i := 0; i < count; i++ {
		//从redisSampleStringResp获取resp转化后的字符串
		newBuf, _ := client.redisSampleResp()
		buf = append(buf, newBuf...)
	}
	return
}

// 获取数据操作
func (client *RedisInfo) redisSampleResp() (buf []string, err error) {
	//这个方法获取数据获取的是几乎所有数据，所以我们首先要去读出来
	//ReadLine源码注释为：The text returned from ReadLine does not include the line end ("\r\n" or "\n").
	//所以readline是从reader中读取不包含行尾crlf的数据
	newReadLine, _, _ := client.Reader.ReadLine()
	switch newReadLine[0] {
	//格式是整数
	case redisInt:
		buf = append(buf, string(newReadLine[1:]))

	//格式是字符串
	case RedisString:
		n, _ := protocol.ByteToInt(newReadLine[1:])
		newBuf := make([]byte, n)
		client.Reader.Read(newBuf)
		buf = append(buf, string(newBuf))
		crlf := make([]byte, 2)
		client.Reader.Read(crlf)
	}
	return
}
func (client *RedisInfo) Run() {

}
func (client *RedisInfo) auth() {

}

// 定义方法api
func (client *RedisInfo) Set(key string, value string) ([]string, error) {
	fmt.Printf("this is set method,key:%v,value:%v", key, value, "\n")
	//第一步，准备传入的数值
	var request []string
	//set是置数进入服务端，所以需要拼接字符串，set命令，key和value的内容
	request = append(request, "set")
	request = append(request, key)
	request = append(request, value)
	//第二步，转换为字节数组准备写入
	buf, err := ToResp(request)
	if err != nil {
		return nil, err
	}
	fmt.Println("等待发送的报文", request)
	//第三步，连接服务端写入
	client.RespWriter(buf)
	return client.RespReader()
}

// get方法是通过key找到value，将resp字符串转变为可读性高的string
// 127.0.0.1:6379> get site
// $6\r\
// ljheee\r\n
// ---
// "ljheee"
func (client *RedisInfo) Get(key string) (value []string, err error) {
	fmt.Printf("Get操作,key :%s", key)
	var response []string
	response = append(response, "get")
	response = append(response, key)
	buf, _ := ToResp(response)
	fmt.Println("待发送报文:", response)
	client.RespWriter(buf)
	return client.RespReader()
}

// 由普通字符串转变为resp字符串
func ToResp(writeString []string) (resp []byte, err error) {
	//第一步,判断字符长度
	arrayLen := len(writeString)
	if arrayLen == 0 {
		return
	}
	//字符串格式转化,拼接crlf等,到这一步，拼接的是set的方法的头部
	//SET key value #对应的resp通信协议串
	//*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n
	resp = append(resp, RedisArray)
	resp = append(resp, []byte(strconv.Itoa(arrayLen))...)
	resp = append(resp, []byte(CRLF)...)
	//第二部格式转化，要转化resp字符串的内容
	//v是一个临时变量，有几个元素打印几次
	for _, v := range writeString {
		//第一步绝对是string格式
		resp = append(resp, RedisString)
		vLen := len(v)
		//拼元素长度
		resp = append(resp, []byte(strconv.Itoa(vLen))...)
		resp = append(resp, []byte(CRLF)...)
		//拼内容
		resp = append(resp, []byte(v)...)
		resp = append(resp, []byte(CRLF)...)
	}
	return
}
