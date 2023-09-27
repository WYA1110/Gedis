/*
*

	@author: illumwang
	@since:  2023/9/28
	@desc:  //TODO

*
*/
package connection

import (
	"bufio"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func Test_respReader(t *testing.T) {
	//创建测试RedisInfo实例
	client := RedisInfo{
		reader: nil,
		//Conn: Client("127.0.0.1:6379","",0),
	}
	// 测试数组类型的响应
	arrayResponse := "*3\r\n$5\r\nhello\r\n$5\r\nworld\r\n$5\r\nfoo\r\n"
	client.reader = bufio.NewReader(strings.NewReader(arrayResponse))
	buf, err := client.respReader()
	fmt.Println(buf)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := []string{"hello", "world", "foo"}
	if !reflect.DeepEqual(buf, expected) {
		t.Errorf("Expected %v, got %v", expected, buf)
	}
	//测试整数

}
