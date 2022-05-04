package main

import (
	"fmt"
	"net"
)
type Server struct{
	Ip string
	Port int
	OnlineMap map[string]string 
}

// 创建一个server的接口
func NewServer(ip string, port int, OnlineMap map[string]string) *Server{
	server := &Server{
		Ip: ip,
		Port: port,
		OnlineMap: OnlineMap,
	}
	return server
}
func (this *Server) Hander(com net.Conn){
		//当前链接业务
		OnlineMap := make(map[string]string)
		var OnlineMap = make(map[this.Ip]this.Ip)
		fmt.Printf("%s:%d 链接创建成功",this.Ip,this.Port,this.OnlineMap)
}

//启动服务器的接口
func (this *Server)Start(){
	//socket listen
	listener,err := net.Listen("tcp", fmt.Sprintf("%s: %d",this.Ip,this.Port))
	if err != nil{
		fmt.Println("net.Listen err: ",err)
		return
	}
	//close listen socket
	defer listener.Close()

	for{
		//accept
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener accept err: ",err)
		continue
	}
		//do hander
	go this.Hander(conn)
}
}