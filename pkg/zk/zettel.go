package zk

import (
	"fmt"
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

// The Callback for adding a new Zettel
var newCallback Callback

// The Callback for editing a Zettel's content
var editCallback Callback

func RegisterNewCallback(c Callback) {
	newCallback = c
}

func RegisterEditCallback(c Callback) {
	editCallback = c
}

// `AddChildWithAddress` adds a child to a Zettel, making sure the specified Address is not already occupied
// This is where the uniqueness of Addresses within a collection is enforced.
// Returns:
// 	- error if the Address of the given Zettel is not unique, or `nil`
func (z *Zettel) AddChildWithAddress(zettel *Zettel) error {
	if z.Children == nil {
		z.Children = map[Address]struct{}{
			zettel.Address: {},
		}
	} else {
		if _, ok := z.Children[zettel.Address]; ok {
			return fmt.Errorf("zettel already exists with address %v", zettel.Address)
		}
		z.Children[zettel.Address] = struct{}{}
	}
	return newCallback(zettel)
}

// `NewZettel` creates a new Zettel with the given Address and body, automatically
// setting the `CreatedAt` and `UpdatedAt` values to the current time in the
// current time zone.
// Returns:
// 	- a pointer to the new Zettel
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
