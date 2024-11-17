package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8885")
	defer conn.Close()
	reader := bufio.NewReader(conn)
	shu, _ := reader.ReadString('\n')
	shu = strings.TrimSpace(shu)
	fmt.Println("服务器发送的随机数是：", shu)
	input := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入相同的随机数：")
		text, _ := input.ReadString('\n')
		text = strings.TrimSpace(text)
		_, _ = conn.Write([]byte(text + "\n"))
		break
	}
}
