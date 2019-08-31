package jsonmodel

// Stuff struct
type Stuff struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Stuffs struct
type Stuffs struct {
	Data []Stuff `json:"data"`
}
