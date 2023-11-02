package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"yelp-business/helpers"

	businessUsecase "yelp-business/business/businesses"
	businessController "yelp-business/controller/businesses"
	businessRepo "yelp-business/drivers/databases/business"

	"github.com/gorilla/mux"
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

	businessRepoInterface := businessRepo.NewBusinessRepository(db)
	businessUseCaseInterface := businessUsecase.NewUseCase(businessRepoInterface)
	businessControllerInterface := businessController.NewBusinessController(businessUseCaseInterface)

	r := mux.NewRouter()

	r.HandleFunc("/business/add", businessControllerInterface.Add)

	// listen port
	err = http.ListenAndServe(":3000", r)

	// print any server-based error messages
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
