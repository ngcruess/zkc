package zk

import (
	"fmt"
	"regexp"
	"strings"
)

const pattern string = "([a-z]+)|([1-9]+)"

var re *regexp.Regexp // lazy loaded

type Address string

// AncestorAtDepth returns the Address which is the receiver Address'
// ancestor at the given tree depth.
// Depth 0 is the global structure parent, and Depth 1 is the true semantic
// origin of this Address' ancestry
// Returns:
// 	- Address or `nil` if the specified depth exceeds the depth of the received Address
// 	- error if the specified depth exceeds the depth of the received Address or `nil` if
//  if the ancestor is found
//
// Example: "1a42b7".AncestorAtDepth(3) = "1a42"
func (a Address) AncestorAtDepth(depth int) (*Address, error) {
	// Generation 0 is the  global parent, the content-less structure Zettel
	// used as the trunk of the Zettel tree.
	if depth == 0 {
		origin := Address("0")
		return &origin, nil
	}
	ancestors := a.Ancestry()
	if depth > len(ancestors) {
		return nil, fmt.Errorf("depth %d exceeds depth of this address", depth)
	}
	address := ancestors[depth-1]
	return &address, nil
}

// Parts returns an array of component parts, basically layers, of an address
// in order.
//
// Example: "1a42b7".Parts = ["1", "a", "42", "b", "7"]
func (a Address) Parts() []string {
	if re == nil {
		re = regexp.MustCompile(pattern)
	}
	return re.FindAllString(string(a), -1)
}

// Ancestry returns an array of all Addresses preceeding this one in the
// family tree in order. The first item is the trunk of the tree, while the
// last one will be the receiving Address itself.
//
// Example: "1a42b7".Ancestry = ["1", "1a", "1a42", "1a42b", "1a42b7"]
func (a Address) Ancestry() []Address {
	parts := a.Parts()
	ancestors := make([]Address, len(parts))
	for i := range parts {
		ancestors[i] = Address(strings.Join(parts[:i+1], ""))
	}
	return ancestors
}
