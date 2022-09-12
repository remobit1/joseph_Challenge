package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	r, _ := regexp.Compile(`(^[456][0-9]{3}\-?)([0-9]{4}\-?)([0-9]{4}\-?)([0-9]{4}\-?)$`)

	fmt.Println("Submit input cases... q to quit")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		creditCardNumber := scanner.Text()

		if creditCardNumber == "q" || creditCardNumber == "Q" {
			fmt.Println("Feel free to submit credit card numbers anytime!")
			os.Exit(0)
		}

		if r.MatchString(creditCardNumber) /* && checkForRepeatCharacters(creditCardNumber) */ {
			fmt.Println("Valid")
		} else {
			fmt.Println("Invalid")
		}
	}
}
