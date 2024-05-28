package main

import (
	"database/sql"
	"fmt"
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

	// iterate over multiple rows of the db
	// create a slice of Product objects from the returned data from the db
	// append the returned results into the slice
	data := []Product{}
	rows, err := db.Query("SELECT name, available, price FROM product")
	if err != nil {
		log.Fatal(err)
	}
	// defer close rows is important to stop potential memory leaks
	defer rows.Close()
	// scan DB values
	var name string
	var available bool
	var price float64

	for rows.Next() {
		err := rows.Scan(&name, &available, &price)
		if err != nil {
			log.Fatal(err)
		}
		// append results into the data variable
		data = append(data, Product{name, price, available})
	}
	// print data to the terminal
	fmt.Println(data)
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

// insertProduct func, product of Type Product, which returns an int for the Primary Key of the product
// when the db is called from createProductTable(db) the Primary Key id is needed to perform further queries; var pk
// insert data into the product table.
// reference values using the $ syntax to paramatize the data, that helps stop SQL injection 
func insertProduct(db *sql.DB, product Product) int {
	query := `INSERT INTO product (name, price, available)
		VALUES ($1, $2, $3) RETURNING id`
	
	var pk int
	err := db.QueryRow(query, product.Name, product.Price, product.Available).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	return pk
}