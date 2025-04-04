package main

import (
	"fmt"
	"learning/golang-advance-learning/pkg/interfaces/mysqldb"
	"log"
	"os"
)

type dbContract interface {
	Close()
	InsertUser(userName string) error
	SelectSingleUser(userName string) (string, error)
}

type Application struct {
	db dbContract
}

func (this Application) Run() {
	userName := "user1"

	err := this.db.InsertUser(userName)
	if err != nil {
		log.Println("couldn't insert user: %s", userName)
	}

	user, err := this.db.SelectSingleUser(userName)
	if err != nil {
		log.Println("couldn't fetch user: %s", userName)
	}

	fmt.Println(user)
}

func NewApplication(db dbContract) *Application {
	return &Application{
		db: db,
	}
}

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	db, err := mysqldb.New(dbUser, dbPassword, dbHost, dbPort, dbName)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	defer db.Close()

	app := NewApplication(db)
	app.Run()
}
