package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Even/Odd checker")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter a number:")
		fmt.Print("-> ")
		num, err := readInput(reader)
		if err != nil {
			fmt.Println("Not valid input")
			continue
		}

		if num%2 == 0 {
			fmt.Println("Even!")
		} else {
			fmt.Println("Odd!")
		}

	}

}

func readInput(reader *bufio.Reader) (int, error) {
	input, _ := reader.ReadString('\n')
	num := strings.TrimSpace(input)
	return strconv.Atoi(num)
}
