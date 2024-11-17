package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	Number := rand.Intn(100) + 1

	var guess int
	var tries int

	fmt.Println("欢迎来到猜数字小游戏！")
	fmt.Println("我已经想好了一个1到100之间的数字，你来猜猜看。")

	for {
		fmt.Print("请输入你的猜测: ")
		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Println("输入有误，请重新输入。")
			continue
		}

		tries++

		if guess == Number {
			fmt.Printf("恭喜你，猜对了！你一共猜了 %d 次。\n", tries)
			break
		} else if guess < Number {
			fmt.Println("猜小了，再试试。")
		} else {
			fmt.Println("猜大了，再试试。")
		}
	}
}
