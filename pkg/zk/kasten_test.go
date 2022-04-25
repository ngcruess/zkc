package zk_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zkc/pkg/zk"
)

func TestInsertIntoSemanticOrderEmptyOrder(t *testing.T) {
	// SETUP
	k := zk.NewKasten("test")

	// ACTION
	k.InsertIntoSemanticOrder(&zk.Zettel{
		Address: zk.Address("1"),
	})

	// ASSERTION
	assert.Equal(t, []zk.Address{zk.Address("1")}, k.SemanticOrder)
}

func TestInsertIntoSemanticOrderNewMajor(t *testing.T) {
	// SETUP
	k := zk.NewKasten("test")
	k.Zettels = map[zk.Address]*zk.Zettel{
		zk.Address("1"): {
			Address: zk.Address("1"),
		},
		zk.Address("2"): {
			Address: zk.Address("2"),
		},
	}
	k.SemanticOrder = []zk.Address{
		zk.Address("1"),
		zk.Address("2"),
	}

	expected := []zk.Address{
		zk.Address("1"),
		zk.Address("2"),
		zk.Address("3"),
	}

	// ACTION
	k.InsertIntoSemanticOrder(&zk.Zettel{
		Address: zk.Address("3"),
	})

	// ASSERTION
	assert.Equal(t, expected, k.SemanticOrder)
}

func TestInsertIntoSemanticOrderChildInMiddle(t *testing.T) {
	largestChild := zk.Address("2b")
	parent := zk.Address("2")

	k := zk.NewKasten("test")
	k.Zettels = map[zk.Address]*zk.Zettel{
		zk.Address("1"): {
			Address: zk.Address("1"),
		},
		zk.Address("2"): {
			Address:             zk.Address("2"),
			LargestChildAddress: &largestChild,
			Children: map[zk.Address]struct{}{
				zk.Address("2a"): {},
				zk.Address("2b"): {},
			},
		},
		zk.Address("2a"): {
			Address: zk.Address("2a"),
			Parent:  &parent,
		},
		zk.Address("3"): {
			Address: zk.Address("3"),
		},
		zk.Address("4"): {
			Address: zk.Address("4"),
		},
	}
	k.SemanticOrder = []zk.Address{
		zk.Address("1"),
		zk.Address("2"),
		zk.Address("2a"),
		zk.Address("3"),
		zk.Address("4"),
	}

	expected := []zk.Address{
		zk.Address("1"),
		zk.Address("2"),
		zk.Address("2a"),
		zk.Address("2b"),
		zk.Address("3"),
		zk.Address("4"),
	}

	// ACTION
	k.InsertIntoSemanticOrder(&zk.Zettel{
		Address:             zk.Address("2b"),
		Parent:              &parent,
		LargestChildAddress: &largestChild,
	})

	// ASSERTION
	assert.Equal(t, expected, k.SemanticOrder)
}
