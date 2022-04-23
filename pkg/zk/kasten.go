package zk

import (
	"fmt"
	"time"
)

type Kasten struct {
	Label     string              `json:"label"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
	Zettels   map[Address]*Zettel `json:"zettels"`
}

func (k *Kasten) AddZettel(address Address, body string) (*Zettel, error) {
	z := Zettel{
		Address:   address,
		Body:      body,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := k.AppendAtLocation(k.Zettels, 1, z)
	return &z, err
}

func (k *Kasten) AppendAtLocation(cur map[Address]*Zettel, depth int, z Zettel) error {
	addr, err := z.Address.AncestorAtDepth(depth)
	// This is only possible if the specified starting depth exceeds the address depth
	if err != nil {
		return err
	}
	if inner, ok := cur[*addr]; ok {
		return k.AppendAtLocation(inner.Children, depth+1, z)
	} else if *addr == z.Address {
		cur[z.Address] = &z
		return nil
	} else {
		return fmt.Errorf("invalid address for new zettel: %s", z.Address)
	}
}
