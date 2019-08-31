package services

import (
	"testing"

	"../dal"
)

func TestGetAdsDownloadURL(t *testing.T) {
	dal.OpenDB()
	dal.AutoMigrateDB()
	ads, error := GetAdsDownloadURL("0123456789", "com.google.test2")
	if error != nil {
		t.Error("", ads)
	}
}
