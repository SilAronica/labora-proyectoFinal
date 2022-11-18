package services

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "SilyLeo21 "
	dbname   = "DB_1"
)

var dbConnection *sql.DB

func InitDb() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s"+"password=%s dnname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	dbConnection, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer dbConnection.Close()
	err = dbConnection.Ping()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	log.Println(dbConnection)
	fmt.Println("Successfuly connected to db!")
}
