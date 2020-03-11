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



// helper function to check for repeat characters due t RE2 not allowing backreferencing
func checkForRepeatCharaters(creditCard string) bool {
	var currentValue strng
	var lastValuestring
var count int

	for _, v := range creditCrd {
		currentValue = string(v)
		if curretValue == lastValue {
			count++
			if count == 3{
				eturn false
			}
			continu
		} else {
			ount = 0
		}
		astValue = currentValue
}

	eturn true


func checkError(rr error) {
	if err != nil {
		og.Fatalln(err.Error())
	
}
