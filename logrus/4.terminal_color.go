package main

import "fmt"

const (
	cBlack = iota
	cRed
	cGreen
	cYellow
	cBlue
	cPurple
	cCyan
	cGray
)

func PrintColor(colorCode int, text string, isBackground bool) {
	if isBackground {
		fmt.Printf("\033[4%dm %s \033[0m\n", colorCode, text)
	} else {
		fmt.Printf("\033[3%dm %s \033[0m\n", colorCode, text)
	}
}

func main() {
	PrintColor(cBlue, "藍色", false)
	PrintColor(cBlue, "藍色", true)

	// 前景色
	fmt.Println("\033[30m 黑色 \033[0m")
	fmt.Println("\033[31m 紅色 \033[0m")
	fmt.Println("\033[32m 綠色 \033[0m")
	fmt.Println("\033[33m 黃色 \033[0m")
	fmt.Println("\033[34m 藍色 \033[0m")
	fmt.Println("\033[35m 紫色 \033[0m")
	fmt.Println("\033[36m 青色 \033[0m")
	fmt.Println("\033[37m 灰色 \033[0m")

	// 背景色
	fmt.Println("\033[40m 黑色 \033[0m")
	fmt.Println("\033[41m 紅色 \033[0m")
	fmt.Println("\033[42m 綠色 \033[0m")
	fmt.Println("\033[43m 黃色 \033[0m")
	fmt.Println("\033[44m 藍色 \033[0m")
	fmt.Println("\033[45m 紫色 \033[0m")
	fmt.Println("\033[46m 青色 \033[0m")
	fmt.Println("\033[47m 灰色 \033[0m")
}
