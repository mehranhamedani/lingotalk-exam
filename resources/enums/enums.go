package enums

type DriverType int

// ToString func
func (driverType DriverType) ToString() string {
	result := ""
	switch driverType {
	case DriverTypeFile:
		result = "file"
	case DriverTypeDB:
		result = "database"
	case DriverTypeMemory:
		result = "memory"
	}
	return result
}

func GetDriverTypeByString(value string) DriverType {
	result := DriverTypeFile
	switch value {
	case DriverTypeFile.ToString():
		result = DriverTypeFile
	case DriverTypeDB.ToString():
		result = DriverTypeDB
	case DriverTypeMemory.ToString():
		result = DriverTypeMemory
	}
	return result
}

const (
	DriverTypeFile   DriverType = 1
	DriverTypeDB     DriverType = 2
	DriverTypeMemory DriverType = 3
)
