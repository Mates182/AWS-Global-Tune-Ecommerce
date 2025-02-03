package dbcontext

import "delete_shipping_details/internal/data/models"

type Replicator struct {
	Repository ShippingDetailsRepository
}

func NewReplicator(repository ShippingDetailsRepository) *Replicator {
	return &Replicator{Repository: repository}
}

func (r *Replicator) Replicate(ShippingDetails models.ShippingDetails, action string) error {
	switch action {
	case "create":
		return r.Repository.Create(ShippingDetails)
	case "update":
		return r.Repository.Update(ShippingDetails)
	case "delete":
		return r.Repository.Delete(ShippingDetails.ID)
	default:
		return nil
	}
}