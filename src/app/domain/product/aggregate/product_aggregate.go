package product_aggregate

import product_model "../model"

type ProductAggregate struct {
	Product     *product_model.Product
	ProductInfo *product_model.ProductInfo
}
