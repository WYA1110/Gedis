/*
*

	@author: illumwang
	@since:  2023/9/23
	@desc:  //负责对redis底层协议resp协议穿的拼装

*
*/
package protocol

// 常量定义RESP中固定的报文值（开头除了以下几种没别的可能了）
const (
	CRLF              = "\r\n"
	redisArray        = "*"
	redisNum          = ":"
	redisString       = "$"
	redisError        = "-"
	redisSampleString = "+"
)

// 接受传进来的resp数据
func Accept() {

}
