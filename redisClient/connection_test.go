/*
*

	@author: illumwang
	@since:  2023/9/28
	@desc:  //TODO

*
*/
package redisClient

import (
	"bufio"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestRedisInfo_RespReader(t *testing.T) {
	//创建测试RedisInfo实例
	client := RedisInfo{
		Reader: nil,
		//Conn: Client("127.0.0.1:6379","",0),
	}
	// 测试数组类型的响应
	arrayResponse := "*3\r\n$5\r\nhello\r\n$5\r\nworld\r\n$3\r\nfoo\r\n"
	client.Reader = bufio.NewReader(strings.NewReader(arrayResponse))
	buf, err := client.RespReader()
	fmt.Println(buf)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := []string{"hello", "world", "foo"}
	if !reflect.DeepEqual(buf, expected) {
		t.Errorf("Expected %v, got %v", expected, buf)
	}
	//测试整型数字
	intResponse := ":2\r\n"
	client.Reader = bufio.NewReader(strings.NewReader(intResponse))
	intBuf, err := client.RespReader()
	fmt.Println(intBuf)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	intexpected := []string{"2"}
	if !reflect.DeepEqual(intexpected, intBuf) {
		t.Errorf("Expected %v, got %v", intexpected, buf)
	}
	//测试多行字符串
	stringResponse := "$6\r\nkkkkkk\r\n"
	client.Reader = bufio.NewReader(strings.NewReader(stringResponse))
	stringBuf, err := client.RespReader()
	fmt.Println(stringBuf)
	stringExpected := []string{"kkkkkk"}
	if !reflect.DeepEqual(stringBuf, stringExpected) {
		t.Errorf("No equal")
	}
}

// 客户端写测试用例
func TestRedisInfo_RespWriter(t *testing.T) {

}
