// Package valueobject holds value objects that cannot change state and usually have no identifier
package valueobject

import (
	"github.com/google/uuid"
	"time"
)

// Transaction is a payment between two parties
type Transaction struct {
	// all values lowercase since they are immutable
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
