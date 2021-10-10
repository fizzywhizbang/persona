package main

import (
	"fmt"

	cc "github.com/fizzywhizbang/persona/ccgen"
)

func main() {
	// var available ListOfKnownTypes

	// for key, value := range AvailableCardTypes {
	// 	fmt.Println(key, " ", value.LongName)

	// }
	//get length of args
	ssn := ss.ssngen()

	// for !ss.checkValid(ssn) {
	//  	ssn = ss.genSSN()
	// }
	fmt.Println(ssn)
	creditCard := cc.GenerateCards("visa", 1)
	card := creditCard[0]
	fmt.Println(card.Pan.Formatted)
	// GenerateCards("visa", 1)
}
