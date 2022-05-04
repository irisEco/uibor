package main

import "net"
type User struct{
	Name string
	Addr string
	C chan string
	conn net.Conn
	server *Server
}

//创建一个用户的Api
func NewUser(conn net.Conn,server *Server)*User{
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C: make(chan string),
		conn: conn,
		server: server,
	}
	// 启动监听
	go user.ListenMessage()
	return user
}

//业务解耦
//用户上线功能
func (u *User)Online(){
	//将新用户加入已在线的用户中
	u.server.mapLock.Lock()
	u.server.OnlineMap[u.Name] = u
	u.server.mapLock.Unlock()
	//广播用户上线
	u.server.BroadCast(u,"已上线!")
}
//用户下线功能
func (u *User)Offline(){
  u.server.mapLock.Lock()
	u.server.OnlineMap.delete
}
//处理消息功能
func (u *User)Domessage(){

}


//监听当前user channel 的方法,一旦有消息,就直接写回对端客户端
func (this *User)ListenMessage(){
   for {
		 msg := <-this.C
		 this.conn.Write([]byte(msg+ "\n"))
	 }
}