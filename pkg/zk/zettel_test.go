package zk_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zkc/pkg/zk"
)

func TestAddChild(t *testing.T) {
	// SETUP
	origin := prepareOrigin()

	child := zk.NewZettel(zk.Address("1"), "child", "", "")

	// ACTION
	origin.AddChildWithAddress(child)

	// ASSERTIONS
	actual, ok := origin.Children[child.Address]

	assert.True(t, ok)
	assert.Equal(t, child, actual)
	assert.Equal(t, child.Parent, origin)
}

func TestAddDescendant(t *testing.T) {
	// SETUP
	origin := prepareTree()

	// ACTION
	descendant := zk.NewZettel(zk.Address("1c"), "1c", "", "")
	err := origin.AddDescendent(origin, 1, descendant)

	// ASSERTIONS
	one, _ := origin.Children[zk.Address("1")]
	actual, ok := one.Children[zk.Address("1c")]

	assert.Nil(t, err)
	assert.True(t, ok)
	assert.Equal(t, descendant, actual)
}

func TestGetDescendentByAddress(t *testing.T) {
	// SETUP
	origin := prepareTree()

	// ACTION
	descendent, err := origin.GetDescendentByAddress(zk.Address("1b"))

	// ASSERTIONS
	assert.Nil(t, err)
	assert.Equal(t, zk.Address("1b"), descendent.Address)
	assert.Equal(t, "1b", descendent.Body)
}

func prepareOrigin() *zk.Zettel {
	return zk.NewZettel(zk.Address("0"), "origin", "", "")
}

func prepareTree() *zk.Zettel {
	origin := prepareOrigin()
	origin.Children = map[zk.Address]*zk.Zettel{
		zk.Address("1"): zk.NewZettel(zk.Address("1"), "1", "", ""),
		zk.Address("2"): zk.NewZettel(zk.Address("2"), "2", "", ""),
	}
	origin.Children[zk.Address("1")].Children = map[zk.Address]*zk.Zettel{
		zk.Address("1a"): zk.NewZettel(zk.Address("1a"), "1a", "", ""),
		zk.Address("1b"): zk.NewZettel(zk.Address("1b"), "1b", "", ""),
	}
	return origin
}
