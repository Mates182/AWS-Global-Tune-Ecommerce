// auto-generated with ginshot
package service

import (
	requests "delete_product/internal/data/requests"
	responses "delete_product/internal/data/responses"
)

type DeleteProductService interface {
	DeleteProductHandler(request requests.DeleteProductRequest) (int, responses.DeleteProductResponse)
}
