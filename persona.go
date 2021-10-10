package main

import (
	"fmt"

	cc "github.com/fizzywhizbang/persona/ccgen"
	ss "github.com/fizzywhizbang/persona/ssngen"
)

func main() {
	// var available ListOfKnownTypes

	// for key, value := range AvailableCardTypes {
	// 	fmt.Println(key, " ", value.LongName)

	// }
	//get length of args

	creditCard := cc.GenerateCards("visa", 1)
	card := creditCard[0]
	fmt.Println(card.Pan.Formatted)
	// GenerateCards("visa", 1)
	ssn := ss.SSNgen()
	fmt.Println(ssn)

	for !ss.CheckValid(ssn) {
		ssn = ss.SSNgen()
	}
	fmt.Println(ssn)
}
