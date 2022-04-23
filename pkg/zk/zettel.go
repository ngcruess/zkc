package zk

import "time"

type Zettel struct {
	Address   `json:"address"`
	Body      string              `json:"body"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
	Children  map[Address]*Zettel `json:"children"`
}

func (z *Zettel) NewChild(zettel Zettel) *Zettel {
	if z.Children == nil {
		z.Children = map[Address]*Zettel{
			zettel.Address: &zettel,
		}
	} else {
		z.Children[zettel.Address] = &zettel
	}
	return &zettel
}
