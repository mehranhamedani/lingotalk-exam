package model

// JSONConfig struct
type JSONConfig struct {
	Host string       `json:"host"`
	Port string       `json:"port"`
	DB   JSONConfigDB `json:"db"`
}
