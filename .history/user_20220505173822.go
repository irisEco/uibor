package main

import (
	"fmt"
	"net"
	"strings"
)

// todo  终端不可输入中文, 会乱码
type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

//创建一个用户的Api
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	// 启动监听
	go user.ListenMessage()
	return user
}

//业务解耦
//用户上线功能
func (u *User) Online() {
	//将新用户加入已在线的用户中
	u.server.mapLock.Lock()
	// // user.Name = s.server.Ip 错误写法 测试ip都是一致的!!!!所以无法广播
	u.server.OnlineMap[u.Name] = u
	u.server.mapLock.Unlock()
	//广播用户上线
	u.Domessage("已上线")
}

//用户下线功能
func (u *User) Offline() {
	u.server.mapLock.Lock()
	delete(u.server.OnlineMap, u.Name)
	u.server.mapLock.Unlock()
	//广播用户已经下线
	u.Domessage("已下线")
}

//处理消息业务功能
func (u *User) Domessage(msg string) {
	fmt.Println(msg)
	if strings.Contains(msg, "rename|") { // 判断是否包含 rename | 关键字来修改用户名
		fmt.Println("进入修改用户名功能")
		u.Rename(msg)
	} else if strings.Contains(msg, "to|") {
		fmt.Println("进入私聊")
		index := strings.Split(msg, "|")[1]
		data := strings.Split(index, " ")[0]
		msg := strings.Split(index, " ")[1]
		fmt.Println("index: \n", index)
		fmt.Println("data: \n", data)
		fmt.Println("len: \n", len(data)
		fmt.Println("msg: \n", msg)
		// u.ShowList()
		// chatter, ok := u.server.OnlineMap[data]
		u.server.mapLock.Lock()
		for _, user := range u.server.OnlineMap {
			fmt.Printf("%s %T\n", user.Name, user.Name)
			fmt.Printf("%s %T\n", data, data)
			if user.Name == data {
				// chatter := user
				user.SendMsg(u.Name + ": " + msg)
			} else {
				u.SendMsg("没有此人!")
			}
		}
		u.server.mapLock.Unlock()
		// if ok {
		// 	chatter.SendMsg(u.Name + ": " + msg)
		// } else {
		// 	u.SendMsg("没有此人!")
		// }
		// u.server.mapLock.Unlock()
	} else if strings.Contains(msg, "list") {
		fmt.Println("进入查询在线用户功能")
		u.ShowList()
	} else {
		fmt.Println("进入广播功能")
		u.server.BroadCast(u, msg)
	}
}

//查询在线的所有用户
func (u *User) ShowList() {
	u.server.mapLock.Lock()
	for _, user := range u.server.OnlineMap {
		u.SendMsg(user.Name)
	}
	u.server.mapLock.Unlock()
}

//修改用户姓名
func (u *User) Rename(msg string) {
	name := strings.Split(msg, "|")[1]
	//判断是否存在这个用户名
	_, ok := u.server.OnlineMap[name]
	if ok {
		u.SendMsg("该用户名已经存在!")
	} else {
		u.server.mapLock.Lock()
		delete(u.server.OnlineMap, u.Name)
		u.Name = name
		u.server.OnlineMap[name] = u
		u.server.mapLock.Unlock()
		u.SendMsg("您已经更新用户名：" + u.Name + "\n")
	}

}

//监听当前user channel 的方法,一旦有消息,就直接写回对端客户端
func (u *User) ListenMessage() {
	for {
		msg := <-u.C
		u.conn.Write([]byte(msg + "\n"))
	}
}

//发送消息
func (u *User) SendMsg(msg string) {
	u.conn.Write([]byte(msg + "\n"))
}
