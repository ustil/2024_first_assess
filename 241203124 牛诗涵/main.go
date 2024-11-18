package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	numberToGuess := rand.Intn(100) + 1
	maxTries := 10
	tries := 0
	fmt.Println("猜数字游戏！范围是1-100。你有10次机会。")
	for tries < maxTries {
		fmt.Print("请输入你猜的数字：")
		var guess int
		_, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Println("输入错误，请再次输入")
			continue5
		}
		tries++
		if guess < numberToGuess {
			fmt.Println("低了低了，再来一把。")
		} else if guess > numberToGuess {
			fmt.Println("高了搞了，再试一把。")
		} else {
			fmt.Printf("恭喜你！你猜对了数字是 %d，你用了 %d 次机会。\n", numberToGuess, tries)
			break
		}
		if tries == maxTries {
			fmt.Printf("很遗憾，你没有猜中。数字是 %d。\n", numberToGuess)
		}
	}
}
