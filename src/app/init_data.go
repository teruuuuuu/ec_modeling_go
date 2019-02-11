package app

import (
	"time"

	"github.com/jinzhu/gorm"

	coupon_model "./domain/coupon/model"
	coupon_repository "./domain/coupon/repository"
	order_aggregate "./domain/order/aggregate"
	order_constant "./domain/order/constant"
	order_model "./domain/order/model"
	order_repository "./domain/order/repository"
	product_aggregate "./domain/product/aggregate"
	product_model "./domain/product/model"
	product_repository "./domain/product/repository"
	user_aggregate "./domain/user/aggregate"
	user_model "./domain/user/model"
	user_repository "./domain/user/repository"
)

func initData(db *gorm.DB) {
	tx := db.Begin()
	userInit(tx)
	productInit(tx)
	couponInit(tx)
	orderInit(tx)
	tx.Commit()
}

func couponInit(db *gorm.DB) {
	coupon_repository.DeleteAll(db)
	coupon_repository.Save(db, &coupon_model.Coupon{CouponNumber: "1234567890", Discount: 1000})
	coupon_repository.Save(db, &coupon_model.Coupon{CouponNumber: "2345678901", Discount: 2000})
	coupon_repository.Save(db, &coupon_model.Coupon{CouponNumber: "3456789012", Discount: 3000})
	coupon_repository.Save(db, &coupon_model.Coupon{CouponNumber: "4567890123", Discount: 4000})
	coupon_repository.Save(db, &coupon_model.Coupon{CouponNumber: "5678901234", Discount: 5000})
	coupon_repository.Save(db, &coupon_model.Coupon{CouponNumber: "6789012345", Discount: 6000})
}

func productInit(db *gorm.DB) {
	product_repository.DeleteAll(db)
	product_repository.Save(db, makeProductAggregate("product1", 100, "product1_description"))
	product_repository.Save(db, makeProductAggregate("product2", 200, "product2_description"))
	product_repository.Save(db, makeProductAggregate("product3", 300, "product3_description"))
	product_repository.Save(db, makeProductAggregate("product4", 400, "product4_description"))
	product_repository.Save(db, makeProductAggregate("product5", 500, "product5_description"))
	product_repository.Save(db, makeProductAggregate("product6", 600, "product6_description"))
	product_repository.Save(db, makeProductAggregate("product7", 700, "product7_description"))
	product_repository.Save(db, makeProductAggregate("product8", 800, "product8_description"))
	product_repository.Save(db, makeProductAggregate("product9", 900, "product9_description"))
	product_repository.Save(db, makeProductAggregate("product10", 1000, "product10_description"))
}

func makeProductAggregate(name string, price uint, description string) *product_aggregate.ProductAggregate {
	var product = product_model.Product{Name: name, Price: price}
	var productInfo = product_model.ProductInfo{Description: description}
	return &product_aggregate.ProductAggregate{Product: &product, ProductInfo: &productInfo}
}

func userInit(db *gorm.DB) {
	user_repository.DeleteAll(db)
	user_repository.Save(db, makeUserAggregate("user1", "password", "user1_address", "user1_postal_code"))
	user_repository.Save(db, makeUserAggregate("user2", "password", "user2_address", "user2_postal_code"))
	user_repository.Save(db, makeUserAggregate("user3", "password", "user3_address", "user3_postal_code"))
	user_repository.Save(db, makeUserAggregate("user4", "password", "user4_address", "user4_postal_code"))
	user_repository.Save(db, makeUserAggregate("user5", "password", "user5_address", "user5_postal_code"))
}

func makeUserAggregate(name string, password string, address string, postalCode string) *user_aggregate.UserAggregate {
	var user = user_model.User{Name: name, Password: password}
	var userInfo = user_model.UserInfo{Address: address, PostalCode: postalCode}
	return &user_aggregate.UserAggregate{User: user, UserInfo: userInfo}
}

func orderInit(db *gorm.DB) {
	order_repository.DeleteAll(db)

	var user1 user_model.User
	db.First(&user1, "name = ?", "user1")

	var products []product_model.Product
	db.Find(&products)
	var coupons []coupon_model.Coupon
	db.Find(&coupons)

	order1 := order_repository.Cart(db, &user1.UserId)
	order1.UpdateItem(&products[0].ProductId, products[0].Price, 2)
	order1.UpdateItem(&products[0].ProductId, products[0].Price, 3)
	order1.UpdateItem(&products[1].ProductId, products[1].Price, 4)
	order1.UpdateItem(&products[2].ProductId, products[2].Price, 1)
	order1.UpdateItem(&products[1].ProductId, products[1].Price, 0)
	order1.Confirm(order_constant.BANK)
	order1.Payed()
	order_repository.Save(db, order1)

	var user2 user_model.User
	db.First(&user2, "name = ?", "user2")
	order2 := order_repository.Cart(db, &user2.UserId)
	order2.UpdateItem(&products[2].ProductId, products[2].Price, 1)
	order2.UpdateItem(&products[3].ProductId, products[3].Price, 4)
	order2.UpdateItem(&products[4].ProductId, products[4].Price, 8)
	order2.AddCoupon(&coupons[2].CouponId, coupons[2].Discount)
	order2.Confirm(order_constant.CREDIT)
	order_repository.Save(db, order2)

	var user3 user_model.User
	db.First(&user3, "name = ?", "user3")
	order3 := order_repository.Cart(db, &user3.UserId)
	order3.UpdateItem(&products[0].ProductId, products[0].Price, 1)
	order3.UpdateItem(&products[1].ProductId, products[1].Price, 2)
	order_repository.Save(db, order3)
}

func makeOrderAggregate(user user_model.User, Items []order_aggregate.ItemAggregate, UsedCoupons []order_model.UsedCoupon) *order_aggregate.OrderAggregate {

	var Order = order_model.Order{OrderStatus: order_constant.Shopping, UserId: &user.UserId}
	return &order_aggregate.OrderAggregate{Order: &Order, Items: Items, UsedCoupon: UsedCoupons}
}

func makeItemAggregate(product product_model.Product, number uint) order_aggregate.ItemAggregate {
	var item = order_model.Item{ProductId: &product.ProductId, Price: product.Price, Number: number, UpdateDate: time.Now()}
	return order_aggregate.ItemAggregate{Item: &item, CurrentPrice: product.Price}
}
