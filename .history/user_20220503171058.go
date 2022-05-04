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
	userAddr := conn.RemoteAddr().string()
	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C: make(chan string),
		conn: conn,
	}
	return user
}

//监听当前user channel 的方法,一旦有消息,就直接写回客户端
func sendMessage(* this.){

}