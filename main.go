package main

import (
	"log"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/database"
)

func main() {
	if err := database.Open(); err != nil {
		log.Fatal(err)
	}
}
