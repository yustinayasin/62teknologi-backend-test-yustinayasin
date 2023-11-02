package businesses

import businesses "yelp-business/business/businesses"

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
