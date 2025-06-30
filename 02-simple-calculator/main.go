package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Simple Calculator!")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter first digit:")
		fmt.Print("-> ")
		first_digit, err := strconv.Atoi(readInput(reader))
		if err != nil {
			fmt.Println("invalid input")
		}

		fmt.Println("Enter operator:")
		fmt.Print("-> ")
		operator := readInput(reader)

		fmt.Println("Enter second digit:")
		fmt.Print("-> ")
		second_digit, err := strconv.Atoi(readInput(reader))
		if err != nil {
			fmt.Println("invalid input")
		}

		fmt.Println("Result: ", calculateResult(first_digit, operator, second_digit))
	}
}

func readInput(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func calculateResult(fd int, op string, sd int) int {
	result := 0
	switch op {
	case "+":
		result = fd + sd
	case "-":
		result = fd - sd
	case "*":
		result = fd * sd
	case "/":
		result = fd / sd
	}
	return result
}
