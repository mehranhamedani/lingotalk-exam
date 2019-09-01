package services

import (
	"lingotalk-exam/resources/enums"
	"lingotalk-exam/utilities"
	"math/rand"
	"testing"
)

// TestSave func
func TestSave(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("panic error:%v", r)
		}
	}()

	stuffMocks := utilities.GetStuffMocks()

	for index := 0; index < len(stuffMocks); index++ {
		key := stuffMocks[rand.Intn(len(stuffMocks))]
		value := stuffMocks[rand.Intn(len(stuffMocks))]
		error := Save(key, value, enums.DriverTypeDB.ToString())
		if error != nil {
			t.Errorf(error.Error())
		}
		error = Save(key, value, enums.DriverTypeFile.ToString())
		if error != nil {
			t.Errorf(error.Error())
		}
		error = Save(key, value, enums.DriverTypeMemory.ToString())
		if error != nil {
			t.Errorf(error.Error())
		}
	}
}

// TestLoad func
func TestLoad(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("panic error:%v", r)
		}
	}()

	stuffMocks := utilities.GetStuffMocks()

	for index := 0; index < len(stuffMocks); index++ {
		key := stuffMocks[rand.Intn(len(stuffMocks))]
		_, error := Load(key, enums.DriverTypeDB.ToString())
		if error != nil {
			t.Errorf(error.Error())
		}
		_, error = Load(key, enums.DriverTypeFile.ToString())
		if error != nil {
			t.Errorf(error.Error())
		}
		_, error = Load(key, enums.DriverTypeMemory.ToString())
		if error != nil {
			t.Errorf(error.Error())
		}
	}
}
