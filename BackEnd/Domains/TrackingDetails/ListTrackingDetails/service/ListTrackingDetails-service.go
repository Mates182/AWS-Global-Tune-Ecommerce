// auto-generated with ginshot
package service

import (
	requests "list-tracking-details/data/requests"
	responses "list-tracking-details/data/responses"
)

type ListTrackingDetailsService interface {
	ListTrackingDetailsHandler(request requests.ListTrackingDetailsRequest) (int, responses.ListTrackingDetailsResponse)
}
