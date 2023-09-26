/*
*

	@author: illumwang
	@since:  2023/9/23
	@desc:  //负责完成对redis的网络连接，实现tcp通信，连接6379端口并从服务端获取信息

*
*/
package connection

import (
	"Gedis/protocol"
	"bufio"
	"fmt"
	"net"
)

// 定义所有接受的数据是同一种类型
type RedisInfo struct {
	//存储redis服务端信息
	address string
	//存储连接
	conn net.Conn
	//读取链接
	reader  *bufio.Reader
	bufRead []string
}

// 常量定义RESP中固定的报文值（开头除了以下几种没别的可能了）
const (
	CRLF              = "\r\n"
	redisArray        = "*"
	redisNum          = ":"
	redisString       = "$"
	redisError        = "-"
	redisSampleString = "+"
)

func Client(address string, password string, port int) (client RedisInfo, err error) {
	//1. 链接redis
	fmt.Printf("connecting Redis server....\naddress:%s \npassword:%s \n port:%d \n", address, password, port)
	client.address = address
	client.conn, err = net.Dial("tcp", client.address)
	if err != nil {
		fmt.Printf("failed to create connection ", err)
	}
	client.reader = bufio.NewReader(client.conn)
	//2.redis设置
	if password != "" {
		//密码校验，不为空则err
	}
	if port != 0 {
		//索引校验，不为0则err
	}
	return
}

// 发送数据
func (client *RedisInfo) respWriter(buf []byte) (err error) {
	writenLine := 0
	for writenLine < len(buf) {
		n, err := client.conn.Write(buf[writenLine:])
		if err != nil {
			return err
		}
		writenLine += n
	}
	return nil
}

// 获取数据，返回获取的resp字符串,并且按照读取类型不同分派出去
func (client *RedisInfo) respReader() (buf []string, err error) {
	//从输入流中读取数据
	newLine, _, _ := client.reader.ReadLine()
	//通过读取数据的第一个字符来判断类型
	switch newLine[0] {
	//数组类型，读取到数组类型的时候，目的是转换阅读型更好的字符串，数组格式的字符串发送来格式为
	case '*':
		count, err := protocol.ByteToInt(newLine[1:])
		if err != nil {
			fmt.Println("数据初次转换错误")
			return nil, err
		}
		return client.redisArrayResp(count)
	case ':':

	case '$':
	case '-':
	case '+':

	}
	return
}

// 数组格式数据接受，接受时为int数据，转换为字符串可阅读
func (client *RedisInfo) redisArrayResp(count int) (buf []string, err error) {

	return
}

// 获取其他数据
func (client *RedisInfo) redisSampleStringResp() {

}
