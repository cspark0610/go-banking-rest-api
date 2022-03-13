package domain

import (
	"database/sql"
	"log"
	"time"

	"github.com/cspark0610/go-banking-rest-api/errs"
	_ "github.com/go-sql-driver/mysql"
)


type CustomerRepositoryDb struct {
	client *sql.DB
}

// make a function to finnd all customers
func (db CustomerRepositoryDb) FindAll() ([]Customer, error) {
	// make sql find all sentence
	findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"
	rows, err := db.client.Query(findAllSql)
	if err != nil {
		// si existe un error se imprime el mismo y se lo retorna
		log.Println("error while querying customers", err.Error())
		return nil, err
	}
	// make a slice of customers
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("error while scanning customers", err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (db CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	// make sql sentence
	findByIdSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = ?"
	// QueryRow method always returns only one row in DB
	row := db.client.QueryRow(findByIdSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found " + id)
		} 
		log.Println("error while scanning customer", err.Error())
		return nil, errs.NewNotFoundError("unexpected db error")
	}
	return &c, nil
}

//define a helper function to handle connection pool
func NewCustomerRepositoryDb() CustomerRepositoryDb {
	// define client and connection
	client, err := sql.Open("mysql", "root:root@tcp(localhost:3300)/banking")
	if err != nil {
		panic(err)
	}
	// set connection timeout
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	
	return CustomerRepositoryDb{client}
}
