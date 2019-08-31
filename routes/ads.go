package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"../jsonModel"
	"../resources"
	"../services"
	"../utilities"
)

func setAdsRoutes() {
	router.HandleFunc("/ads/createAds", createAds).Methods("POST")
	router.HandleFunc("/ads/deleteAds", deleteAds).Methods("DELETE")
	router.HandleFunc("/ads/updateAds", updateAds).Methods("PUT")
	router.HandleFunc("/ads/getAdsDownloadURL/{googleAdID}/{packageName}", getAdsDownloadURL).Methods("GET")
	router.HandleFunc("/ads/getAllAds", getAllAds).Methods("GET")
	router.HandleFunc("/ads/adsInstalled", adsInstalled).Methods("PUT")
	router.HandleFunc("/ads/adsDone", adsDone).Methods("PUT")
	router.HandleFunc("/ads/getAllInstalledAds", getAllInstalledAds).Methods("GET")
}

func createAds(response http.ResponseWriter, request *http.Request) {
	bodyDecoder := json.NewDecoder(request.Body)
	jsonAds := jsonModel.Ads{}
	error := bodyDecoder.Decode(&jsonAds)
	if error != nil {
		fmt.Println(error.Error())
		utilities.FillHTTPResponse(response, http.StatusBadRequest, true, error.Error(), nil)
		return
	}
	if jsonAds.PackageName == "" || jsonAds.UserID <= 0 || jsonAds.Title == "" || jsonAds.Description == "" {
		utilities.FillHTTPResponse(response, http.StatusInternalServerError, true, resources.FaStrings.EVSKN, nil)
		return
	}
	error = services.CreateAds(jsonAds.PackageName, jsonAds.UserID, jsonAds.Title, jsonAds.Description)
	if error != nil {
		fmt.Println(error.Error())
		utilities.FillHTTPResponse(response, http.StatusInternalServerError, true, error.Error(), nil)
		return
	}
	utilities.FillHTTPResponse(response, http.StatusOK, false, resources.FaStrings.ABMAS, nil)
}

func deleteAds(response http.ResponseWriter, request *http.Request) {
	bodyDecoder := json.NewDecoder(request.Body)
	jsonAds := jsonModel.Ads{}
	error := bodyDecoder.Decode(&jsonAds)
	if error != nil {
		fmt.Println(error.Error())
		utilities.FillHTTPResponse(response, http.StatusBadRequest, true, error.Error(), nil)
		return
	}
	error = services.DeleteAds(jsonAds.AdsID)
	if error != nil {
		fmt.Println(error.Error())
		utilities.FillHTTPResponse(response, http.StatusInternalServerError, true, error.Error(), nil)
		return
	}
	utilities.FillHTTPResponse(response, http.StatusOK, false, resources.FaStrings.ABMAS, nil)
}

func updateAds(response http.ResponseWriter, request *http.Request) {
	bodyDecoder := json.NewDecoder(request.Body)
	jsonAds := jsonModel.Ads{}
	error := bodyDecoder.Decode(&jsonAds)
	if error != nil {
		fmt.Println(error.Error())
		utilities.FillHTTPResponse(response, http.StatusBadRequest, true, error.Error(), nil)
		return
	}
	if jsonAds.AdsID <= 0 || jsonAds.PackageName == "" || jsonAds.UserID <= 0 {
		utilities.FillHTTPResponse(response, http.StatusInternalServerError, true, resources.FaStrings.EVSKN, nil)
		return
	}
	error = services.UpdateAds(jsonAds.AdsID, jsonAds.PackageName, jsonAds.UserID)
	if error != nil {
		fmt.Println(error.Error())
		utilities.FillHTTPResponse(response, http.StatusInternalServerError, true, error.Error(), nil)
		return
	}
	utilities.FillHTTPResponse(response, http.StatusOK, false, resources.FaStrings.ABMAS, nil)
}

func getAdsDownloadURL(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	googleAdID := params["googleAdID"]
	packageName := params["packageName"]
	if googleAdID == "" || packageName == "" {
		utilities.FillHTTPResponse(response, http.StatusInternalServerError, true, resources.FaStrings.EVSKN, nil)
		return
	}
	data, error := services.GetAdsDownloadURL(googleAdID, packageName)
	if error != nil {
		fmt.Println(error.Error())
		utilities.FillHTTPResponse(response, http.StatusInternalServerError, true, error.Error(), nil)
		return
	}
	utilities.FillHTTPResponse(response, http.StatusOK, false, resources.FaStrings.ABMAS, data)
}

func getAllAds(response http.ResponseWriter, request *http.Request) {
	ads, error := services.GetAllAds()
	if error != nil {
		fmt.Println(error.Error())
		utilities.FillHTTPResponse(response, http.StatusInternalServerError, true, error.Error(), nil)
		return
	}
	utilities.FillHTTPResponse(response, http.StatusOK, false, resources.FaStrings.ABMAS, ads)
}

func adsInstalled(response http.ResponseWriter, request *http.Request) {
	bodyDecoder := json.NewDecoder(request.Body)
	adsInstalledRequest := jsonModel.AdsInstalledRequest{}
	error := bodyDecoder.Decode(&adsInstalledRequest)
	if error != nil {
		fmt.Println(error.Error())
		utilities.FillHTTPResponse(response, http.StatusBadRequest, true, error.Error(), nil)
		return
	}
	if adsInstalledRequest.GoogleAdID == "" || adsInstalledRequest.AdsID == "" {
		utilities.FillHTTPResponse(response, http.StatusInternalServerError, true, resources.FaStrings.EVSKN, nil)
		return
	}
	error = services.AdsInstalled(adsInstalledRequest.GoogleAdID, adsInstalledRequest.AdsID)
	if error != nil {
		fmt.Println(error.Error())
		utilities.FillHTTPResponse(response, http.StatusInternalServerError, true, error.Error(), nil)
		return
	}
	utilities.FillHTTPResponse(response, http.StatusOK, false, resources.FaStrings.ABMAS, nil)
}

func adsDone(response http.ResponseWriter, request *http.Request) {
	bodyDecoder := json.NewDecoder(request.Body)
	adsDoneRequest := jsonModel.AdsDoneRequest{}
	error := bodyDecoder.Decode(&adsDoneRequest)
	if error != nil {
		fmt.Println(error.Error())
		utilities.FillHTTPResponse(response, http.StatusBadRequest, true, error.Error(), nil)
		return
	}
	if adsDoneRequest.GoogleAdID == "" || adsDoneRequest.AdsID == "" {
		utilities.FillHTTPResponse(response, http.StatusInternalServerError, true, resources.FaStrings.EVSKN, nil)
		return
	}
	error = services.AdsDone(adsDoneRequest.GoogleAdID, adsDoneRequest.AdsID)
	if error != nil {
		fmt.Println(error.Error())
		utilities.FillHTTPResponse(response, http.StatusInternalServerError, true, error.Error(), nil)
		return
	}
	utilities.FillHTTPResponse(response, http.StatusOK, false, resources.FaStrings.ABMAS, nil)
}

func getAllInstalledAds(response http.ResponseWriter, request *http.Request) {
	bodyDecoder := json.NewDecoder(request.Body)
	getAllInstalledAdsRequest := jsonModel.GetAllInstalledAdsRequest{}
	error := bodyDecoder.Decode(&getAllInstalledAdsRequest)
	if error != nil {
		fmt.Println(error.Error())
		utilities.FillHTTPResponse(response, http.StatusBadRequest, true, error.Error(), nil)
		return
	}
	if getAllInstalledAdsRequest.GoogleAdID == "" || getAllInstalledAdsRequest.PackageName == "" {
		utilities.FillHTTPResponse(response, http.StatusInternalServerError, true, resources.FaStrings.EVSKN, nil)
		return
	}
	data, error := services.GetAllInstalledAds(getAllInstalledAdsRequest.GoogleAdID, getAllInstalledAdsRequest.PackageName)
	if error != nil {
		fmt.Println(error.Error())
		utilities.FillHTTPResponse(response, http.StatusInternalServerError, true, error.Error(), nil)
		return
	}
	utilities.FillHTTPResponse(response, http.StatusOK, false, resources.FaStrings.ABMAS, data)
}
