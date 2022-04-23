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

func (z *Zettel) AddChild(zettel Zettel) *Zettel {
	if z.Children == nil {
		z.Children = map[Address]*Zettel{
			zettel.Address: &zettel,
		}
	} else {
		z.Children[zettel.Address] = &zettel
	}
	zettel.Parent = z
	return &zettel
}

func NewZettel(address Address, body string) (*Zettel, error) {
	newZ := Zettel{
		Address:   address,
		Body:      body,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return &newZ, nil
}

func (z *Zettel) AddDescendent(cur *Zettel, depth int, newZ Zettel) error {
	addr, err := newZ.Address.AncestorAtDepth(depth)
	// This is only possible if the specified starting depth exceeds the address depth
	if err != nil {
		return err
	}
	if inner, ok := cur.Children[*addr]; ok {
		return z.AddDescendent(inner, depth+1, newZ)
	} else if *addr == z.Address {
		cur.AddChild(newZ)
		newZ.Parent = cur
		return nil
	} else {
		return fmt.Errorf("invalid address for new zettel: %s", newZ.Address)
	}
}
