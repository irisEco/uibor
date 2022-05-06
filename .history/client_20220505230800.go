package main

import (
	"fmt"
	"io"
	"net"
)

type Client struct {
	ServerIp   string
	serverPort int
	conn       net.Conn
}

func NewClient(serverIp string, serverPort int) {
	//创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		serverPort: serverPort,
	}
	// 连接server
	conn, err := net.Dial("tcp", fmt.Sprint("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial err: ", err)
		return
	}
	client.conn = conn
	return client
}

func main(){
	client := NewClient("127.0.0.1",3000)
	
}
