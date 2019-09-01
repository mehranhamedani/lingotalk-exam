package services

// Save func
func Save(key string, value string, driver string) error {
	helper := getHelperByDriver(driver)
	return helper.Save(key, value)
}

// Load func
func Load(key string, driver string) (string, error) {
	helper := getHelperByDriver(driver)
	return helper.Load(key)
}
