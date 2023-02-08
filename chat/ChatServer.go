package chat

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// Client 客户端连接
type Client struct {
	conn     net.Conn
	username string
}

var clients map[net.Conn]Client

type ServerConfig struct {
	ConnType string
	Host     string
	Port     string
}

func InitServer(config *ServerConfig) {
	// 启动聊天服务器
	fmt.Println("Starting chat server...")
	l, err := net.Listen(config.ConnType, config.Host+":"+config.Port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer l.Close()

	// 初始化客户端连接列表
	clients = make(map[net.Conn]Client)

	// 监听
	listen(l, config)

}

// 处理客户端请求
func handleRequest(conn net.Conn) {
	username := conn.RemoteAddr().String()
	fmt.Println("New client connected:", username)

	// 接收客户端消息
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed:", username)
			return
		}

		// 将消息发送给所有客户端
		if strings.TrimSpace(message) != "" {
			fmt.Println("Received message from", username, ":", string(message))
			broadcast(username + ": " + message)
		}
	}
}

// 监听并接受客户端连接
func listen(l net.Listener, config *ServerConfig) {
	fmt.Println("Waiting for connections on port " + config.Port + "...")
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			return
		}
		go handleRequest(conn)
	}
}

// 广播消息
func broadcast(message string) {

	for conn, _ := range clients {
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error broadcast: ", err.Error())
			return
		}
	}

	//for client := range clients {
	//	client.conn.Write([]byte(message))
	//}
}
