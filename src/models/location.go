package models

type Location struct {
	ID        int `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Latitude  string
	Longitude string
}

// TableName sets the table name
func (l *Location) TableName() string {
	return "location"
}
