package persongen

type first struct {
	id   int
	name string
	sex  string
}

type surname struct {
	id   int
	name string
}

type Person struct {
	Fid   int
	First string
	Lid   int
	Last  string
	Sex   string
}
