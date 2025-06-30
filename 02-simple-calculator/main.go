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
		first_digit, err := readIntInput(reader)
		if err != nil {
			fmt.Println("invalid input! Please enter a valid number.")
			continue
		}

		fmt.Println("Enter operator (+, -, *, /):")
		fmt.Print("-> ")
		operator := readInput(reader)

		fmt.Println("Enter second digit:")
		fmt.Print("-> ")
		second_digit, err := readIntInput(reader)
		if err != nil {
			fmt.Println("invalid input! Please enter a valid number.")
		}

		result, err := calculateResult(first_digit, operator, second_digit)
		if err != nil {
			fmt.Println("Error:", err)
		}

		fmt.Printf("Result: %d \n", result)
	}
}

func readInput(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func readIntInput(reader *bufio.Reader) (int, error) {
	input := readInput(reader)
	return strconv.Atoi(input)
}

func calculateResult(fd int, op string, sd int) (int, error) {
	switch op {
	case "+":
		return fd + sd, nil
	case "-":
		return fd - sd, nil
	case "*":
		return fd * sd, nil
	case "/":
		if sd == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}
		return fd / sd, nil
	default:
		return 0, fmt.Errorf("unsupported operator: %s", op)
	}
}
