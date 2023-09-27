/*
*

	@author: illumwang
	@since:  2023/9/23
	@desc:  //TODO

*
*/
package main

import (
	"Gedis/connection"
	"fmt"
)

func main() {
	//开启client端口
	address := "127.0.0.1:6379"
	result, err := connection.Client(address, "", 0)
	if err != nil {
		fmt.Println("Failed to connect to Redis:", err)
		return
	}
	defer result.Conn.Close()
	fmt.Println("success")
}
