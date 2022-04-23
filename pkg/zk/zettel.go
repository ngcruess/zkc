package zk

import (
	"fmt"
	"time"
)

type Zettel struct {
	Address   `json:"address"`
	Body      string              `json:"body"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
	Parent    *Zettel             `json:"parent"`
	Children  map[Address]*Zettel `json:"children"`
}

// `AddChild` adds a child to a Zettel, making sure the specified Address is not already occupied
// This is where the uniqueness of Addresses within a collection is enforced.
// Returns:
// 	- error if the Address of the given Zettel is not unique, or `nil`
func (z *Zettel) AddChild(zettel *Zettel) error {
	if z.Children == nil {
		z.Children = map[Address]*Zettel{
			zettel.Address: zettel,
		}
	} else {
		if _, ok := z.Children[zettel.Address]; ok {
			return fmt.Errorf("zettel already exists with address %v", zettel.Address)
		}
		z.Children[zettel.Address] = zettel
	}
	zettel.Parent = z
	return nil
}

// `NewZettel` creates a new Zettel with the given Address and body, automatically
// setting the `CreatedAt` and `UpdatedAt` values to the current time in the
// current time zone.
// Returns:
// 	- a pointer to the new Zettel
func NewZettel(address Address, body string) *Zettel {
	newZ := Zettel{
		Address:   address,
		Body:      body,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return &newZ
}

// `AddDescendent` uses the Address of the given Zettel to insert it into the correct place
// in the tree.
// Returns:
// 	- error if the Address of the given Zettel is invalid or `nil`
func (z *Zettel) AddDescendent(cur *Zettel, depth int, newZ *Zettel) error {
	addr, err := newZ.Address.AncestorAtDepth(depth)
	// This is only possible if the specified starting depth exceeds the address depth
	if err != nil {
		return err
	}
	if *addr == newZ.Address {
		return cur.AddChild(newZ)
	}
	if inner, ok := cur.Children[*addr]; ok {
		return z.AddDescendent(inner, depth+1, newZ)
	}
	return fmt.Errorf("invalid address for new zettel: %s", newZ.Address)
}

// `GetDescendentByAddress` uses the `Ancestry` of the given Address to
// navigate the tree of the receiver Zettel, finding the Zettel with the matching Address.
// Returns:
//	- pointer to the matching Zettel if it is found, or `nil` if not
//	- error if the given address does not match a Zettel, or `nil` if the Zettel was found
func (z *Zettel) GetDescendentByAddress(address Address) (*Zettel, error) {
	if address == z.Address {
		return z, nil
	}
	descendent := z
	for _, ancestor := range address.Ancestry() {
		if c, ok := descendent.Children[ancestor]; ok {
			descendent = c
		} else {
			return nil, fmt.Errorf("address '%v' not found", address)
		}
	}
	return descendent, nil
}
