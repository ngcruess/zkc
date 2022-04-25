package zk

import (
	"fmt"
	"time"
)

// A `Kasten` is a container and metastore for the for the Zettel tree.
// It's main functional purpose is to keep sets of all Addresses in orders commonly required
// for display.
// Semantic Order: a flattened representation of the Zettel tree in depth-first order.
//	Example: ["1", "1a", "1a1", "1a2", "1b", "2", "2a"]
// Chronological Order: a flattened representation of the Zettel tree in descending chronological order.
// 	This newest (most recently created) Address will be first, and the oldest will be last.
// Zettel callbacks are used to make sure the metastore is updated after Zettel events
type Kasten struct {
	Label              string    `json:"label"`
	Origin             *Zettel   `json:"origin"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	SemanticOrder      []Address `json:"semantic_order"`
	ChronologicalOrder []Address `json:"chronological_order"`
}

// `NewKasten` returns a new Kasten. The origin is the
// global parent, the Zettel with Address="0" and no Parent.
// CreatedAt and UpdatedAt will both be the current time in the
// current time zone. The Semantic and Chronological orders arrays
// will be populated with the Address of the Origin.
// Returns:
//	- the new Kasten
func NewKasten(label string) *Kasten {
	now := time.Now()
	address := Address("0")
	k := &Kasten{
		Label: label,
		Origin: &Zettel{
			Address:            address,
			CreatedAt:          now,
			UpdatedAt:          now,
			LatestChildAddress: address,
		},
		CreatedAt:          now,
		UpdatedAt:          now,
		SemanticOrder:      []Address{address},
		ChronologicalOrder: []Address{address},
	}
	RegisterNewCallback(func(z *Zettel) error {
		k.SemanticOrder, _ = InsertIntoSemanticOrder(k.SemanticOrder, z)
		k.ChronologicalOrder = append([]Address{z.Address}, k.ChronologicalOrder...)
		k.UpdatedAt = z.UpdatedAt
		return nil
	})
	RegisterEditCallback(func(z *Zettel) error {
		k.UpdatedAt = z.UpdatedAt
		return nil
	})
	return k
}

func InsertIntoSemanticOrder(semanticOrder []Address, z *Zettel) ([]Address, error) {
	for i, a := range semanticOrder {
		if a == z.Parent.LatestChildAddress {
			if i == len(semanticOrder)-1 {
				return append(semanticOrder, z.Address), nil
			}
			semanticOrder = append(semanticOrder[:i+1], semanticOrder[i:]...) // index < len(a)
			semanticOrder[i+1] = z.Address
			return semanticOrder, nil
		}
	}
	// the only way for this to happen is that a Zettel's parent Zettel has not had it's
	// LatestChildAddress set or has had it set to something invalid
	return nil, fmt.Errorf("failed to insert address %v into semantic order", z.Address)
}
