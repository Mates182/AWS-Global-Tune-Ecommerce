package request

import "update_product/internal/data/models"

type UpdateProductRequest struct {
	Product models.Product `json:"Product"`
}
