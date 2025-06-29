package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Number Guesser!")
	fmt.Println(("-------------------------"))
	fmt.Println("> I've picked a random number (1-100). Guess it.")

	num := generateRandomNumber()

	for {
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')

		input = strings.Replace(input, "\n", "", -1)
		i, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid Input: ", err)
		}

		if i < num {
			fmt.Println("higher")
		} else if i > num {
			fmt.Println("lower")
		} else {
			fmt.Println("Nice... you guessed it correctly.")
			fmt.Println("Want to play again? Y/N")
			fmt.Print("-> ")

			input, _ := reader.ReadString('\n')
			input = strings.Replace(input, "\n", "", -1)

			if input == "Y" {
				num = generateRandomNumber()
				fmt.Println("I picked a new number...")
				continue
			} else {
				fmt.Println("see ya!")
				break
			}
		}
	}
}

func generateRandomNumber() int {
	num := rand.Intn(100)
	return num
}
