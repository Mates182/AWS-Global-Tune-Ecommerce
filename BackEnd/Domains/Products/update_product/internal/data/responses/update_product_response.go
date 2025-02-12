package response

import "update_product/internal/data/models"

type UpdateProductResponse struct {
	Product models.Product `json:"Product"`
	Message string `json:"Message"`
}
