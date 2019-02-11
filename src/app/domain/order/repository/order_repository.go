package order_repository

import (
	"time"

	"github.com/jinzhu/gorm"

	order_aggregate "../aggregate"
	order_constant "../constant"
	order_model "../model"
)

func DeleteAll(db *gorm.DB) {
	db.Delete(order_model.BankPay{})
	db.Delete(order_model.CreditPay{})
	db.Delete(order_model.Item{})
	db.Delete(order_model.Order{})
	db.Delete(order_model.PaymentInfo{})
	db.Delete(order_model.UsedCoupon{})
}

func Cart(db *gorm.DB, userId *uint) *order_aggregate.OrderAggregate {
	var order *order_model.Order = &order_model.Order{}
	db.Where("user_id = ?", userId).Where("order_status= ?", order_constant.Shopping).
		FirstOrInit(&order, order_model.Order{OrderStatus: order_constant.Shopping, UserId: userId})
	return findByOrder(db, order)
}

func findByOrder(db *gorm.DB, order *order_model.Order) *order_aggregate.OrderAggregate {
	var orderId = &order.OrderId

	var items []order_aggregate.ItemAggregate
	rows, _ := db.Table("items").Select("items.item_id, products.product_id, items.price, items.number, items.update_date, products.price as current_price").
		Joins("left join products on products.product_id = items.product_id").
		Where("items.order_id = ?", orderId).Rows()
	for rows.Next() {
		var itemId uint
		var productId uint
		var price uint
		var number uint
		var updateDate time.Time
		var currentPrice uint
		rows.Scan(&itemId, &productId, &price, &number, &updateDate, &currentPrice)
		// fmt.Println("itemId: " + fmt.Sprint(itemId) + " productId: " + fmt.Sprint(productId) + " price: " +
		// 	fmt.Sprint(price) + " number: " + fmt.Sprint(number) + " updateDate: " + updateDate.String() + " currentPrice: " + fmt.Sprint(currentPrice))
		itemAggregate := order_aggregate.ItemAggregate{
			Item:         &order_model.Item{ItemId: itemId, OrderId: orderId, ProductId: &productId, Price: price, Number: number, UpdateDate: updateDate},
			CurrentPrice: currentPrice}
		items = append(items, itemAggregate)
	}

	var usedCoupons []order_model.UsedCoupon
	db.Where("order_id = ?", orderId).Find(&usedCoupons)

	var paymentInfo *order_model.PaymentInfo
	var bankPay *order_model.BankPay
	var creditPay *order_model.CreditPay
	payment_err := db.Where("order_id = ?", orderId).First(paymentInfo)
	if payment_err != nil {
		paymentInfo = nil
	}
	bank_err := db.Where("order_id = ?", orderId).First(bankPay)
	if bank_err != nil {
		bankPay = nil
	}
	credit_err := db.Where("order_id = ?", orderId).First(creditPay)
	if credit_err != nil {
		creditPay = nil
	}

	paymentAggregate := &order_aggregate.PaymentAggregate{
		PaymentInfo:   paymentInfo,
		PaymentDetail: &order_aggregate.PaymentDetail{BankPay: bankPay, CreditPay: creditPay}}

	return &order_aggregate.OrderAggregate{Order: order, Items: items, UsedCoupon: usedCoupons, PaymentAggregate: paymentAggregate}
}

func Save(db *gorm.DB, oa *order_aggregate.OrderAggregate) {
	db.Save(oa.Order)
	orderId := &oa.Order.OrderId
	for i := 0; i < len(oa.Items); i++ {
		oa.Items[i].Item.OrderId = orderId
		if oa.Items[i].Item.Number == 0 {
			db.Delete(order_model.Item{}, "order_id = ? and product_id = ?", orderId, oa.Items[i].Item.ProductId)
		} else {
			db.Save(oa.Items[i].Item)
		}
	}

	db.Delete(order_model.UsedCoupon{}, "order_id = ?", orderId)
	for i := 0; i < len(oa.UsedCoupon); i++ {
		oa.UsedCoupon[i].OrderId = orderId
		db.Save(oa.UsedCoupon[i])
	}

	if oa.PaymentAggregate.PaymentInfo != nil {
		oa.PaymentAggregate.PaymentInfo.OrderId = orderId
		db.Save(oa.PaymentAggregate.PaymentInfo)
		var paymentId = &oa.PaymentAggregate.PaymentInfo.PaymentId
		var paymentDetail = oa.PaymentAggregate.PaymentDetail
		if paymentDetail.BankPay != nil {
			paymentDetail.BankPay.PaymentId = paymentId
			db.Save(paymentDetail.BankPay)
		}
		if paymentDetail.CreditPay != nil {
			paymentDetail.CreditPay.PaymentId = paymentId
			db.Save(paymentDetail.CreditPay)
		}
	}
}
