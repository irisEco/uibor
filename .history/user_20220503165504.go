package main

import "net"
type User struct{
	Name string
	Addr string
	C chan string
	conn net.Conn
}

//创建一个用户的Api
func NewUser(conn net.Conn)*User{
	userAddr := 
}