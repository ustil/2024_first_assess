package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

func jianyan(conn net.Conn) {
	defer conn.Close()
	rand.Seed(time.Now().UnixNano())
	shuchu := rand.Intn(10000) + 1
	fmt.Fprintln(conn, shuchu)
	fmt.Println("随机数已发送到客户端：", shuchu)
	reader := bufio.NewReader(conn)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	shuru, _ := strconv.Atoi(line)
	if shuru == shuchu {
		fmt.Println("验证通过。客户端发送了正确的数字。")
	} else {
		fmt.Println("验证失败。客户端发送了错误的数字。")
	}
}
func main() {
	listener, _ := net.Listen("tcp", "127.0.0.1:8885")
	defer listener.Close()
	for {
		conn, _ := listener.Accept()
		go jianyan(conn)
	}
}
