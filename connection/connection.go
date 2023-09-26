/*
*

	@author: illumwang
	@since:  2023/9/23
	@desc:  //负责完成对redis的网络连接，实现tcp通信，连接6379端口并从服务端获取信息

*
*/
package connection

import (
	"bufio"
	"fmt"
	"net"
)

// 定义所有接受的数据是同一种类型
type Message struct {
	//存储redis服务端信息
	Address string
	//存储连接
	Conn net.Conn
	//读取链接
	Reader  *bufio.Reader
	BufRead []string
}

func Client(address string, password string, port int) (client Message, err error) {
	//1. 链接redis
	fmt.Printf("connecting Redis server....\naddress:%s \npassword:%s \n port:%d \n", address, password, port)
	client.Address = address
	client.Conn, err = net.Dial("tcp", client.Address)
	if err != nil {
		fmt.Printf("failed to create connection ", err)
	}
	client.Reader = bufio.NewReader(client.Conn)
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
func (client *Message) respWriter(buf []byte) (err error) {
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
func (client *Message) respReader() (buf []string, err error) {
	//从输入流中读取数据
	newLine, _, _ := client.Reader.ReadLine()
	//通过读取数据的第一个字符来判断类型
	switch newLine[0] {
	//数组类型，读取到数组类型的时候，目的是转换阅读型更好的字符串，数组格式的字符串发送来格式为
	//
	//case redisArray:
	//	count,err:=byteToInt(newLine[1:])
	//
	//case redisNum:
	//case redisString:
	//case redisError:
	//case redisSampleString:

	}
	return
}

// 数组格式数据接受，接受时为int数据，转换为字符串可阅读
func (client *Message) resqArray(count int) (buf []string, err error) {

	return
}
