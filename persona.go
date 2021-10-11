package main

import (
	"flag"
	"fmt"

	cc "github.com/fizzywhizbang/persona/ccgen"
	p "github.com/fizzywhizbang/persona/persongen"
	ss "github.com/fizzywhizbang/persona/ssngen"
)

func main() {

	typeFlag := flag.String("t", "visa", "amex, dci, dcu, discover, jcb, mae, maeui, mc, visa")
	qtyFlag := flag.Int("q", 1, "An integer for quantity")

	flag.Parse()
	fmt.Println("")
	fmt.Println("######################################")
	fmt.Println("type:", *typeFlag)
	fmt.Println("qty:", *qtyFlag)

	fmt.Println("######################################")
	fmt.Println("")

	creditCard := cc.GenerateCards(*typeFlag, *qtyFlag)

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
