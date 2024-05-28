package main

import (
	"database/sql"
	"log"

	// import pkg for side effects
	_ "github.com/lib/pq"
)

type Product struct {
	Name string
	Price float64
	Available bool
}

/* Product Table
- ID
- Name
- Price
- Available
- Date Created
*/

func main(){
	connStr := "postgres://postgres:secretpsd@localhost:5432/gopgtest?sslmode=disable"
	// do not hard code the password, conn to gopgtest db

	db, err := sql.Open("postgres", connStr)

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	createProductTable(db)
}

// take a pointer to it as an argument
func createProductTable(db *sql.DB){
	query := `CREATE TABLE IF NOT EXISTS product (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		price NUMERIC(6,2) NOT NULL,
		available BOOLEAN,
		created timestamp DEFAULT NOW()
	)`

	// returns a result and an err
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
		// quit the program
	}


}