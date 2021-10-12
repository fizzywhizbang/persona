package persongen

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func PersonGen(count int) []Person {
	db, err := sql.Open("sqlite3", "database/persona.db") // Open the created SQLite File
	if err != nil {
		log.Println(err.Error())
		fmt.Println("coldn't open database")
	}
	smax := 151671
	fmax := 2000
	//loop
	var Personlist []Person

	for i := 0; i < count; i++ {
		r := rand.Intn(smax - 1)
		var lastname surname
		defer db.Close()
		q := `SELECT id,name FROM surnames WHERE id=?`
		err = db.QueryRow(q, r).Scan(&lastname.id, &lastname.name)

		if err != nil {
			log.Println(err.Error())
			fmt.Println("Something went wrong getting lastname")
		}
		var P Person
		P.Lid = lastname.id
		P.Last = strings.Title(strings.ToLower(lastname.name))

		r = rand.Intn(fmax - 1)
		q = `SELECT id,name,sex FROM firstnames WHERE id=?`
		var firstname first
		err = db.QueryRow(q, r).Scan(&firstname.id, &firstname.name, &firstname.sex)

		if err != nil {
			log.Println(err.Error())
			fmt.Println("Something went wrong getting firstname")
		}
		var person = Person{
			Fid:   firstname.id,
			First: firstname.name,
			Lid:   lastname.id,
			Last:  lastname.name,
			Sex:   firstname.sex,
		}
		Personlist = append(Personlist, person)
	}
	return Personlist
}
