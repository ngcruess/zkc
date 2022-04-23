package zk_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zkc/pkg/zk"
)

func TestParts(t *testing.T) {
	// SETUP
	type addressPartsPair struct {
		address zk.Address
		parts   []string
	}
	tests := []addressPartsPair{
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
	type ancestorTestConf struct {
		address  zk.Address
		depth    int
		ancestor zk.Address
		fails    bool
	}
	tests := []ancestorTestConf{
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

func TestAncestry(t *testing.T) {
	// SETUP
	address := zk.Address("1a42b7")
	expected := []zk.Address{
		zk.Address("1"),
		zk.Address("1a"),
		zk.Address("1a42"),
		zk.Address("1a42b"),
		address,
	}

	// ACTION
	ancestors := address.Ancestry()

	// ASSERTION
	assert.Equal(t, expected, ancestors)
}

func TestIncrement(t *testing.T) {
	// SETUP
	type originalIncrementedPair struct {
		original zk.Address
		expected zk.Address
	}
	tests := []originalIncrementedPair{
		{
			original: zk.Address("0"),
			expected: zk.Address("1"),
		},
		{
			original: zk.Address("1a42b7"),
			expected: zk.Address("1a42b8"),
		},
		{
			original: zk.Address("1a"),
			expected: zk.Address("1b"),
		},
		{
			original: zk.Address("1z"),
			expected: zk.Address("1za"),
		},
		{
			original: zk.Address("1ay"),
			expected: zk.Address("1az"),
		},
		{
			original: zk.Address("1a42b77"),
			expected: zk.Address("1a42b78"),
		},
	}

	for _, test := range tests {
		t.Run(string(test.original), func(t *testing.T) {
			// ACTION
			actual := test.original.Increment()

			// ASSERTION
			assert.Equal(t, test.expected, actual)
		})
	}
}
