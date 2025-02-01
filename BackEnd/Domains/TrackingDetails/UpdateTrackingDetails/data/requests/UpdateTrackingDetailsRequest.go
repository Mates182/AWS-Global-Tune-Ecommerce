package request

import "update-tracking-details-service/models"

type UpdateTrackingDetailsRequest struct {
	TrackingDetails models.TrackingDetails `json:"TrackingDetails"`
}
