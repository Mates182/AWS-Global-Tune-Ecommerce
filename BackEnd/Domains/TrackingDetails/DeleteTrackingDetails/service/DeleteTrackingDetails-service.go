// auto-generated with ginshot
package service

import (
	requests "delete-tracking-details-service/data/requests"
	responses "delete-tracking-details-service/data/responses"
)

type DeleteTrackingDetailsService interface {
	DeleteTrackingDetailsHandler(request requests.DeleteTrackingDetailsRequest) (int, responses.DeleteTrackingDetailsResponse)
}
