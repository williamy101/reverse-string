package main

import "fmt"

func main() {
	var input string
	newStr := ""
	fmt.Scanln(&input)
	n := len(input)
	for i := n - 1; i >= 0; i-- {
		newStr += string(input[i])
	}
	fmt.Print(newStr)
}
