package resources

type adsStatus struct {
	Enable  int
	Disable int
}

type deviceDisplayerAdsStatus struct {
	Enable    int
	Installed int
	Done      int
}

type enums struct {
	AdsStatus                adsStatus
	DeviceDisplayerAdsStatus deviceDisplayerAdsStatus
}

// Enums enums
var Enums = enums{
	AdsStatus:                adsStatus{Enable: 1, Disable: 2},
	DeviceDisplayerAdsStatus: deviceDisplayerAdsStatus{Enable: 1, Installed: 2, Done: 3}}
