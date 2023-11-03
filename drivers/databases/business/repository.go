package businesses

import (
	"errors"
	"fmt"
	"log"
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

func (repo *BusinessRepository) Edit(business businesses.Business, id int) (businesses.Business, error) {
	businessDB := FromUsecase(business)

	var newBusiness Business

	//connection database
	db, err := helpers.NewDatabase()

	if err != nil {
		return businesses.Business{}, err
	}

	if db == nil {
		fmt.Println("Database connection is nil")
		return businesses.Business{}, errors.New("database connection is nil")
	}

	_, err = db.DB.Exec("UPDATE businesses SET alias = $1, name = $2, image_url = $3, is_closed = $4, url = $5, review_count = $6, rating = $7, id_coordinates = $8, id_transactions = $9, price = $10, id_locations = $11, phone = $12, display_phone = $13, distance = $14 WHERE id = $15", businessDB.Alias, businessDB.Name, businessDB.ImageUrl, businessDB.IsClosed, businessDB.Url, businessDB.ReviewCount, businessDB.Rating, businessDB.CoordinateId, businessDB.TransactionId, businessDB.Price, businessDB.LocationId, businessDB.Phone, businessDB.DisplayPhone, businessDB.Distance, id)

	if err != nil {
		fmt.Println("Error in repo: ", err)
		return businesses.Business{}, err
	}

	err = db.DB.QueryRow("SELECT id, alias, name, image_url, is_closed, url, review_count, rating, id_coordinates, id_transactions, price, id_locations, phone, display_phone, distance FROM businesses WHERE id = $1", id).Scan(&newBusiness.Id, &newBusiness.Alias, &newBusiness.Name, &newBusiness.ImageUrl, &newBusiness.IsClosed, &newBusiness.Url, &newBusiness.ReviewCount, &newBusiness.Rating, &newBusiness.CoordinateId, &newBusiness.TransactionId, &newBusiness.Price, &newBusiness.LocationId, &newBusiness.Phone, &newBusiness.DisplayPhone, &newBusiness.Distance)

	if err != nil {
		log.Fatal(err)
	}

	return newBusiness.ToUsecase(), nil
}

func (repo *BusinessRepository) Delete(id int) (businesses.Business, error) {
	var businessDB Business

	//connection database
	db, err := helpers.NewDatabase()

	if err != nil {
		return businesses.Business{}, err
	}

	if db == nil {
		fmt.Println("Database connection is nil")
		return businesses.Business{}, errors.New("database connection is nil")
	}

	_, err = db.DB.Exec("DELETE FROM businesses WHERE ID = $1", id)

	//kalo ngecek ga ada id kayak gitu pake result kah?
	if err != nil {
		return businesses.Business{}, err
	}

	return businessDB.ToUsecase(), nil
}

// func (repo *BusinessRepository) GetBusiness() ([]businesses.Business, error) {
// 	var scannedBusiness ScannedBusiness

// 	//connection database
// 	db, err := helpers.NewDatabase()

// 	if err != nil {
// 		return []businesses.Business{}, err
// 	}

// 	if db == nil {
// 		fmt.Println("Database connection is nil")
// 		return []businesses.Business{}, errors.New("database connection is nil")
// 	}

// 	row := db.DB.QueryRow("SELECT businesses.*, (SELECT json_agg(categories) FROM categories WHERE categories.id IN (SELECT category_id FROM business_categories WHERE business_id = businesses.id)) AS categories, json_build_object('id', coordinates.id, 'latitude', coordinates.latitude, 'longitude', coordinates.longitude) AS coordinate, json_build_object('id', locations.id, 'address1', locations.address1, 'address2', locations.address2, 'address3', locations.address3, 'city', locations.city, 'country', locations.country, 'zip_code', locations.zip_code, 'state', locations.state, 'display_address', locations.display_address) AS location,(SELECT json_agg(json_build_object('id', transactions.id, 'name', transactions.name)) FROM transactions WHERE transactions.id IN (SELECT transaction_id FROM business_transactions WHERE business_id = businesses.id)) AS transactions FROM businesses LEFT JOIN coordinates ON businesses.id_coordinates = coordinates.id LEFT JOIN locations ON businesses.id_locations = locations.id")

// 	if err := row.Scan(&scannedBusiness.ID, &scannedBusiness.Alias, &scannedBusiness.Name, &scannedBusiness.ImageURL, &scannedBusiness.IsClosed, &scannedBusiness.URL, &scannedBusiness.ReviewCount, &scannedBusiness.Rating, &scannedBusiness.Coordinates.ID, &scannedBusiness.Coordinates.Latitude, &scannedBusiness.Coordinates.Longitude, &scannedBusiness.Locations.ID, &scannedBusiness.Locations.Address1, &scannedBusiness.Locations.Address2, &scannedBusiness.Locations.Address3, &scannedBusiness.Locations.City, &scannedBusiness.Locations.Country, &scannedBusiness.Locations.ZipCode, &scannedBusiness.Locations.State, &scannedBusiness.Locations.DisplayAddress, &scannedBusiness.Categories, &scannedBusiness.Transactions); err != nil {
// 		if err == sql.ErrNoRows {
// 			return scannedBusiness, fmt.Errorf("albumsById %d: no such album", id)
// 		}
// 		return scannedBusiness, fmt.Errorf("albumsById %d: %v", id, err)
// 	}

// 	if result.Error != nil {
// 		return []businesses.Business{}, result.Error
// 	}

// 	Kalo mau mengubah array
// 	return ToUsecaseList(businessesDb), nil
// }
