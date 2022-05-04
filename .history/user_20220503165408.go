package main

import "net"
type User struct{
	Name string
	Addr string
	C chan string
	conn net.Conn
}