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

	// ssn := genSSN()
	// for !checkValid(ssn) {
	// 	ssn = genSSN()
	// }
	// fmt.Println(ssn)
	creditCard := cc.GenerateCards("visa", 1)
	card := creditCard[0]
	fmt.Println(card.Issuer)
	// GenerateCards("visa", 1)
}
