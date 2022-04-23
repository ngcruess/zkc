package zk

import (
	"fmt"
	"regexp"
	"strings"
)

const pattern string = "([a-z]+)|([1-9]+)"

var re *regexp.Regexp

type Address string

func (a Address) AncestorAtDepth(depth int) (*Address, error) {
	if depth == 0 {
		origin := Address("0")
		return &origin, nil
	}
	parts := a.Parts()
	if depth > len(parts) {
		return nil, fmt.Errorf("depth %d exceeds depth of this address", depth)
	}
	address := Address(strings.Join(parts[:depth], ""))
	return &address, nil
}

func (a Address) Parts() []string {
	if re == nil {
		re = regexp.MustCompile(pattern)
	}
	return re.FindAllString(string(a), -1)
}

// Ancestry gets an array of all addresses preceeding this one in the
// family tree in order. The first item is the trunk of the family tree, while the
// last one will be the receiving Address itself.
//
// Example: "1a42b7" -> ["1", "1a", "1a42", "1a42b", "1a42b7"]
func (a Address) Ancestry() []Address {
	parts := a.Parts()
	ancestors := make([]Address, len(parts))
	for i := range parts {
		ancestors[i] = Address(strings.Join(parts[:i+1], ""))
	}
	return ancestors
}
