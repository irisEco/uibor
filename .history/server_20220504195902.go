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
func (s *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + ":" + user.Name + ": " + msg
	s.Message <- sendMsg
}

//监听用户消息
func (s *Server) ListenMessage() {
	for {
		msg := <-s.Message

		//将msg发送给全部在线的User
		s.mapLock.Lock()
		for _, cli := range s.OnlineMap {
			cli.C <- msg
		}
		s.mapLock.Unlock()
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
func (s *Server) Hander(conn net.Conn) {
	//用户上线
	user := NewUser(conn,s)
	user.Online()

	//监听用户是否活跃的channel
	isLive := make(chan bool)

	//接受客户端发送的消息
	go func() {
		buf := make([]byte, 4098)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Printf("Conn read err: ", err)
				return
			}
			msg := string(buf[:n-1])

			user.Domessage(msg)	

			// 任何消息,代表当前用户活跃
			isLive <- true	
		}
	}()

	 // 当前handler阻塞
    for
}



//启动服务器的接口
func (s *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s: %d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}
	//close listen socket
	defer listener.Close()

	//启动监听massage的goroutine
	go s.ListenMessage()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err: ", err)
			continue
		}
		//do hander
		go s.Hander(conn)

	}
}
