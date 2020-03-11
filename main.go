package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	r, _ := regexp.Compile(`(^[456][0-9]{3}\-?)([0-9]{4}\-?)([0-9]{4}\-?)([0-9]{4}\-?)$`)

	fmt.Println("Submit input cases...")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		creditCardNumber := scanner.Text()

		if creditCardNumber == "q" || creditCardNumber == "Q" {
			fmt.Println("Feel free to submit credit card numbers anytime!")
			os.Exit(0)
		}

		if r.MatchString(creditCardNumber) && checkForRepeatCharacters(creditCardNumber) {
			fmt.Println("Valid")
		} else {
			fmt.Println("Invalid")
		}
	}

}

// helper function to check for repeat characters due to RE2 not allowing backreferencing
func checkForRepeatCharacters(creditCard string) bool {
	var currentValue string
	var lastValue string
	var count int

	for _, v := range creditCard {
		currentValue = string(v)
		if currentValue == lastValue {
			count++
			if count == 3 {
				return false
			}
			continue
		} else {
			count = 0
		}
		lastValue = currentValue
	}

	return true
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
