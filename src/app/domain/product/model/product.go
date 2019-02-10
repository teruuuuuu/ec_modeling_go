package product_model

type Product struct {
	ProductId uint `gorm:"primary_key"`
	Name      string
	Price     uint
}
