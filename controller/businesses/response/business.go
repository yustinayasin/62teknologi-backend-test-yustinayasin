package response

import "yelp-business/business/businesses"

type BusinessResponse struct {
	Id            int     `json:"id"`
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
	DisplayPhone  string  `json:"displayName"`
	Distance      float64 `json:"distance"`
}

func FromUsecase(business businesses.Business) BusinessResponse {
	return BusinessResponse{
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

func FromUsecaseList(user []businesses.Business) []BusinessResponse {
	var businessResponse []BusinessResponse

	for _, v := range user {
		businessResponse = append(businessResponse, FromUsecase(v))
	}

	return businessResponse
}
