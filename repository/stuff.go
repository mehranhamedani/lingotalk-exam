package repository

import (
	"encoding/json"
	"io/ioutil"
	"lingotalk-exam/jsonmodel"
	"lingotalk-exam/model"
	"lingotalk-exam/utilities"
	"path"
)

var stuffMemory map[string]string
var dbFilePath string

func init() {
	stuffMemory = make(map[string]string)
	dbFilePath = path.Join(utilities.GetRootPath(), "db/db.json")
}

func SaveStuffToMemory(key string, value string) error {
	_, error := GetStuffFromMemory(key)
	if error == nil {
		return &utilities.CustomError{Message: "key is exist"}
	}
	stuffMemory[key] = value
	return nil
}

func GetStuffFromMemory(key string) (string, error) {
	if value, ok := stuffMemory[key]; ok {
		return value, nil
	}

	return "", &utilities.CustomError{Message: "key does not exist"}
}

func SaveStuffToDB(key string, value string) error {
	_, error := GetStuffFromDB(key)
	if error == nil {
		return &utilities.CustomError{Message: "key is exist"}
	}
	stuff := model.Stuff{Key: key, Value: value}
	db := GormDB.Create(&stuff)
	return db.Error
}

func GetStuffFromDB(key string) (*model.Stuff, error) {
	var error error
	stuff := model.Stuff{}
	db := GormDB.First(&stuff, &model.Stuff{Key: key})
	error = db.Error
	if error != nil && error.Error() == "record not found" {
		error = &utilities.CustomError{Message: "key does not exist"}
	}
	return &stuff, error
}

func SaveStuffToFile(key string, value string) error {
	file, error := ioutil.ReadFile(dbFilePath)
	if error != nil {
		return error
	}
	stuffs := jsonmodel.Stuffs{}
	json.Unmarshal(file, &stuffs)
	_, error = getStuffByKey(key, stuffs.Data)
	if error == nil {
		return &utilities.CustomError{Message: "key is exist"}
	}
	stuff := jsonmodel.Stuff{Key: key, Value: value}
	stuffs.Data = append(stuffs.Data, stuff)
	stuffJson, error := json.Marshal(stuffs)
	if error != nil {
		return error
	}
	return ioutil.WriteFile(dbFilePath, stuffJson, 0644)
}

func GetStuffFromFile(key string) (*model.Stuff, error) {
	var stuff model.Stuff
	file, error := ioutil.ReadFile(dbFilePath)
	if error != nil {
		return nil, error
	}
	stuffs := jsonmodel.Stuffs{}
	error = json.Unmarshal(file, &stuffs)
	if error != nil {
		return nil, error
	}
	jsonStuff, error := getStuffByKey(key, stuffs.Data)
	if error == nil {
		stuff = model.Stuff{Key: jsonStuff.Key, Value: jsonStuff.Value}
	}
	return &stuff, error
}

func getStuffByKey(key string, stuffs []jsonmodel.Stuff) (*jsonmodel.Stuff, error) {
	for _, stuff := range stuffs {
		if stuff.Key == key {
			return &jsonmodel.Stuff{Key: stuff.Key, Value: stuff.Value}, nil
		}
	}
	return nil, &utilities.CustomError{Message: "key does not exist"}
}
