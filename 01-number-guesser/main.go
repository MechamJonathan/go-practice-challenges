package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to Number Guesser!")
	fmt.Println(("-------------------------"))

	reader := bufio.NewReader(os.Stdin)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		target := r.Intn(100) + 1
		fmt.Println("> I've picked a random number (1-100). Guess it.")

		for {
			fmt.Print("-> ")
			input := readInput(reader)

			users_guess, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Invalid Input: ", err)
			}

			if users_guess < target {
				fmt.Println("higher")
			} else if users_guess > target {
				fmt.Println("lower")
			} else {
				fmt.Println("Nice... You guessed it correctly.")
				break
			}
		}

		fmt.Println("Want to play again? Y/N")
		if strings.ToUpper(readInput(reader)) != "Y" {
			break
		}
	}
}

func readInput(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
