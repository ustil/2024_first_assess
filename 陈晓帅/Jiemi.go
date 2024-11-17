package main

import (
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8888")
	defer conn.Close()

	conn.Write([]byte("Jie"))

	for {

		var JiamiText, key string
		text := make([]byte, 1000)

		fmt.Println("请输入要解密的文本:")
		fmt.Scanln(&JiamiText)
		fmt.Println("请输入密钥:")
		fmt.Scanln(&key)

		conn.Write([]byte(JiamiText))
		conn.Write([]byte(key))

		D, _ := conn.Read(text)

		fmt.Printf("解密后的文本：%s\n", string(text[:D]))
	}
}
