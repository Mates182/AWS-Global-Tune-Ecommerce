package request

import "delete-tracking-details-service/models"

type DeleteTrackingDetailsRequest struct {
	TrackingDetails models.TrackingDetails `json:"TrackingDetails"`
}
