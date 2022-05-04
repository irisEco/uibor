package main

import (
	"fmt"
	"io"
	"math/rand"
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

		//将msg发送给全部在线的User
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

//随机字符串
//todo 可扩展为随机中文字符串
func GetRandomString(n int) string {
   str := "0123456789abcdefghijklmnopqrstuvwxyz"
   bytes := []byte(str)
   var result []byte
   for i := 0; i < n; i++ {
      result = append(result, bytes[rand.Intn(len(bytes))])
   }
   return string(result)
}


//处理客户端业务
func (this *Server) Hander(conn net.Conn) {
	//用户上线
	user := NewUser(conn)
	//将用户加入在线人数中
	this.mapLock.Lock()
	// user.Name = this.Ip 错误写法 测试ip都是一致的!!!!所以无法广播
	user.Name = GetRandomString(6)
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()
	this.BroadCast(user, "已上线!")
	//当前hander 阻塞
	select {}

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
			this.BroadCast(user, msg)
		}
	}()
}

//业务解耦
//用户上线功能
func (this *Server)Online(conn net.Conn){
	//判断有链接,便创建新用户
	user := NewUser(conn)
	//将新用户加入已在线的用户中
	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()
	//广播用户上线
	this.BroadCast()

}
//用户下线功能
func (this *Server)Offline(){

}
//处理消息功能
func (this *Server)Domessage(){

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

	//启动监听massage的goroutine
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
