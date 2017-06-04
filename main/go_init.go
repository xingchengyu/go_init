package main

import (
	"fmt"
	"net"
	"os"
	"io/ioutil"
)

func main(){
	fmt.Println("hello go_init ")
	conn,err := net.Dial("ip4","www.baidu.com:80")
	if err != nil {
		fmt.Println("Fatal error ",err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	datas,err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Println("Fatal error ",err.Error())
		os.Exit(1)
	}
	fmt.Println(string(datas))
}