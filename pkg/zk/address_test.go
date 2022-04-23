package zk_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zkc/pkg/zk"
)

func TestParts(t *testing.T) {
	// SETUP
	type AddressPartsPair struct {
		address zk.Address
		parts   []string
	}
	tests := []AddressPartsPair{
		{
			address: zk.Address("1"),
			parts:   []string{"1"},
		},
		{
			address: zk.Address("1a"),
			parts:   []string{"1", "a"},
		},
		{
			address: zk.Address("1a42b7"),
			parts:   []string{"1", "a", "42", "b", "7"},
		},
		{
			address: zk.Address("1ac4b7"),
			parts:   []string{"1", "ac", "4", "b", "7"},
		},
	}
	for _, test := range tests {
		t.Run(string(test.address), func(t *testing.T) {
			// ACTION + ASSERTION
			assert.Equal(t, test.parts, test.address.Parts())
		})
	}
}

func TestAncestorAtDepth(t *testing.T) {
	// SETUP
	type AncestorTestConf struct {
		address  zk.Address
		depth    int
		ancestor zk.Address
		fails    bool
	}
	tests := []AncestorTestConf{
		{
			address:  zk.Address("1"),
			depth:    0,
			ancestor: zk.Address("0"),
			fails:    false,
		},
		{
			address:  zk.Address("1a42b7"),
			depth:    1,
			ancestor: zk.Address("1"),
			fails:    false,
		},
		{
			address:  zk.Address("1a42b7"),
			depth:    2,
			ancestor: zk.Address("1a"),
			fails:    false,
		},
		{
			address:  zk.Address("1a42b7"),
			depth:    5,
			ancestor: zk.Address("1a42b7"),
			fails:    false,
		},
		{
			address: zk.Address("1a42b7"),
			depth:   6,
			fails:   true,
		},
	}
	for _, test := range tests {
		t.Run(string(test.address), func(t *testing.T) {
			// ACTION
			actual, err := test.address.AncestorAtDepth(test.depth)

			// ASSERTION
			if test.fails {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, test.ancestor, *actual)
			}
		})
	}
}
