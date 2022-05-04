package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int
	//在线用户的列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex
	//消息广播的channel
	Message chan string
}

// 创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}
//广播用户上线消息
func (this *Server) BroadCast(user *User, msg string) {
  sendMsg := "["+user.Addr+"]" + ":" + user.Name+": " + msg
	this.Message <- sendMsg
}
//监听用户消息
func (this *Server)ListenMessage(){
	for{
		msg := <- this.Message

		this.mapLock.Lock()
		user.ListenMessage
		this.mapLock.Unlock()
	}
}

func (this *Server) Hander(conn net.Conn) {
	//用户上线
	user := NewUser(conn)
	//将用户加入在线人数中
	this.mapLock.Lock()
	this.OnlineMap[this.Ip] = user
	this.mapLock.Unlock()

	fmt.Printf("%s:%d 链接创建成功,当前用户为: %s", this.Ip, this.Port)
}

//启动服务器的接口
func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s: %d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}
	//close listen socket
	defer listener.Close()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err: ", err)
			continue
		}
		//do hander
		go this.Hander(conn)
		 this.BroadCast(this)
	}
}
