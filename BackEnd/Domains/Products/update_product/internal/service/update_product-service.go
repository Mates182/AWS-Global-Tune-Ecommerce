// auto-generated with ginshot
package service

import (
	requests "update_product/internal/data/requests"
	responses "update_product/internal/data/responses"
)

type UpdateProductService interface {
	UpdateProductHandler(request requests.UpdateProductRequest) (int, responses.UpdateProductResponse)
}
