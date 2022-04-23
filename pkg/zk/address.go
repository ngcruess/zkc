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
