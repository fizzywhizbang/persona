package ccgen

//export for persona
type Errors []Error

//struct export for persona
type Error struct {
	Parameter string `json:"parameter"`
	Issue     string `json:"error"`
}

//struct export for persona
type Card struct {
	Issuer     string     `json:"issuer"`
	Pan        PAN        `json:"pan"`
	ExpiryDate ExpiryDate `json:"expiryDate"`
	CVV        string     `json:"cvv"`
}

//struct export for persona
type ExpiryDate struct {
	Month int `json:"month"`
	Year  int `json:"year"`
}

//struct export for persona
type PAN struct {
	Raw       string `json:"raw"`
	Formatted string `json:"formatted"`
}

//struct export for persona
type CardProperties struct {
	LongName string
	Prefix   []string
	PanSize  int
	CvvSize  int
}

var Bin = map[string]CardProperties{
	//https://en.wikipedia.org/wiki/Payment_card_number
	"amex": {
		LongName: "American Express",
		Prefix:   []string{"37", "34"},
		PanSize:  15,
		CvvSize:  4,
	},
	"dci": {
		LongName: "Diners Club International",
		Prefix:   []string{"36", "38"},
		PanSize:  16,
		CvvSize:  3,
	},
	"dcu": {
		LongName: "Diners Club USA & Canada",
		Prefix:   []string{"54"},
		PanSize:  16,
		CvvSize:  3,
	},
	"discover": {
		LongName: "Discover",
		Prefix:   []string{"6011", "65", "644", "645", "646", "647", "648", "649"}, //6011, 622126–622925, 644-649, 65
		PanSize:  16,
		CvvSize:  3,
	},
	"jcb": {
		LongName: "Japan Credit Bureau",
		Prefix:   []string{"3528", "3529", "3530", "3531", "3532", "3533", "3534", "3535"}, //3528–3589
		PanSize:  16,
		CvvSize:  3,
	},
	"maeuk": {
		LongName: "Maestro UK",
		Prefix:   []string{"6759", "676770", "676774"},
		PanSize:  16, //12-19
		CvvSize:  3,
	},
	"mae": {
		LongName: "Maestro",
		Prefix:   []string{"5018", "5020", "5038", "5893", "6304", "6759", "6761", "6762", "6763"},
		PanSize:  16, //12-19
		CvvSize:  3,
	},
	"mc": {
		LongName: "Mastercard",
		Prefix:   []string{"51", "52", "53", "54", "55"},
		PanSize:  16,
		CvvSize:  3,
	},
	"visa": {
		LongName: "Visa",
		Prefix:   []string{"4"},
		PanSize:  16,
		CvvSize:  3,
	},
}
