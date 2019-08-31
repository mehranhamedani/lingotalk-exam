package services

import (
	"errors"

	"../dal"
	"../jsonModel"
	"../model"
	"../resources"
	"github.com/speps/go-hashids"
)

// CreateAds func
func CreateAds(packageName string, userID int, title string, description string) error {
	var isUserExist bool
	var isAdsExist bool
	var error error
	isUserExist, error = dal.GetIsUserExistByID(userID)
	if error != nil {
		return error
	}
	if !isUserExist {
		return errors.New(resources.FaStrings.KMPN)
	}
	isAdsExist, error = dal.GetIsAdsExistByPackageName(packageName)
	if error != nil {
		return error
	}
	if isAdsExist {
		return errors.New(resources.FaStrings.YTFBINBVD)
	}
	error = dal.CreateAds(packageName, userID, title, description)
	return error
}

// UpdateAds func
func UpdateAds(adsID int, packageName string, userID int) error {
	isUserExist, error := dal.GetIsUserExistByID(userID)
	if error != nil {
		return error
	}
	if !isUserExist {
		return errors.New(resources.FaStrings.KMPN)
	}
	error = dal.UpdateAds(adsID, packageName, userID)
	return error
}

// DeleteAds func
func DeleteAds(adsID int) error {
	error := dal.DeleteAds(adsID)
	return error
}

// GetAdsDownloadURL func
func GetAdsDownloadURL(googleAdID string, packageName string) (*jsonModel.GetAdsDownloadURLResponse, error) {
	getAdsDownloadURLResponse := jsonModel.GetAdsDownloadURLResponse{}
	var error error
	deviceDisplayer, error := dal.GetOrCreateDeviceDisplayer(googleAdID, packageName)
	if error != nil {
		return nil, error
	}
	ads, deviceDisplayerAds, error := dal.GetDeviceDisplayerAvailableAds(deviceDisplayer.DeviceDisplayerID)
	if error != nil {
		return nil, error
	}
	if ads != nil && deviceDisplayerAds != nil {
		getAdsDownloadURLResponse.FromDeviceDisplayerAdsAndAds(deviceDisplayerAds, ads)
	}
	return &getAdsDownloadURLResponse, error
}

// GetAllAds func
func GetAllAds() ([]model.Ads, error) {
	return dal.GetAllAds()
}

// AdsInstalled func
func AdsInstalled(googleAdID string, hashedAdsID string) error {
	hd := hashids.NewData()
	hd.Salt = resources.DownloadURLAdsIDSalt
	hd.MinLength = resources.DownloadURLAdsIDMinLength
	h, _ := hashids.NewWithData(hd)
	adsIDs, error := h.DecodeWithError(hashedAdsID)
	if error != nil {
		return error
	}
	if len(adsIDs) < 1 {
		return errors.New(resources.FaStrings.SVSMN)
	}
	deviceDisplayerAdsID := adsIDs[0]
	deviceDisplayerAds, error := dal.GetDeviceDisplayerAdsByID(deviceDisplayerAdsID)
	if error != nil {
		return error
	}
	if deviceDisplayerAds.Status != resources.Enums.DeviceDisplayerAdsStatus.Enable {
		return errors.New(resources.FaStrings.SVSMN)
	}
	return dal.UpdateDeviceDisplayerAdsStatus(adsIDs[0], resources.Enums.DeviceDisplayerAdsStatus.Installed)
}

// AdsDone func
func AdsDone(googleAdID string, hashedAdsID string) error {
	hd := hashids.NewData()
	hd.Salt = resources.DownloadURLAdsIDSalt
	hd.MinLength = resources.DownloadURLAdsIDMinLength
	h, _ := hashids.NewWithData(hd)
	adsIDs, error := h.DecodeWithError(hashedAdsID)
	if error != nil {
		return error
	}
	if len(adsIDs) < 1 {
		return errors.New(resources.FaStrings.SVSMN)
	}
	deviceDisplayerAdsID := adsIDs[0]
	deviceDisplayerAds, error := dal.GetDeviceDisplayerAdsByID(deviceDisplayerAdsID)
	if error != nil {
		return error
	}
	if deviceDisplayerAds.Status != resources.Enums.DeviceDisplayerAdsStatus.Installed {
		return errors.New(resources.FaStrings.AMNNNA)
	}
	return dal.UpdateDeviceDisplayerAdsStatus(adsIDs[0], resources.Enums.DeviceDisplayerAdsStatus.Done)
}

// GetAllInstalledAds func
func GetAllInstalledAds(googleAdID string, packageName string) (*[]jsonModel.GetAdsDownloadURLResponse, error) {
	var getAdsDownloadURLResponses []jsonModel.GetAdsDownloadURLResponse
	deviceDisplayer, error := dal.GetDeviceDisplayer(googleAdID, packageName)
	if error != nil {
		return nil, error
	}
	deviceDisplayerAds, error := dal.GetDeviceDisplayerAds(deviceDisplayer.DeviceDisplayerID, resources.Enums.DeviceDisplayerAdsStatus.Installed)
	if error != nil {
		return nil, error
	}
	for index := 0; index < len(*deviceDisplayerAds); index++ {
		getAdsDownloadURLResponse := jsonModel.GetAdsDownloadURLResponse{}
		deviceDisplayerAd := (*deviceDisplayerAds)[index]
		getAdsDownloadURLResponse.FromDeviceDisplayerAdsAndAds(&deviceDisplayerAd, &deviceDisplayerAd.Ads)
		getAdsDownloadURLResponses = append(getAdsDownloadURLResponses, getAdsDownloadURLResponse)
	}
	return &getAdsDownloadURLResponses, error
}
