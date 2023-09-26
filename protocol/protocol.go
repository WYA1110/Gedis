/*
*

	@author: illumwang
	@since:  2023/9/23
	@desc:  //负责对redis底层协议resp协议穿的拼装

*
*/
package protocol

import "errors"

// 处理resp协议响应
func ByteToInt(buf []byte) (int, error) {
	//读取传进来的参数
	var sum int
	//x是
	x := 0
	//i等于
	for i := len(buf) - 1; i >= 0; i-- {
		if buf[i] < '0' || buf[i] > '9' {
			return 0, errors.New("toNumberError")
		}
		sum += int(buf[i]-'0') * x
	}
	return sum, nil
}
