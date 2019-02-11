package order_model

import "time"

type PaymentInfo struct {
	PaymentId   uint  `gorm:"primary_key"`
	OrderId     *uint `gorm:"index"`
	IsPayed     uint
	PaymentType uint
	Price       uint
	DueDate     *time.Time
	PaymentDate *time.Time
}
