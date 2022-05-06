package main

import (
	"fmt"
	"net"
	"flag"
)

type Client struct {
	ServerIp   string
	serverPort int
	conn       net.Conn
}

func NewClient(serverIp string, serverPort int) *Client {
	fmt.Println(serverIp,serverPort)
	//创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		serverPort: serverPort,
	}
	// 连接server
	addr := fmt.Sprintf("%s:%d", serverIp, serverPort)
		fmt.Println(addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("net.Dial err: ", err)
		return nil
	}
	client.conn = conn
	return client
}

func init(){
	flag.StringVar(&serverIp,"ip","127.0.0.1","设置服务器IP地址(默认是127.0.0.1)")
  flag.IntVar(&ServerPort,"prot",3000,"设置服务器端口(默认是3000)")

	//命令行解析
	flag.Parse()
}

func main() {
	client := NewClient("127.0.0.1",3000)
	if client == nil {
		fmt.Println(">>>>>服务器连接失败")
	}else{
		
	}
	fmt.Println(">>>>>服务器连接成功")
	//启动客户端业务
	select {}
}
