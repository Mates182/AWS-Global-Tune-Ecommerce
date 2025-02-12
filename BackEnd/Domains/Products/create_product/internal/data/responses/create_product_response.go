package response

import "create_product/internal/data/models"

type CreateProductResponse struct {
	Product models.Product `json:"Product"`
	Message string `json:"Message"`
}
