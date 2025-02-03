// auto-generated with ginshot
package service

import (
	requests "delete_shipping_details/internal/data/requests"
	responses "delete_shipping_details/internal/data/responses"
)

type DeleteShippingDetailsService interface {
	DeleteShippingDetailsHandler(request requests.DeleteShippingDetailsRequest) (int, responses.DeleteShippingDetailsResponse)
}
