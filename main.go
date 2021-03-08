package main

import (
	"fmt"
	"net"
	"strconv"
)

func portscanner()  {
	ip:=""
	port:= 80

	fmt.Println("Enter target ip address:")
	fmt.Scanln(&ip)
	fmt.Println("Enter target port number:")
	fmt.Scanln(&port)


	new_port := strconv.Itoa(port)

	_,err := net.Dial("tcp", ip+":"+new_port)
	if err !=nil{
		fmt.Println("port is closed",port)
	} else{
		fmt.Println("port is open",port)
	}
}

func main()  {

	portscanner()
}