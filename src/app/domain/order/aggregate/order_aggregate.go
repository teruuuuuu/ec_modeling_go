package payment_aggregate

import (
	"fmt"
	"time"

	order_constant "../constant"
	order_model "../model"
)

type OrderAggregate struct {
	Order            *order_model.Order
	Items            []ItemAggregate
	UsedCoupon       []order_model.UsedCoupon
	PaymentAggregate *PaymentAggregate
}

func (oa *OrderAggregate) UpdateItem(productId *uint, price uint, number uint) bool {
	if oa.Order.OrderStatus != order_constant.Shopping {
		return false
	}

	var items = oa.Items
	fmt.Println(items)
	var item *ItemAggregate = nil
	for i := 0; i < len(items); i++ {
		if *items[i].Item.ProductId == *productId {
			item = &items[i]
			break
		}
	}
	fmt.Println(item)
	if item != nil {
		item.Item.Price = price
		item.Item.Number = number
		item.Item.UpdateDate = time.Now()
		item.CurrentPrice = price
	} else {
		oa.Items = append(oa.Items,
			ItemAggregate{
				Item: &order_model.Item{
					OrderId:    &oa.Order.OrderId,
					ProductId:  productId,
					Price:      price,
					Number:     number,
					UpdateDate: time.Now(),
				},
				CurrentPrice: price,
			})
	}
	return true
}

func (oa *OrderAggregate) AddCoupon(couponId *uint, discount uint) bool {
	if oa.Order.OrderStatus != order_constant.Shopping {
		return false
	}
	oa.UsedCoupon = append(oa.UsedCoupon,
		order_model.UsedCoupon{OrderId: &oa.Order.OrderId, CouponId: couponId, Discount: discount})
	return true
}

func (oa *OrderAggregate) Confirm(paymentType int) bool {
	if len(oa.Items) == 0 {
		return false
	}
	if oa.Order.OrderStatus != order_constant.Shopping {
		return false
	}
	oa.Order.OrderStatus = order_constant.Confirm

	var paymentDetail *PaymentDetail = &PaymentDetail{}
	var paymentAggregate *PaymentAggregate = &PaymentAggregate{}
	paymnetInfo := &order_model.PaymentInfo{OrderId: &oa.Order.OrderId, IsPayed: 0,
		PaymentType: uint(paymentType), Price: oa.price()}
	if oa.PaymentAggregate.PaymentInfo != nil {
		paymnetInfo.IsPayed = oa.PaymentAggregate.PaymentInfo.IsPayed
		paymnetInfo.PaymentId = oa.PaymentAggregate.PaymentInfo.PaymentId
		paymnetInfo.DueDate = oa.PaymentAggregate.PaymentInfo.DueDate
		paymnetInfo.PaymentDate = oa.PaymentAggregate.PaymentInfo.PaymentDate
	} else {
		var dueDate = time.Now().AddDate(0, 0, 7)
		paymnetInfo.IsPayed = 0
		paymnetInfo.DueDate = &dueDate
		paymnetInfo.PaymentDate = nil
	}
	if paymentType == order_constant.CREDIT {
		paymentDetail.CreditPay = &order_model.CreditPay{PaymentId: &paymnetInfo.PaymentId}
	}
	if paymentType == order_constant.BANK {
		paymentDetail.BankPay = &order_model.BankPay{PaymentId: &paymnetInfo.PaymentId, BankAccount: "abcdef"}
	}
	paymentAggregate.PaymentInfo = paymnetInfo
	paymentAggregate.PaymentDetail = paymentDetail
	oa.PaymentAggregate = paymentAggregate

	return true
}

func (oa *OrderAggregate) Payed() bool {
	if oa.Order.OrderStatus != order_constant.Confirm {
		return false
	}
	paymentDate := time.Now()
	oa.PaymentAggregate.PaymentInfo.IsPayed = 1
	oa.PaymentAggregate.PaymentInfo.PaymentDate = &paymentDate
	return true
}

func (oa *OrderAggregate) price() uint {
	var result uint = 0
	for i := 0; i < len(oa.Items); i++ {
		result += oa.Items[i].Sum()
	}
	var discount uint = 0
	for i := 0; i < len(oa.UsedCoupon); i++ {
		discount += oa.UsedCoupon[i].Discount
	}
	if discount > result {
		discount = result
	}
	result -= discount
	return result
}
