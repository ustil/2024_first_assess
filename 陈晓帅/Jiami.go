package main

import (
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8888")
	defer conn.Close()

	conn.Write([]byte("Jia"))

	for {

		var text, key string
		JiamiText := make([]byte, 1024)

		fmt.Println("请输入要加密的文本:")
		fmt.Scanln(&text)
		fmt.Println("请输入密钥:")
		fmt.Scanln(&key)

		conn.Write([]byte(text))
		conn.Write([]byte(key))

		E, _ := conn.Read(JiamiText)

		fmt.Printf("加密后的文本：%s\n", string(JiamiText[:E]))

	}
}
