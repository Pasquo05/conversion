package funcserver

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type money struct {
	Id        uint `gorm:"primary key;autoIncrement"`
	Currfrom  string
	Currto    string
	Value     float64
	Createdat string
}

type rates struct {
	EUR float64
}

var db *gorm.DB

var moneyTable []money

func NewConnection(dsn string) {

	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("connection filed")
	}

	fmt.Println(db)
	fmt.Println("CONNECTED")

	err = db.Find(&moneyTable).Error

	if err != nil {
		fmt.Println("Errore", err)
	}

}

func addMoney(usd Prova) money {

	val := money{}

	val.Currfrom = "EUR"
	val.Currto = "USD"
	val.Value = usd.Rates.EUR
	val.Createdat = time.Unix(usd.Timestamp, 0).Format("2006-01-02 15:04:05")

	db.Create(&val)

	/*
		err := db.Find(&moneyTable).Error

		if err != nil {
			fmt.Println("Errore", err)
		}

		fmt.Println(moneyTable)

	*/
	return val
}

func GetConversion(index string) money {

	err := db.Find(&moneyTable).Error

	if err != nil {
		panic(err)
	}

	usd := money{}
	for _, conversion := range moneyTable {
		if fmt.Sprintf("%v", conversion.Id) == index {
			usd = conversion
			fmt.Println(usd)
			return usd
		}
	}

	fmt.Println(usd)
	return usd

}

func DeleteConversion(index string) {

	db.Find(&moneyTable)
	for _, conversion := range moneyTable {
		if fmt.Sprintf("%v", conversion.Id) == index {
			db.Delete(&moneyTable, conversion.Id)
			db.Find(&moneyTable)
		}
	}

	fmt.Println(moneyTable)
}
