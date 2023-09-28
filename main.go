/*
*

	@author: illumwang
	@since:  2023/9/23
	@desc:  //TODO

*
*/
package main

import (
	"Gedis/redisClient"
	"fmt"
)

func main() {
	//第一步、链接服务端
	client, err := redisClient.Client("localhost:6379", "", 0)
	if err != nil {
		panic(err)
	}
	//Set方法
	result, err := client.Set("name", "pdudo")
	if err != nil {
		fmt.Println("set 失败", err)
		return
	}
	fmt.Println("set 结果:", result)
	//get方法
	response, err := client.Get("name")
	fmt.Println("get 结果：", response)
}
