package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	ServerIp   string
	serverPort int
	conn       net.Conn
	flag       int //当前客户端的选择模式
}

func NewClient(serverIp string, serverPort int) *Client {
	//创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		serverPort: serverPort,
		flag:       99,
	}
	// 连接server
	addr := fmt.Sprintf("%s:%d", serverIp, serverPort)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("net.Dial err: ", err)
		return nil
	}
	client.conn = conn
	return client
}

func (c *Client) Navigation() bool {
	var flag int
	fmt.Println("请选择以下数字进入相应的功能")
	fmt.Println("1.公聊模式")
	fmt.Println("2.修改用户名")
	fmt.Println("3.私聊模式")
	fmt.Println("0.退出聊天室")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		c.flag = flag
		return true
	} else {
		fmt.Println("---------请输入合法范围内的数字------------")
		return false
	}
}

func (c *Client)DealResponse(){
	io.Co
	n,err := c.conn.Write([]byte)
	if err != nil{
		fmt.Print("Dial Write err: ",err)
	}

}

var (
	serverIp   string
	serverPort int
)

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址(默认是127.0.0.1)")
	flag.IntVar(&serverPort, "prot", 3000, "设置服务器端口(默认是3000)")
}

func (c *Client) Run() {

	for c.flag != 0 {
		for !c.Navigation() {
		}
		var msg string
		//根据不同的模式处理不同的业务
		switch c.flag {
		case 1:
			//进入公聊模式
			fmt.Println("进入 公聊 模式")
			fmt.Scanln(&msg)
			c.conn.Write([]byte(msg))
			// break
		case 2:
			//更新用户名
			fmt.Println("进入 更新用户名 ")
			// break
		case 3:
			//私聊模式
			fmt.Println("进入 私聊 模式")
			// break
		}
	}
	fmt.Println("退出!")
}

func main() {

	//命令行解析
	flag.Parse()
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>服务器连接失败")
	} else {
		fmt.Println(">>>>>服务器连接成功")
	}

	//启动客户端业务
	client.Run()
	// select {}
}
