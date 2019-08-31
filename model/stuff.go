package model

// Stuff struct
type Stuff struct {
	StuffID int    `gorm:"column:StuffID;primary_key:yes;auto_increment=true"`
	Key     string `gorm:"column:Key"`
	Value   string `gorm:"column:Value"`
}

// TableName func
func (Stuff) TableName() string {
	return "Stuff"
}
