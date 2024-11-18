package main
import(
	"fmt"
	"net")

func main(){
	fmt.Println("服务器")
	listen,err := net.Listen("tcp",“127.0.0.1:8888")
	if err != nil{fmt.Println("监听失败，err:",err)
	return
}
	for{
	    conn,err2 := listen.Accept()
		if err2!= nil {
		fmt.Println("失败,err2:",err2)
	}else{
	fmt.Printf("链接成功，con=%v ，客户端信息:%v\n",conn,conn.RemoteAddr().String())
    }
	}
}