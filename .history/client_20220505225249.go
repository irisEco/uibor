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
  client := &Client{
		ServerIp: serverIp,
		serverPort: s,
	}
}


