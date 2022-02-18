package models

import "time"

type Booking struct {
	ID                int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	IDUser            int       `gorm:"column:id_user" json:"id_user" valid:"required"`
	IDCab             int       `gorm:"column:id_cab" json:"id_cab" valid:"required"`
	IDFromLocation    int       `gorm:"column:id_from_location" json:"id_from_location" valid:"required"`
	IDToLocation      int       `gorm:"column:id_to_location" json:"id_to_location" valid:"required"`
	CancelledByUser   bool      `gorm:"column:cancelled_by_user" json:"cancelled_by_user" `
	CancelledByDriver bool      `gorm:"column:cancelled_by_driver" json:"cancelled_by_driver" `
	AcceptedByDriver  bool      `gorm:"column:accepted_by_driver" json:"accepted_by_driver" `
	Completed         bool      `gorm:"column:completed" json:"completed" `
	BookingDate       time.Time `gorm:"column:booking_date" json:"booking_date" `
}

// TableName sets the table name
func (b *Booking) TableName() string {
	return "booking"
}
