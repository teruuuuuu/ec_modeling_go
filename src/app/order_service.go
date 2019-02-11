package app

import (
	"fmt"

	order_repository "./domain/order/repository"
	product_repository "./domain/product/repository"
	order_query "./query/order"
)

func (app *App) UpdateItem(userId uint, productId uint, number uint) bool {
	isFind, product := product_repository.FindById(app.db, productId)
	if !isFind {
		return false
	}
	cart := order_repository.Cart(app.db, &userId)
	cart.UpdateItem(&product.Product.ProductId, product.Product.Price, number)
	tx := app.db.Begin()
	order_repository.Save(app.db, cart)
	tx.Commit()
	fmt.Println("update item endddddd")
	return true
}

func (app *App) CartItems(userId uint) []order_query.ItemsView {
	return order_query.CartItems(app.db, userId)
}
