package main
import(
	"fmt"
	"net")
func main(){
	fmt.Println("客户端")
	conn,err := net.dial("tcp","127.0.0.1:8888")
    if err !=nil {
		fmt.Println("连接失败：err:",err)
		return
	} 
    fmt.Println("连接成功， conn："，conn)
}


