package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	cc "github.com/fizzywhizbang/persona/ccgen"
	p "github.com/fizzywhizbang/persona/persongen"
	ss "github.com/fizzywhizbang/persona/ssngen"
)

func main() {

	typeFlag := flag.String("t", "visa", "amex, dci, dcu, discover, jcb, mae, maeui, mc, visa")
	qtyFlag := flag.Int("q", 1, "An integer for quantity")
	writeToFile := flag.Bool("w", false, "if included it will be true")

	flag.Parse()
	fmt.Println("")
	fmt.Println("######################################")
	fmt.Println("-t type:", *typeFlag)
	fmt.Println("-q qty:", *qtyFlag)
	fmt.Println("-w write to file:", *writeToFile)
	fmt.Println("######################################")
	fmt.Println("")

	creditCard := cc.GenerateCards(*typeFlag, *qtyFlag)

	//loop through cards and supply ssn and persona
	for i := 0; i < len(creditCard); i++ {
		card := creditCard[i]
		person := p.PersonGen()
		payload := card.Issuer + "\n"
		payload += card.Pan.Formatted + "\n"
		payload += "exp:" + strconv.Itoa(card.ExpiryDate.Month) + strconv.Itoa(card.ExpiryDate.Year) + " CVV:" + card.CVV + "\n"
		payload += person.First + person.Last + "(" + person.Sex + ")\n"
		ssn := ss.SSNgen()
		for !ss.CheckValid(ssn) {
			ssn = ss.SSNgen()
		}
		payload += "SSN:" + ssn
		fmt.Println(payload)
		fmt.Println("") //line break
		if *writeToFile {
			fname := person.First + "_" + person.Last + ".txt"
			writeToFileFunc(fname, payload)
		}
	}

}

func writeToFileFunc(filename string, payload string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Could not write file")
	}

	defer f.Close()

	_, err2 := f.WriteString(payload)

	if err2 != nil {
		fmt.Println("Could not write to file")
	}

}
