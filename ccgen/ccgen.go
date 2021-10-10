package ccgen

import (
	"fmt"
	"math/rand"
	"time"
)

const MaxCount = 500

func GenerateCards(cardType string, count int) []Card {

	// count, countError := strconv.Atoi(countString)
	cardProperties, exists := bin[cardType]

	var errors Errors

	if !exists {
		errors = append(errors, Error{
			Parameter: "type",
			Issue:     "invalid card type supplied",
		})
	}

	if len(errors) > 0 {
		fmt.Println(errors)
	}

	if count > MaxCount {
		count = MaxCount
	}

	var cardList []Card
	for i := 0; i < count; i++ {

		rand.Seed(time.Now().UnixNano())

		var card = Card{
			Issuer:     cardProperties.LongName,
			Pan:        GeneratePAN(cardProperties),
			ExpiryDate: GenerateExpiryDate(),
			CVV:        GenerateCVV(cardProperties.CvvSize),
		}

		cardList = append(cardList, card)
	}
	return cardList
	// c.JSON(http.StatusOK, cardList)
}
