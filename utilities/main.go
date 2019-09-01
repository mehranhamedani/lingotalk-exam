package utilities

import (
	"net/http"
	"os"
	"path"

	"encoding/json"

	"lingotalk-exam/jsonmodel"
)

// FillHTTPResponse func
func FillHTTPResponse(responseWriter http.ResponseWriter, status int, hasError bool, message string, data interface{}) {
	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(status)
	apiResponse := jsonmodel.APIResponse{HasError: hasError, Message: message, Data: data}
	json.NewEncoder(responseWriter).Encode(apiResponse)
}

// GetRootPath func
func GetRootPath() string {
	gp := os.Getenv("GOPATH")
	return path.Join(gp, "src/lingotalk-exam")
}

// GetStuffMocks func
func GetStuffMocks() []string {
	return []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
}
