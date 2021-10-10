package main

import (
	"fmt"
)

func main() {
	// var available ListOfKnownTypes

	// for key, value := range AvailableCardTypes {
	// 	fmt.Println(key, " ", value.LongName)

	// }
	//get length of args

	ssn := genSSN()
	for !checkValid(ssn) {
		ssn = genSSN()
	}
	fmt.Println(ssn)

	GenerateCards("visa", 1)
}
