package request

import "yelp-business/business/businesses"

type Business struct {
	Alias         string  `json:"alias"`
	Name          string  `json:"name"`
	ImageUrl      string  `json:"imageUrl"`
	IsClosed      bool    `json:"isClosed"`
	Url           string  `json:"url"`
	ReviewCount   int     `json:"reviewCount"`
	Rating        float64 `json:"rating"`
	CoordinateId  int     `json:"coordinateId"`
	TransactionId int     `json:"transactionId"`
	Price         float64 `json:"price"`
	LocationId    int     `json:"locationId"`
	Phone         string  `json:"phone"`
	DisplayPhone  string  `json:"displayPhone"`
	Distance      float64 `json:"distance"`
}

func (business *Business) ToUsecase() *businesses.Business {
	return &businesses.Business{
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
