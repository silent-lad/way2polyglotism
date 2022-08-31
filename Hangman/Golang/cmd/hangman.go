package main

import (
	"fmt"
	"example.com/hangman/pkg/utils"
)

func main(){
	fmt.Println("Welcome to hangman!!!")

	fmt.Println(utils.GetHangmanAscii(6))

	utils.ClearTerminal()

	fmt.Println(utils.GetWord())
}