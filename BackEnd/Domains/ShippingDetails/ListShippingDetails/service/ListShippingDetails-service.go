// auto-generated with ginshot
package service

import (
	requests "list-shipping-details/data/requests"
	responses "list-shipping-details/data/responses"
)

type ListShippingDetailsService interface {
	ListShippingDetailsHandler(request requests.ListShippingDetailsRequest) (int, responses.ListShippingDetailsResponse)
}
