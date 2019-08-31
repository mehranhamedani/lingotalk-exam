package utilities

// CustomError struct
type CustomError struct {
	Message string
}

func (customError *CustomError) Error() string {
	return customError.Message
}
