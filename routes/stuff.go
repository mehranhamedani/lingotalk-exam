package routes

import (
	"fmt"
	"lingotalk-exam/resources/texts"
	"lingotalk-exam/services"
	"lingotalk-exam/utilities"
	"net/http"

	"github.com/gorilla/mux"
)

func setStuffRoutes() {
	router.HandleFunc("/stuff/save/{key}/{value}/{driver}", save).Methods("POST")
	router.HandleFunc("/stuff/load/{key}/{driver}", load).Methods("GET")
}

func save(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	key := params["key"]
	value := params["value"]
	driver := params["driver"]
	if key == "" || value == "" || driver == "" {
		utilities.FillHTTPResponse(response, http.StatusInternalServerError, true, texts.EN_III, nil)
		return
	}
	error := services.Save(key, value, driver)
	if error != nil {
		fmt.Println(error.Error())
		utilities.FillHTTPResponse(response, http.StatusInternalServerError, true, error.Error(), nil)
		return
	}
	utilities.FillHTTPResponse(response, http.StatusOK, false, texts.EN_DONE, nil)
}

func load(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	key := params["key"]
	driver := params["driver"]
	if key == "" || driver == "" {
		utilities.FillHTTPResponse(response, http.StatusInternalServerError, true, texts.EN_III, nil)
		return
	}
	value, error := services.Load(key, driver)
	if error != nil {
		fmt.Println(error.Error())
		utilities.FillHTTPResponse(response, http.StatusInternalServerError, true, error.Error(), nil)
		return
	}
	utilities.FillHTTPResponse(response, http.StatusOK, false, texts.EN_DONE, value)
}
