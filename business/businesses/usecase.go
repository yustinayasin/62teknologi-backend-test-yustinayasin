package businesses

import (
	"errors"
)

type BusinessUseCase struct {
	repo BusinessRepoInterface
}

func NewUseCase(businessRepo BusinessRepoInterface) BusinessUseCaseInterface {
	return &BusinessUseCase{
		repo: businessRepo,
	}
}

func (businessUseCase *BusinessUseCase) Add(business Business) (Business, error) {
	//check request
	if business.Name == "" {
		return Business{}, errors.New("Name empty")
	}

	if business.Alias == "" {
		return Business{}, errors.New("Alias empty")
	}

	if business.ImageUrl == "" {
		return Business{}, errors.New("Image url empty")
	}

	if business.Url == "" {
		return Business{}, errors.New("Url empty")
	}

	if business.ReviewCount == 0 {
		return Business{}, errors.New("Review count empty")
	}

	if business.Rating == 0.0 {
		return Business{}, errors.New("Rating empty")
	}

	if business.Price == 0.0 {
		return Business{}, errors.New("Price empty")
	}

	if business.Phone == "" {
		return Business{}, errors.New("Phone empty")
	}

	if business.DisplayPhone == "" {
		return Business{}, errors.New("Display phone empty")
	}

	if business.Distance == 0.0 {
		return Business{}, errors.New("Distance empty")
	}

	if business.CoordinateId == 0 {
		return Business{}, errors.New("Coordinate empty")
	}

	if business.TransactionId == 0 {
		return Business{}, errors.New("Transaction empty")
	}

	if business.LocationId == 0 {
		return Business{}, errors.New("Location empty")
	}

	businessRepo, err := businessUseCase.repo.Add(business)

	if err != nil {
		return Business{}, err
	}

	return businessRepo, nil
}

func (businessUseCase *BusinessUseCase) Edit(business Business, id int) (Business, error) {
	if id == 0 {
		return Business{}, errors.New("Business ID empty")
	}

	if business.Name == "" {
		return Business{}, errors.New("Name empty")
	}

	if business.Alias == "" {
		return Business{}, errors.New("Alias empty")
	}

	if business.ImageUrl == "" {
		return Business{}, errors.New("Image url empty")
	}

	if business.Url == "" {
		return Business{}, errors.New("Url empty")
	}

	if business.ReviewCount == 0 {
		return Business{}, errors.New("Review count empty")
	}

	if business.Rating == 0.0 {
		return Business{}, errors.New("Rating empty")
	}

	if business.Price == 0.0 {
		return Business{}, errors.New("Price empty")
	}

	if business.Phone == "" {
		return Business{}, errors.New("Phone empty")
	}

	if business.DisplayPhone == "" {
		return Business{}, errors.New("Display phone empty")
	}

	if business.Distance == 0.0 {
		return Business{}, errors.New("Distance empty")
	}

	if business.CoordinateId == 0 {
		return Business{}, errors.New("Coordinate empty")
	}

	if business.TransactionId == 0 {
		return Business{}, errors.New("Transaction empty")
	}

	if business.LocationId == 0 {
		return Business{}, errors.New("Location empty")
	}

	businessRepo, err := businessUseCase.repo.Edit(business, id)

	if err != nil {
		return Business{}, err
	}

	return businessRepo, nil
}
