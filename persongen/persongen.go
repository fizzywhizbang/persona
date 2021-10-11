package persongen

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func PersonGen() Person {
	db, err := sql.Open("sqlite3", "database/persona.db") // Open the created SQLite File
	if err != nil {
		log.Println(err.Error())
		fmt.Println("coldn't open database")
	}
	max := 151671
	r := rand.Intn(max - 1)
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

	max = 2000
	r = rand.Intn(max - 1)
	q = `SELECT id,name,sex FROM firstnames WHERE id=?`
	var firstname first
	err = db.QueryRow(q, r).Scan(&firstname.id, &firstname.name, &firstname.sex)

	if err != nil {
		log.Println(err.Error())
		fmt.Println("Something went wrong getting firstname")
	}
	P.Fid = firstname.id
	P.First = firstname.name
	P.Sex = firstname.sex

	return P
}
