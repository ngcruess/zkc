package zk_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zkc/pkg/zk"
)

func TestParts(t *testing.T) {
	// SETUP
	type AddressPartsPair struct {
		Address zk.Address
		parts   []string
	}
	tests := []AddressPartsPair{
		{
			Address: zk.Address("1"),
			parts:   []string{"1"},
		},
		{
			Address: zk.Address("1a"),
			parts:   []string{"1", "a"},
		},
		{
			Address: zk.Address("1a42b7"),
			parts:   []string{"1", "a", "42", "b", "7"},
		},
		{
			Address: zk.Address("1ac4b7"),
			parts:   []string{"1", "ac", "4", "b", "7"},
		},
	}
	for _, test := range tests {
		t.Run(string(test.Address), func(t *testing.T) {
			assert.Equal(t, test.parts, test.Address.Parts())
		})
	}
}
