package response

import "list-tracking-details/models"

type ListTrackingDetailsResponse struct {
	TrackingDetails []models.TrackingDetails `json:"TrackingDetails"`
	Message         string                   `json:"Message"`
}
