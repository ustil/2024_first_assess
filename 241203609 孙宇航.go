package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 创建HTTP客户端
	client := &http.Client{}
	// 百度首页
	req, err := http.NewRequest("GET", "https://www.baidu.com", nil)
	if err != nil {
		fmt.Printf("创建请求出错: %v\n", err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("发送请求出错: %v\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应内容出错: %v\n", err)
		return
	}
	fmt.Printf("响应状态码: %d\n", resp.StatusCode)
	fmt.Printf("响应内容: %s\n", string(body))
}
