package zk

import "time"

// A `Kasten` is a container and metastore for the for the Zettel tree.
// It's main functional purpose is to keep sets of all Addresses in orders commonly required
// for display.
// Semantic Order: a flattened representation of the Zettel tree in depth-first order.
//	Example: ["1", "1a", "1a1", "1a2", "1b", "2", "2a"]
// Chronological Order: a flattened representation of the Zettel tree in descending chronological order.
// 	This newest (most recently created) Address will be first, and the oldest will be last.
// Key Zettel features are exposed here with wrapper methods that call their respective Zettel
// implementations and also ensure the metadata fields of the Kasten are kept up to date.
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
	return &Kasten{
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
}
