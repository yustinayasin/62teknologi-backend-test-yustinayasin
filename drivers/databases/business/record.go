package businesses

import (
	businesses "yelp-business/business/businesses"
	"yelp-business/business/categories"
	"yelp-business/business/coordinates"
	"yelp-business/business/locations"
	"yelp-business/business/transactions"
)

type Business struct {
	Id            int
	Alias         string
	Name          string
	ImageUrl      string
	IsClosed      bool
	Url           string
	ReviewCount   int
	Rating        float64
	CoordinateId  int // ID of associated coordinate
	TransactionId int // ID of associated transaction
	Price         float64
	LocationId    int // ID of associated location
	Phone         string
	DisplayPhone  string
	Distance      float64
}

type ScannedBusiness struct {
	Id            int
	Alias         string
	Name          string
	ImageUrl      string
	IsClosed      bool
	Url           string
	ReviewCount   int
	Rating        float64
	CoordinateId  int // ID of associated coordinate
	TransactionId int // ID of associated transaction
	Price         float64
	LocationId    int // ID of associated location
	Phone         string
	DisplayPhone  string
	Distance      float64
	Coordinates   coordinates.Coordinate
	Transactions  transactions.Transaction
	Locations     locations.Location
	Categories    []categories.Category
}

func (business Business) ToUsecase() businesses.Business {
	return businesses.Business{
		Id:            business.Id,
		Alias:         business.Alias,
		Name:          business.Name,
		ImageUrl:      business.ImageUrl,
		IsClosed:      business.IsClosed,
		Url:           business.Url,
		ReviewCount:   business.ReviewCount,
		Rating:        business.Rating,
		CoordinateId:  business.CoordinateId,
		TransactionId: business.TransactionId,
		Price:         business.Price,
		LocationId:    business.LocationId,
		Phone:         business.Phone,
		DisplayPhone:  business.DisplayPhone,
		Distance:      business.Distance,
	}
}

func (business ScannedBusiness) ToUsecase() businesses.ScannedBusiness {
	return businesses.ScannedBusiness{
		Id:            business.Id,
		Alias:         business.Alias,
		Name:          business.Name,
		ImageUrl:      business.ImageUrl,
		IsClosed:      business.IsClosed,
		Url:           business.Url,
		ReviewCount:   business.ReviewCount,
		Rating:        business.Rating,
		CoordinateId:  business.CoordinateId,
		TransactionId: business.TransactionId,
		Price:         business.Price,
		LocationId:    business.LocationId,
		Phone:         business.Phone,
		DisplayPhone:  business.DisplayPhone,
		Distance:      business.Distance,
		Categories:    business.Categories,
		Coordinates:   business.Coordinates,
		Transactions:  business.Transactions,
		Locations:     business.Locations,
	}
}

func ToUsecaseList(business []Business) []businesses.Business {
	var newBusinesses []businesses.Business

	for _, v := range business {
		newBusinesses = append(newBusinesses, v.ToUsecase())
	}
	return newBusinesses
}

func FromUsecase(business businesses.Business) Business {
	return Business{
		Id:            business.Id,
		Alias:         business.Alias,
		Name:          business.Name,
		ImageUrl:      business.ImageUrl,
		IsClosed:      business.IsClosed,
		Url:           business.Url,
		ReviewCount:   business.ReviewCount,
		Rating:        business.Rating,
		CoordinateId:  business.CoordinateId,
		TransactionId: business.TransactionId,
		Price:         business.Price,
		LocationId:    business.LocationId,
		Phone:         business.Phone,
		DisplayPhone:  business.DisplayPhone,
		Distance:      business.Distance,
	}
}

// func (business ScannedBusiness) FromUsecase(business businesses.ScannedBusiness) ScannedBusiness {
// 	return ScannedBusiness{
// 		Id:            business.Id,
// 		Alias:         business.Alias,
// 		Name:          business.Name,
// 		ImageUrl:      business.ImageUrl,
// 		IsClosed:      business.IsClosed,
// 		Url:           business.Url,
// 		ReviewCount:   business.ReviewCount,
// 		Rating:        business.Rating,
// 		CoordinateId:  business.CoordinateId,
// 		TransactionId: business.TransactionId,
// 		Price:         business.Price,
// 		LocationId:    business.LocationId,
// 		Phone:         business.Phone,
// 		DisplayPhone:  business.DisplayPhone,
// 		Distance:      business.Distance,
// 		Categories:    business.Categories,
// 		Coordinates:   business.Coordinates,
// 		Transactions:  business.Transactions,
// 		Locations:     business.Locations,
// 	}
// }
