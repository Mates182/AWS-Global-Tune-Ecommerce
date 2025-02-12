package message

import "create_product/internal/data/models"

type ProductTopicMessage struct {
	Action string `json:"Action"`
	Target string `json:"Target"`
	Product models.Product `json:"Product"`
	Service string `json:"Service"`
}
