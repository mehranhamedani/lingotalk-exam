package services

import (
	"fmt"
	"lingotalk-exam/repository"
	"lingotalk-exam/resources/enums"
)

type IHelper interface {
	Save(key string, value string) error
	Load(key string) (string, error)
}

type FileHelper struct {
}

func (fileHelper FileHelper) Save(key string, value string) error {
	return repository.SaveStuffToFile(key, value)
}

func (fileHelper FileHelper) Load(key string) (string, error) {
	stuff, error := repository.GetStuffFromFile(key)
	return stuff.Value, error
}

type MemoryHelper struct {
}

func (memoryHelper MemoryHelper) Save(key string, value string) error {
	return repository.SaveStuffToMemory(key, value)
}

func (memoryHelper MemoryHelper) Load(key string) (string, error) {
	return repository.GetStuffFromMemory(key)
}

type DBHelper struct {
}

func (dbHelper DBHelper) Save(key string, value string) error {
	return repository.SaveStuffToDB(key, value)
}

func (dbHelper DBHelper) Load(key string) (string, error) {
	stuff, error := repository.GetStuffFromDB(key)
	return stuff.Value, error
}

func init() {
	fmt.Println("Services init")
}

func getHelperByDriver(driver string) IHelper {
	var helper IHelper
	var driverType = enums.GetDriverTypeByString(driver)
	switch driverType {
	case enums.DriverTypeDB:
		helper = DBHelper{}
	case enums.DriverTypeFile:
		helper = FileHelper{}
	case enums.DriverTypeMemory:
		helper = MemoryHelper{}
	}
	return helper
}
