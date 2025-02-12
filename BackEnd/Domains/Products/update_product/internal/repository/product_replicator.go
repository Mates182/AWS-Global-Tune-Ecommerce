package dbcontext

import "update_product/internal/data/models"

type Replicator struct {
	Repository ProductRepository
}

func NewReplicator(repository ProductRepository) *Replicator {
	return &Replicator{Repository: repository}
}

func (r *Replicator) Replicate(Product models.Product, action string) error {
	switch action {
	case "create":
		return r.Repository.Create(Product)
	case "update":
		return r.Repository.Update(Product)
	case "delete":
		return r.Repository.Delete(Product.ID)
	default:
		return nil
	}
}