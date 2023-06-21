package funcserver

import (
	"fmt"
	"time"
)

func PostRequest(url string) {

	usd := GetRespondBody(url)

	addMoney(usd)

}

func GetAll() {
	fmt.Println(moneyTable)
}

func GetById() {
	index := ""
	fmt.Println("ID? ")
	fmt.Scanln(&index)
	GetConversion(index)
}

func DeleteById() {
	index := ""
	fmt.Println("ID? ")
	fmt.Scanln(&index)
	DeleteConversion(index)
}

func PostTimer(url string) {

	for i := 0; i < 2; i++ {

		newtimer := time.NewTimer(10 * time.Second)
		<-newtimer.C
		PostRequest(url)

	}

}
