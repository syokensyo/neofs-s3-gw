package contracts

import (
	"github.com/nspcc-dev/neofs-sdk-go/user"
)

// NoOpNNSResolver is a no-op implementation of NNS resolver that always returns false.
type NoOpNNSResolver struct {
}

// HasUser always returns false; it is a no-op implementation.
func (r *NoOpNNSResolver) HasUser(_ string, _ user.ID) (bool, error) {
	return false, nil
}
