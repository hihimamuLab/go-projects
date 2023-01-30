package main

import (
	"fmt"
)

func main() {
	// 文字出力
	fmt.Println("Hello,World!")

	// 演算
	fmt.Println(3 + 2)
	fmt.Println("3" + "2")

	// 変数
	var message string
	message = "Hello, World!"
	var message2 = "Hello, World!"
	message3 := "Hello, World!"
	number := 42
	println(message, message2, message3, number)

	// 変数の演算
	number2 := 13
	number3 := number + number2
	fmt.Println(number, number2, number3)

	// 条件分岐if文
	fmt.Println(if_result(90))

	// 条件分岐switch文
	fmt.Println(game_win(1))
}

// 条件分岐処理関数
func if_result(score int) string {
	if score >= 80 {
		return "合格"
	} else {
		return "不合格"
	}
}

// 条件分岐処理関数
func game_win(score int) string {
	switch score {
	case 1:
		return "勝"
	case 0:
		return "負"
	}
	return "1か0を代入してください"
}
