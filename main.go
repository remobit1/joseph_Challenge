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

	fmt.Println("Checking input cases...")
	// testStr := "5223-6222-7961-4443"
	rdr := bufio.NewReader(os.Stdin)
	n, err := rdr.ReadBytes('\n')
	fmt.Println(int(n[0]))
	checkError((err))
	for i := 0; i < int(n[0]); i++ {
		creditCardNumber, err := rdr.ReadString('\n')
		fmt.Println(creditCardNumber)
		checkError(err)

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
