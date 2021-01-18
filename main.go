package main

import (
	"log"
	"time"
)

func main() {

	budgetFirstDay, err := time.Parse("20060102", "20210401")
	lastDay, err := time.Parse("20060102", "20210301")

	log.Println(budgetFirstDay.Before(lastDay))
	log.Println(err)
}
