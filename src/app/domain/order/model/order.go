package order_model

type Order struct {
	OrderId     uint `gorm:"primary_key"`
	OrderStatus uint
	UserId      uint
}
