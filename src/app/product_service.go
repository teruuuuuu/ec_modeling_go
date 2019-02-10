package app

import (
	product_query "./query/product"
)

func (app *App) ShowProducts() []product_query.ProductView {
	return product_query.Search(app.db, "")
}
