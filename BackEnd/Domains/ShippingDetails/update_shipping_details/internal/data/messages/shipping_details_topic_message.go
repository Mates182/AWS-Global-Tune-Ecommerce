package message

import "update_shipping_details/internal/data/models"

type ShippingDetailsTopicMessage struct {
	Action string `json:"Action"`
	Target string `json:"Target"`
	ShippingDetails models.ShippingDetails `json:"ShippingDetails"`
	Service string `json:"Service"`
}
