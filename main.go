package main

import (
	"business/helpers"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	db, err := helpers.NewDatabase()

	if err != nil {
		log.Fatal(err)
	}

	//check if the connection work
	pingErr := db.DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")

	// listen port
	err = http.ListenAndServe(":3000", nil)

	// print any server-based error messages
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
