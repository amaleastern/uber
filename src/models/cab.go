package models

type Cab struct {
	ID         int `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	NoOfSeats  int
	DriverName string
	CarModel   string
	IDLocation int
	Active     bool
}

// TableName sets the table name
func (c *Cab) TableName() string {
	return "cab"
}
