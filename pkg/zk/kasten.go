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
	Label              string               `json:"label"`
	Origin             *Zettel              `json:"origin"`
	CreatedAt          time.Time            `json:"created_at"`
	UpdatedAt          time.Time            `json:"updated_at"`
	SemanticOrder      map[Address]struct{} `json:"semantic_order"`
	ChronologicalOrder map[Address]struct{} `json:"chronological_order"`
}
