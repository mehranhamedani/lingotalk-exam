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
	if _, ok := stuffMemory[key]; !ok {
		stuffMemory[key] = value
		return nil
	}

	return &utilities.CustomError{Message: "key is exist"}
}

func GetStuffFromMemory(key string) (string, error) {
	if value, ok := stuffMemory[key]; ok {
		stuffMemory[key] = value
		return value, nil
	}

	return "", &utilities.CustomError{Message: "key is not exist"}
}

func SaveStuffToDB(key string, value string) error {
	stuff := model.Stuff{Key: key, Value: value}
	db := GormDB.Create(&stuff)
	return db.Error
}

func GetStuffFromDB(key string) (*model.Stuff, error) {
	stuff := model.Stuff{}
	db := GormDB.First(&stuff, &model.Stuff{Key: key})
	return &stuff, db.Error
}

func SaveStuffToFile(key string, value string) error {
	file, error := ioutil.ReadFile(dbFilePath)
	if error != nil {
		return error
	}
	stuffs := jsonmodel.Stuffs{}
	json.Unmarshal(file, &stuffs)
	stuff := jsonmodel.Stuff{Key: key, Value: value}
	stuffs.Data = append(stuffs.Data, stuff)
	stuffJson, error := json.Marshal(stuffs)
	if error != nil {
		return error
	}
	return ioutil.WriteFile(dbFilePath, stuffJson, 0644)
}

func GetStuffFromFile(key string) (*model.Stuff, error) {
	file, error := ioutil.ReadFile(dbFilePath)
	if error != nil {
		return nil, error
	}
	stuffs := jsonmodel.Stuffs{}
	error = json.Unmarshal(file, &stuffs)
	if error != nil {
		return nil, error
	}
	for _, stuff := range stuffs.Data {
		if stuff.Key == key {
			return &model.Stuff{StuffID: 0, Key: stuff.Key, Value: stuff.Value}, nil
		}
	}
	return nil, &utilities.CustomError{Message: "key is not exist"}
}
