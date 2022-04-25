package zk

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	partsPattern     string = "([a-z]+)|([0-9]+)"
	endDigitsPattern string = "\\w*?(?[0-9]+)$"
)

type Address string

// `AncestorAtDepth` returns the Address which is the receiver Address'
// ancestor at the given tree depth.
// Returns:
// 	- Address or `nil` if the specified depth exceeds the depth of the received Address
// 	- error if the specified depth exceeds the depth of the received Address or `nil` if
//  if the ancestor is found
//
// Example: "1a42b7".AncestorAtDepth(3) = "1a42"
func (a Address) AncestorAtDepth(depth int) (*Address, error) {
	// Generation 0 is the  global parent, the content-less structure Zettel
	// used as the trunk of the Zettel tree.
	ancestors := a.Ancestry()
	if depth > len(ancestors) {
		return nil, fmt.Errorf("depth %d exceeds depth of this address", depth)
	}
	address := ancestors[depth]
	return &address, nil
}

// `Parts` returns an array of component parts, basically layers, of an address
// in order.
//
// Example: "1a42b7".Parts = ["1", "a", "42", "b", "7"]
func (a Address) Parts() []string {
	re := regexp.MustCompile(partsPattern)
	if re == nil {
		re = regexp.MustCompile(partsPattern)
	}
	return re.FindAllString(string(a), -1)
}

// `Ancestry` returns an array of all Addresses preceeding this one in the
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

// `NextSibling` returns a new Address which is one larger, in Zettel semantics, than
// the receiving Address.
// Returns:
//	- the new Address
// Examples:
// 	- "1".NextSibling = "2"
// 	- "1a".NextSibling = "1b"
// 	- "1z".NextSibling = "1za"
// 	- "1a47".NextSibling = "1a48"
func (a Address) NextSibling() Address {
	parts := a.Parts()
	last := parts[len(parts)-1]
	if strings.HasSuffix(last, "z") {
		return Address(string(a) + "a")
	}
	// the current part is numeric, so just add 1
	if i, err := strconv.Atoi(last); err == nil {
		return Address(strings.Join(parts[:len(parts)-1], "") + fmt.Sprint(i+1))
	}
	incremented := last[len(last)-1] + 1
	return Address(string(a)[:len(a)-1] + string(incremented))
}

func (a Address) NewChild() Address {
	parts := a.Parts()
	last := parts[len(parts)-1]
	// the current part is numeric so switch to an alphabetical part
	if _, err := strconv.Atoi(last); err == nil {
		return Address(strings.Join(parts, "") + "a")
	}
	// the current part is alphabetical so switch to numeric
	return Address(strings.Join(parts, "") + "1")
}
