package main

import (
	"fmt"
	"io"
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
	sendMsg := "[" + user.Addr + "]" + ":" + user.Name + ": " + msg
	this.Message <- sendMsg
}

//监听用户消息
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message

		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

func (this *Server) Hander(conn net.Conn) {
	//用户上线
	user := NewUser(conn)
	//将用户加入在线人数中
	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()
	this.BroadCast(this.OnlineMap[this.Ip], "已上线!")
	//当前hander
	select{}

	//接受客户端发送的消息
	go func() {
		buf := make([]byte, 4098)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				this.BroadCast(user, "已下线")
				return
			}
			if err != nil && err != io.EOF {
				fmt.Printf("Conn read err: ", err)
				return
			}
			msg := string(buf[:n-1])
			this.BroadCast(this.OnlineMap[this.Ip], msg)
		}

	}()
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
	go this.ListenMessage()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err: ", err)
			continue
		}
		//do hander
		go this.Hander(conn)

	}
}
