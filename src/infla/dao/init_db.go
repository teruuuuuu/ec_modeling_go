package dao

import (
	coupon_model "../../app/domain/coupon/model"
	order_model "../../app/domain/order/model"
	product_model "../../app/domain/product/model"
	user_model "../../app/domain/user/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initDb(db *gorm.DB) {
	db.AutoMigrate(&coupon_model.Coupon{})
	db.AutoMigrate(&order_model.BankPay{})
	db.AutoMigrate(&order_model.CreditPay{})
	db.AutoMigrate(&order_model.BankPay{})
	db.AutoMigrate(&order_model.Item{})
	db.AutoMigrate(&order_model.Order{})
	db.AutoMigrate(&order_model.PaymentInfo{})
	db.AutoMigrate(&order_model.UsedCoupon{})
	db.AutoMigrate(&product_model.Product{})
	db.AutoMigrate(&product_model.ProductInfo{})
	db.AutoMigrate(&user_model.User{})
	db.AutoMigrate(&user_model.UserInfo{})
}
