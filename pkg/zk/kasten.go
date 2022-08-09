package zk

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// A `Kasten` is a container and metastore for the for the Zettel tree.
// Semantic Order: a flattened representation of the Zettel tree in depth-first order.
//
//	Example: ["1", "1a", "1a1", "1a2", "1b", "2", "2a"]
//
// Chronological Order: a flattened representation of the Zettel tree in descending chronological order.
//
//	This newest (most recently created) Address will be first, and the oldest will be last.
type Kasten struct {
	Label              string              `json:"label"`
	Zettels            map[Address]*Zettel `json:"zettels"`
	CreatedAt          time.Time           `json:"created_at"`
	UpdatedAt          time.Time           `json:"updated_at"`
	SemanticOrder      []Address           `json:"semantic_order"`
	ChronologicalOrder []Address           `json:"chronological_order"`
}

// `NewKasten` returns a new Kasten.
// CreatedAt and UpdatedAt will both be the current time in the
// current time zone.
// Returns:
//   - the new Kasten
func NewKasten(label string) *Kasten {
	now := time.Now()
	return &Kasten{
		Label:              label,
		Zettels:            map[Address]*Zettel{},
		CreatedAt:          now,
		UpdatedAt:          now,
		SemanticOrder:      []Address{},
		ChronologicalOrder: []Address{},
	}
}

func (k *Kasten) AddZettel(parent *Address, body string, references string, related ...Address) {
	var address Address
	if parent != nil {
		address = (k.Zettels[*parent].AddChild())
	} else {
		address = k.NextMajorAddress()
	}
	z := NewZettel(address, parent, body, references, related...)
	k.Zettels[z.Address] = z
	k.ChronologicalOrder = append([]Address{z.Address}, k.ChronologicalOrder...)
	k.InsertIntoSemanticOrder(z)
}

func (k *Kasten) InsertIntoSemanticOrder(z *Zettel) error {
	if z.Parent == nil {
		k.SemanticOrder = append(k.SemanticOrder, z.Address)
		return nil
	}
	for i, a := range k.SemanticOrder {
		if i == len(k.SemanticOrder)-1 {
			k.SemanticOrder = append(k.SemanticOrder, z.Address)
			return nil
		}
		if *z.Parent == a {
			// We know this is the most recent child of the parent, so we can
			// use the length of the parent's Children to determine how far
			// ahead in the semantic order this one belongs
			offset := len(k.Zettels[*z.Parent].Children)
			k.SemanticOrder = append(k.SemanticOrder[:i+offset], k.SemanticOrder[i+offset-1:]...) // index < len(a)
			k.SemanticOrder[i+offset] = z.Address
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
	return k.SemanticOrder[len(k.SemanticOrder)-1].Ancestry()[0].NextSibling()
}

func (k *Kasten) PersistZettel(z *Zettel) error {
	content, err := json.MarshalIndent(z, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile(fmt.Sprintf("%s.json", z.Address), content, os.FileMode(0644))
	if err != nil {
		return err
	}
	return nil
}
