package main

import (
	"bufio"
	"fmt"
	"github.com/icntswm/test_project/encr_src"
	"os"
	"strconv"
	"strings"
)

func printHint(pos int) {
	switch pos {
	case 0:
		fmt.Println("Please enter the sentence you want to encrypt:")
	case 1:
		fmt.Println("Please enter encryption shift: ")
	}
	fmt.Print("-> ")
}

func receivingDataFromUser() ([]rune, int) {
	reader := bufio.NewReader(os.Stdin)

	printHint(0)
	text, err := reader.ReadString('\n')
	if err != nil || text[0] == '\n' || len(text) == 1 {
		os.Exit(1)
	}
	printHint(1)
	number_str, err := reader.ReadString('\n')
	number_str = strings.TrimRight(number_str, "\n")
	if err != nil || number_str[0] == '\n' {
		os.Exit(1)
	}
	number, err := strconv.Atoi(number_str)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return []rune(text), number
}

func main() {
	date, number := receivingDataFromUser()
	fmt.Println(string(encr_src.CaesarCipher(date, number)))
}
