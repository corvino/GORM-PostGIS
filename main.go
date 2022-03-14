package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type City struct {
	Name       string
	State      string
	Population int
}

var cities = []City{
	{
		Name:       "Chicago",
		State:      "IL",
		Population: 2_746_388,
	},
	{
		Name:       "New York",
		State:      "NY",
		Population: 8_804_190,
	},
	{
		Name:       "San Francisco",
		State:      "CA",
		Population: 873_965,
	},
	{
		Name:       "Seattle",
		State:      "WA",
		Population: 737_015,
	},
}

func insertData(db *gorm.DB) {
	for _, city := range cities {
		db.Create(&city)
	}
}

func connect(host, database, username, password string) (*gorm.DB, error) {
	dbURI := fmt.Sprintf("sslmode=disable host=%s dbname=%s  user=%s password=%s", host, database, username, password)

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Println("We cant op open a DATABASE")
		return nil, err
	}

	return db.Debug(), nil
}

func main() {
	db, err := connect("localhost", "gorm-postgis", "gorm-postgis", "gorm-postgis")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	db.AutoMigrate(&City{})

	insertData(db)
}
