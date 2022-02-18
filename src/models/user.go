package models

type User struct {
	ID       int `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name     string
	Email    string
	Phone    string
	Username string `gorm:"column:username;unique" json:"username" valid:"required"`
	Password string `gorm:"column:password;not null" json:"password" valid:"required"`
}

// TableName sets the table name
func (u *User) TableName() string {
	return "user"
}
