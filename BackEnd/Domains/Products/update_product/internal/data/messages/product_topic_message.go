package message

import "update_product/internal/data/models"

type ProductTopicMessage struct {
	Service string `json:"Service"`
	Action string `json:"Action"`
	Target string `json:"Target"`
	Product models.Product `json:"Product"`
}
