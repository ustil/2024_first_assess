go
 
package main

import (
    "fmt"
)

func main() {
    var choice int
    fmt.Println("请选择您的性别：")
    fmt.Println("1. 男")
    fmt.Println("2. 女")
    fmt.Println("3. 其他")
    fmt.Scanln(&choice)

    var gender string
    switch choice {
    case 1:
        gender = "男"
    case 2:
        gender = "女"
    case 3:
        gender = "其他"
    default:
        fmt.Println("无效的选择，请重新输入。")
        return
    }

    fmt.Printf("您选择的性别是：%s\n", gender)
}