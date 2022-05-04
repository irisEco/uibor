package main

import (
	"fmt"
	"net"
)
type Server struct{
	Ip string
	Port int
}

// 创建一个server的接口
func NewServer(ip string, port int) *Server{
	server := &Server{
		Ip: ip,
		Port: port,
	}
	return server
}

//启动服务器的接口
func (this *Server)start(){
	//socket listen
	listener,err := net.Listen("tcp", fmt.Sprintf("%s: %d",this.Ip,this.Port))
	if err != nil{
		fmt.Println("net.Listen err: ",err)
		return
	}
	//close listen socket
	defer listener.Close()
	conn,err := listener.Accept()
	if err != nil {
		fmt.Println("listener accept err: ",err)
		continue
	}
}