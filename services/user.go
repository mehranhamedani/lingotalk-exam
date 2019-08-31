package services

import (
	"errors"

	"../dal"
	"../model"
	"../resources"
)

// CreateUser func
func CreateUser(name string) error {
	isUserExist, error := dal.GetIsUserExistByName(name)
	if error != nil {
		return error
	}
	if isUserExist {
		return errors.New(resources.FaStrings.INKVD)
	}
	error = dal.CreateUser(name)
	return error
}

// UpdateUser func
func UpdateUser(userID int, name string) error {
	return dal.UpdateUser(userID, name)
}

// DeleteUser func
func DeleteUser(userID int) error {
	return dal.DeleteUser(userID)
}

// GetAllUsers func
func GetAllUsers() ([]model.User, error) {
	return dal.GetAllUsers()
}
