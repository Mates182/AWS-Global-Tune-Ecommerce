package request

import "create-tracking-details-service/models"

type CreateTrackingDetailsRequest struct {
	TrackingDetails models.TrackingDetails `json:"TrackingDetails"`
}
