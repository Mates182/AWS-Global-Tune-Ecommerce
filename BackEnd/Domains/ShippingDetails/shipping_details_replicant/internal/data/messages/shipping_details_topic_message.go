package message

import "models-replicant/internal/data/models"

type ShippingDetailsTopicMessage struct {
	Service         string                 `json:"Service"`
	Action          string                 `json:"Action"`
	Target          string                 `json:"Target"`
	ShippingDetails models.ShippingDetails `json:"ShippingDetails"`
}
