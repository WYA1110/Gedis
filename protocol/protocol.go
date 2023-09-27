/*
*

	@author: illumwang
	@since:  2023/9/23
	@desc:  //负责对redis底层协议resp协议穿的拼装

*
*/
package protocol

import "errors"

// 处理resp协议响应,计算传入的数组元素的个数
func ByteToInt(buf []byte) (int, error) {
	//读取传进来的参数
	var sum int
	//x是权重
	x := 1
	//i为去掉第一、二位，即“*”和“$”之后的切片长度(第一位应该是个数字
	for i := len(buf) - 1; i >= 0; i-- {
		//在数字0-9之外的符号返回0报错无法解释
		if buf[i] < '0' || buf[i] > '9' {
			return 0, errors.New("OverNumberError")
		}
		sum += int(buf[i]-'0') * x
		x = x * 10
	}
	return sum, nil
}
