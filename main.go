package main

import (
	"log"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/database"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/router"
)

func main() {
	if err := database.Open(); err != nil {
		log.Fatal(err)
	}

	app := router.Start()
	log.Fatal(app.Run("127.0.0.1:8080"))
}
