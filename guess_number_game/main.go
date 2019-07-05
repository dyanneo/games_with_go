package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	low := 1
	high := 100
	tries := 0
	fmt.Println("Please think of a number between", low, " and", high)
	fmt.Println("Press enter when ready")
	scanner.Scan()

	for {
		guess := (low + high) / 2
		fmt.Println("I guess the number is", guess)
		fmt.Println("Is that:")
		fmt.Println("(a) too high?")
		fmt.Println("(b) too low?")
		fmt.Println("(c) correct?")
		scanner.Scan()
		response := scanner.Text()

		if response == "a" {
			if guess == high {
				fmt.Println("You picked a number smaller than 0. No cheating.")
				break
			}
			high = guess - 1
			tries++
		} else if response == "b" {
			if guess == low {
				fmt.Println("You picked a number larger than 100. No cheating.")
				break
			}
			low = guess + 1
			tries++
		} else if response == "c" {
			tries++
			fmt.Printf("I won in %d tries!", tries)
			break
		} else {
			fmt.Println("Invalid response, try again")
		}
	}
}
