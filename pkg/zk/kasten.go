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
	Label              string              `json:"label"`
	Zettels            map[Address]*Zettel `json:"zettels"`
	CreatedAt          time.Time           `json:"created_at"`
	UpdatedAt          time.Time           `json:"updated_at"`
	SemanticOrder      []Address           `json:"semantic_order"`
	ChronologicalOrder []Address           `json:"chronological_order"`
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
	k := &Kasten{
		Label:              label,
		Zettels:            map[Address]*Zettel{},
		CreatedAt:          now,
		UpdatedAt:          now,
		SemanticOrder:      []Address{},
		ChronologicalOrder: []Address{},
	}
	RegisterNewCallback(func(z *Zettel) error {
		k.InsertIntoSemanticOrder(z)
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

func (k *Kasten) AddZettel(parent *Address, body string, references string, related ...Address) {
	var address Address
	if parent != nil {
		largestSibling := k.Zettels[*parent].LargestChildAddress
		if largestSibling != nil {
			address = largestSibling.Increment()
		} else {
			address = parent.NewChild()
		}
		k.Zettels[address].LargestChildAddress = &address
	} else {
		address = k.NextMajorAddress()
	}
	z := NewZettel(address, parent, body, references, related...)
	k.Zettels[z.Address] = z
}

func (k *Kasten) InsertIntoSemanticOrder(z *Zettel) error {
	if z.Parent == nil {
		k.SemanticOrder = append(k.SemanticOrder, z.Address)
		return nil
	}
	for i, a := range k.SemanticOrder {
		if a == *k.Zettels[*z.Parent].LargestChildAddress {
			if i == len(k.SemanticOrder)-1 {
				k.SemanticOrder = append(k.SemanticOrder, z.Address)
				return nil
			}
			k.SemanticOrder = append(k.SemanticOrder[:i+1], k.SemanticOrder[i:]...) // index < len(a)
			k.SemanticOrder[i+1] = z.Address
			return nil
		}
	}
	// the only way for this to happen is that a Zettel's parent Zettel has not had it's
	// LatestChildAddress set or has had it set to something invalid
	return fmt.Errorf("failed to insert address %v into semantic order", z.Address)
}

func (k *Kasten) NextMajorAddress() Address {
	if len(k.SemanticOrder) == 0 {
		return Address("1")
	}
	// Get the last Address in the semantic order, which is guaranteed to be the
	// largest one, and increment generation 0
	return k.SemanticOrder[len(k.SemanticOrder)-1].Ancestry()[0].Increment()
}
