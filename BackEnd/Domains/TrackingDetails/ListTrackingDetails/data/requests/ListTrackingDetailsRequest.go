package request

import "list-tracking-details/models"

type ListTrackingDetailsRequest struct {
	TrackingDetails models.TrackingDetails `json:"TrackingDetails"`
}
