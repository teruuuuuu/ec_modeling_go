package payment_aggregate

import (
	payment_model "../model"
)

type PaymentDetail struct {
	BankPay   *payment_model.BankPay
	CreditPay *payment_model.CreditPay
}

type PaymentAggregate struct {
	PaymentInfo   *payment_model.PaymentInfo
	PaymentDetail *PaymentDetail
}
