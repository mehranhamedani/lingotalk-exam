package repository

import (
	"lingotalk-exam/utilities"
	"math/rand"
	"testing"
)

func TestSaveStuffToMemory(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("panic error:%v", r)
		}
	}()

	newStuffKeys := utilities.GetStuffMocks()
	newStuffValues := utilities.GetStuffMocks()
	for index := 0; index < len(newStuffKeys); index++ {
		newStuffKey := newStuffKeys[rand.Intn(len(newStuffKeys))]
		newStuffValue := newStuffValues[rand.Intn(len(newStuffValues))]
		error := SaveStuffToMemory(newStuffKey, newStuffValue)
		if error != nil {
			t.Errorf(error.Error())
		}
	}
}

func TestGetStuffFromMemory(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("panic error:%v", r)
		}
	}()

	newStuffKeys := utilities.GetStuffMocks()
	for index := 0; index < len(newStuffKeys); index++ {
		newStuffKey := newStuffKeys[rand.Intn(len(newStuffKeys))]
		_, error := GetStuffFromMemory(newStuffKey)
		if error != nil {
			t.Errorf(error.Error())
		}
	}
}

func TestSaveStuffToDB(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("panic error:%v", r)
		}
	}()

	newStuffKeys := utilities.GetStuffMocks()
	newStuffValues := utilities.GetStuffMocks()
	for index := 0; index < len(newStuffKeys); index++ {
		newStuffKey := newStuffKeys[rand.Intn(len(newStuffKeys))]
		newStuffValue := newStuffValues[rand.Intn(len(newStuffValues))]
		error := SaveStuffToDB(newStuffKey, newStuffValue)
		if error != nil {
			t.Errorf(error.Error())
		}
	}
}

func TestGetStuffFromDB(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("panic error:%v", r)
		}
	}()

	newStuffKeys := utilities.GetStuffMocks()
	for index := 0; index < len(newStuffKeys); index++ {
		newStuffKey := newStuffKeys[rand.Intn(len(newStuffKeys))]
		_, error := GetStuffFromDB(newStuffKey)
		if error != nil {
			t.Errorf(error.Error())
		}
	}
}

func TestSaveStuffToFile(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("panic error:%v", r)
		}
	}()

	newStuffKeys := utilities.GetStuffMocks()
	newStuffValues := utilities.GetStuffMocks()
	for index := 0; index < len(newStuffKeys); index++ {
		newStuffKey := newStuffKeys[rand.Intn(len(newStuffKeys))]
		newStuffValue := newStuffValues[rand.Intn(len(newStuffValues))]
		error := SaveStuffToFile(newStuffKey, newStuffValue)
		if error != nil {
			t.Errorf(error.Error())
		}
	}
}

func TestGetStuffFromFile(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("panic error:%v", r)
		}
	}()

	newStuffKeys := utilities.GetStuffMocks()
	for index := 0; index < len(newStuffKeys); index++ {
		newStuffKey := newStuffKeys[rand.Intn(len(newStuffKeys))]
		_, error := GetStuffFromFile(newStuffKey)
		if error != nil {
			t.Errorf(error.Error())
		}
	}
}
