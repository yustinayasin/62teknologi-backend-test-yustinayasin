package businesses

import (
	"errors"
	"fmt"
	businesses "yelp-business/business/businesses"
	"yelp-business/helpers"
)

type BusinessRepository struct {
	db *helpers.Database
}

func NewBusinessRepository(database *helpers.Database) businesses.BusinessRepoInterface {
	//yang direturn adalah interfacenya repo
	return &BusinessRepository{
		db: database,
	}
}

func (repo *BusinessRepository) Add(business businesses.Business) (businesses.Business, error) {
	businessDB := FromUsecase(business)
	//connection database
	db, err := helpers.NewDatabase()

	if err != nil {
		return businesses.Business{}, err
	}

	//check if the db is connect
	if db == nil {
		fmt.Println("Database connection is nil")
		return businesses.Business{}, errors.New("database connection is nil")
	}

	_, err = db.DB.Exec("INSERT INTO businesses (alias, name, image_url, is_closed, url, review_count, rating, id_coordinates, id_transactions, price, id_locations, phone, display_phone, distance) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)", businessDB.Alias, businessDB.Name, businessDB.ImageUrl, businessDB.IsClosed, businessDB.Url, businessDB.ReviewCount, businessDB.Rating, businessDB.CoordinateId, businessDB.TransactionId, businessDB.Price, businessDB.LocationId, businessDB.Phone, businessDB.DisplayPhone, businessDB.Distance)

	if err != nil {
		fmt.Println("Error in repo: ", err)
		return businesses.Business{}, err
	}

	return businessDB.ToUsecase(), nil
}
