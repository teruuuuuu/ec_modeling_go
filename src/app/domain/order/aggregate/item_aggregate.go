package payment_aggregate

import (
	payment_model "../model"
)

type ItemAggregate struct {
	Item         *payment_model.Item
	CurrentPrice uint
}

func (ia *ItemAggregate) Sum() uint {
	return ia.Item.Price * ia.Item.Number
}
