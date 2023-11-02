package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	businesses "yelp-business/business/businesses"
	"yelp-business/controller/businesses/request"
	"yelp-business/utils"
)

type BusinessController struct {
	usecase businesses.BusinessUseCaseInterface
}

// dipasangkan dengan routing
func NewBusinessController(uc businesses.BusinessUseCaseInterface) *BusinessController {
	return &BusinessController{
		usecase: uc,
	}
}

// Add a user handler
func (controller *BusinessController) Add(res http.ResponseWriter, req *http.Request) {
	// check the method
	if req.Method != "POST" {
		// Add the response return message
		HandlerMessage := []byte(`{
		"success": false,
		"message": "Check your HTTP method: Invalid HTTP method executed",
		}`)

		utils.ReturnJsonResponse(res, http.StatusMethodNotAllowed, HandlerMessage)
		return
	}

	var business request.Business

	// Read the request body
	bodyBytes, err := ioutil.ReadAll(req.Body)

	if err != nil {
		HandlerMessage := []byte(`{
		"success": false,
		"message": "Error read request body",
		}`)

		utils.ReturnJsonResponse(res, http.StatusMethodNotAllowed, HandlerMessage)
		return
	}

	// Create a new io.ReadCloser with the same content
	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	// parse the user data into json format
	err = json.NewDecoder(req.Body).Decode(&business)

	if err != nil {
		HandlerMessage := []byte(`{
		"success": false,
		"message": "Error parsing the business data",
		}`)

		utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
		return
	}

	defer req.Body.Close()

	_, errRepo := controller.usecase.Add(*business.ToUsecase())

	if errRepo != nil {
		HandlerMessage := []byte(`{
		"success": false,
		"message": "Error in repo,
		}`)

		utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
		return
	}

	// Add the response return message
	HandlerMessage := []byte(`{
	 "success": true,
	 "message": "Business was successfully created",
	 }`)

	utils.ReturnJsonResponse(res, http.StatusCreated, HandlerMessage)
	return
}
