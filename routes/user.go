package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../jsonModel"
	"../resources"
	"../services"
	"../utilities"
)

func setUserRoutes() {
	router.HandleFunc("/user/createUser", createUser).Methods("POST")
	router.HandleFunc("/user/deleteUser", deleteUser).Methods("DELETE")
	router.HandleFunc("/user/updateUser", updateUser).Methods("PUT")
	router.HandleFunc("/user/getAllUsers", getAllUsers).Methods("GET")
}

func createUser(response http.ResponseWriter, request *http.Request) {
	bodyDecoder := json.NewDecoder(request.Body)
	jsonUser := jsonModel.User{}
	error := bodyDecoder.Decode(&jsonUser)
	if error != nil {
		fmt.Println(error.Error())
		utilities.FillHTTPResponse(response, http.StatusBadRequest, true, error.Error(), nil)
		return
	}
	if jsonUser.Name == "" {
		utilities.FillHTTPResponse(response, http.StatusInternalServerError, true, resources.FaStrings.EVSKN, nil)
		return
	}
	error = services.CreateUser(jsonUser.Name)
	if error != nil {
		fmt.Println(error.Error())
		utilities.FillHTTPResponse(response, http.StatusInternalServerError, true, error.Error(), nil)
		return
	}
	utilities.FillHTTPResponse(response, http.StatusOK, false, resources.FaStrings.ABMAS, nil)
}

func updateUser(response http.ResponseWriter, request *http.Request) {
	bodyDecoder := json.NewDecoder(request.Body)
	jsonUser := jsonModel.User{}
	error := bodyDecoder.Decode(&jsonUser)
	if error != nil {
		fmt.Println(error.Error())
		utilities.FillHTTPResponse(response, http.StatusBadRequest, true, error.Error(), nil)
		return
	}
	if jsonUser.UserID <= 0 || jsonUser.Name == "" {
		utilities.FillHTTPResponse(response, http.StatusInternalServerError, true, resources.FaStrings.EVSKN, nil)
		return
	}
	error = services.UpdateUser(jsonUser.UserID, jsonUser.Name)
	if error != nil {
		fmt.Println(error.Error())
		utilities.FillHTTPResponse(response, http.StatusInternalServerError, true, error.Error(), nil)
		return
	}
	utilities.FillHTTPResponse(response, http.StatusOK, false, resources.FaStrings.ABMAS, nil)
}

func deleteUser(response http.ResponseWriter, request *http.Request) {
	bodyDecoder := json.NewDecoder(request.Body)
	jsonUser := jsonModel.User{}
	error := bodyDecoder.Decode(&jsonUser)
	if error != nil {
		fmt.Println(error.Error())
		utilities.FillHTTPResponse(response, http.StatusBadRequest, true, error.Error(), nil)
		return
	}
	services.DeleteUser(jsonUser.UserID)
	utilities.FillHTTPResponse(response, http.StatusOK, false, resources.FaStrings.ABMAS, nil)
}

func getAllUsers(response http.ResponseWriter, request *http.Request) {
	users, error := services.GetAllUsers()
	if error != nil {
		fmt.Println(error.Error())
		utilities.FillHTTPResponse(response, http.StatusInternalServerError, true, error.Error(), nil)
		return
	}
	utilities.FillHTTPResponse(response, http.StatusOK, false, resources.FaStrings.ABMAS, users)
}
