package main

import (
	"fmt"

	cc "github.com/fizzywhizbang/persona/ccgen"
	p "github.com/fizzywhizbang/persona/persongen"
	ss "github.com/fizzywhizbang/persona/ssngen"
)

func main() {
	// var available ListOfKnownTypes

	// for key, value := range AvailableCardTypes {
	// 	fmt.Println(key, " ", value.LongName)

	// }
	//get length of args

	creditCard := cc.GenerateCards("visa", 2)

	//loop through cards and supply ssn and persona
	for i := 0; i < len(creditCard); i++ {
		card := creditCard[i]
		person := p.PersonGen()
		fmt.Println(card)
		fmt.Println(person)
		ssn := ss.SSNgen()
		for !ss.CheckValid(ssn) {
			ssn = ss.SSNgen()
		}
		fmt.Println(ssn)
	}

}
