package zk

import (
	"time"
)

type Zettel struct {
	Address             `json:"address"`
	Body                string               `json:"body"`
	References          string               `json:"references"`
	CreatedAt           time.Time            `json:"created_at"`
	UpdatedAt           time.Time            `json:"updated_at"`
	Parent              *Address             `json:"parent"`
	Children            map[Address]struct{} `json:"children"`
	Related             map[Address]struct{} `json:"related"`
	LargestChildAddress *Address             `json:"latest_child_address"`
}

// Callbacks are used to make sure Zettel events are reported
// to the tracking Kasten or other metastore
type Callback func(z *Zettel) error

// The Callback for editing a Zettel's content
var editCallback Callback

func RegisterEditCallback(c Callback) {
	editCallback = c
}

// `AddChild` determines the next child Address for this Zettel then adds it
// to the Zettel's Children and updates the LargestChildAddress value
// This is where the uniqueness of Addresses within a collection is enforced.
// Returns:
//   - the new Address
func (z *Zettel) AddChild() Address {
	var address Address
	if z.LargestChildAddress != nil {
		address = z.LargestChildAddress.NextSibling()
		z.Children = map[Address]struct{}{
			address: {},
		}
	} else {
		address = z.NewChild()
		z.Children[address] = struct{}{}
	}
	z.LargestChildAddress = &address
	return address
}

// `NewZettel` creates a new Zettel with the given Address and body, automatically
// setting the `CreatedAt` and `UpdatedAt` values to the current time in the
// current time zone.
// Returns:
//   - a pointer to the new Zettel
func NewZettel(address Address, parent *Address, body string, references string, related ...Address) *Zettel {
	newZ := Zettel{
		Address:    address,
		Body:       body,
		References: references,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Related:    map[Address]struct{}{},
		Parent:     parent,
	}
	for _, relation := range related {
		newZ.Related[relation] = struct{}{}
	}
	return &newZ
}
