package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	cc "github.com/fizzywhizbang/persona/ccgen"
	p "github.com/fizzywhizbang/persona/persongen"
	ss "github.com/fizzywhizbang/persona/ssngen"
)

func main() {

	if os.Args[1] == "help" {
		message := "Welcome to help\n"
		message += "Version 0.1\n"
		message += "Command line usage is\n"
		message += "go run . type=visa qty=1\n"
		message += "Card Types:amex, dci, dcu, discover, jcb, mae, maeui, mc, visa\n"
		fmt.Println(message)
		fmt.Println(cc.Bin)
		fmt.Println("_________________________________________________")

	} else {
		if os.Args[1] == "gui" {
			fmt.Println("start qui")
		} else {
			c := strings.Split(os.Args[1], "=")
			cardType := c[1]

			q := strings.Split(os.Args[2], "=")
			qty, _ := strconv.Atoi(q[1])

			creditCard := cc.GenerateCards(cardType, qty)

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

	}

}
