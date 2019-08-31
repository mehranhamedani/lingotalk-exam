package resources

type faString struct {
	INKVD, ABMAS, KMPN, EVSKN, AMPN, YTFBINBVD, SVSMN, AMNFN, AMNNNA string
}

type enString struct {
	RNF string
}

// FaStrings resource
var FaStrings = faString{
	INKVD:     "این نام کاربری وجود دارد",
	ABMAS:     "عملیات با موفقیت انجام شد",
	KMPN:      "کاربر موردنظر پیدا نشد",
	EVSKN:     "اطلاعات وارد شده کامل نیست",
	AMPN:      "آگهی موردنظر پیدا نشد",
	YTFBINBVD: "یک تبلیغ فعال با این نام‌بسته وجود دارد",
	SVSMN:     "شناسه وارد شده معتبر نیست",
	AMNFN:     "آگهی موردنظرفعال نیست",
	AMNNNA:    "آگهی موردنظر نصب نشده است"}

// EnStrings resource
var EnStrings = enString{
	RNF: "record not found"}

// DownloadURLAdsIDSalt salt
var DownloadURLAdsIDSalt = "AdsIDSalt"

// DownloadURLAdsIDMinLength length
var DownloadURLAdsIDMinLength = 6
