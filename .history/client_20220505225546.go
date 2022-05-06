package main

import (
	"fmt"
	"io"
	"net"
)
type Client struct{
	ServerIp string
	serverPort int
	conn net.Conn
}

func NewClient(serverIp string,serverPort int){
	//创建客户端对象
  client := &Client{
		ServerIp: serverIp,
		serverPort: serverPort,
	}
//
	return client
}


