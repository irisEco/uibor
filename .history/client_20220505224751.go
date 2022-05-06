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


