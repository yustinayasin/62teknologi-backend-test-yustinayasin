package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	businesses "yelp-business/business/businesses"
	"yelp-business/controller/businesses/request"
	"yelp-business/utils"

	"github.com/gorilla/mux"
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

func (controller *BusinessController) Edit(res http.ResponseWriter, req *http.Request) {
	// check the method
	if req.Method != "PUT" {
		// Add the response return message
		HandlerMessage := []byte(`{
		"success": false,
		"message": "Check your HTTP method: Invalid HTTP method executed",
		}`)

		utils.ReturnJsonResponse(res, http.StatusMethodNotAllowed, HandlerMessage)
		return
	}

	vars := mux.Vars(req)
	id := vars["id"]

	var business request.Business

	// Read the request body
	bodyBytes, err := ioutil.ReadAll(req.Body)

	if err != nil {
		// Handle error
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
		// Add the response return message
		HandlerMessage := []byte(`{
		"success": false,
		"message": "Error parsing the user data",
		}`)

		utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
		return
	}

	//defer ensure req.Body.Close() will be executed after the AddMovie or schedule a function
	defer req.Body.Close()

	// Get a single parameter from the query string
	businessId, _ := strconv.Atoi(id)

	_, errRepo := controller.usecase.Edit(*business.ToUsecase(), businessId)

	if errRepo != nil {
		HandlerMessage := []byte(`{
				"success": false,
				"message": "Error in repo",
			}`)

		utils.ReturnJsonResponse(res, http.StatusCreated, HandlerMessage)
		return
	}

	HandlerMessage := []byte(`{
	 "success": true,
	 "message": "User was successfully edited",
	 }`)

	utils.ReturnJsonResponse(res, http.StatusCreated, HandlerMessage)
	return
}

func (controller *BusinessController) Delete(res http.ResponseWriter, req *http.Request) {
	// check the method
	if req.Method != "DELETE" {
		// Add the response return message
		HandlerMessage := []byte(`{
		"success": false,
		"message": "Check your HTTP method: Invalid HTTP method executed",
		}`)

		utils.ReturnJsonResponse(res, http.StatusMethodNotAllowed, HandlerMessage)
		return
	}

	vars := mux.Vars(req)
	id := vars["id"]
	businessId, _ := strconv.Atoi(id)

	_, errRepo := controller.usecase.Delete(businessId)

	if errRepo != nil {
		fmt.Println(errRepo)
		HandlerMessage := []byte(`{
		"success": false,
		"message": "Business not found",
		}`)

		utils.ReturnJsonResponse(res, http.StatusMethodNotAllowed, HandlerMessage)
		return
	}

	HandlerMessage := []byte(`{
		"success": true,
		"message": "Business was successfully deleted",
	}`)

	utils.ReturnJsonResponse(res, http.StatusMethodNotAllowed, HandlerMessage)
	return
}
