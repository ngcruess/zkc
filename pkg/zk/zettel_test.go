package zk_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zkc/pkg/zk"
)

func TestAddChild(t *testing.T) {
	// SETUP
	origin := prepareOrigin()

	child := origin.NewZettel(zk.Address("1"), "child", "", "")

	// ACTION
	origin.AddChildWithAddress(child)

	// ASSERTIONS
	actual, ok := origin.Children[child.Address]

	assert.True(t, ok)
	assert.Equal(t, child, actual)
	assert.Equal(t, child.Parent, origin)
}

func TestNewChild(t *testing.T) {
	// SETUP
	origin := prepareOrigin()

	// ACTION
	child, err := origin.NewChild("child", "", "")

	// ASSERTIONS
	assert.NoError(t, err)
	assert.Equal(t, zk.Address("1"), child.Address)

	actual, ok := origin.Children[child.Address]

	assert.True(t, ok)
	assert.Equal(t, child, actual)
	assert.Equal(t, child.Parent, origin)
}

func TestAddDescendant(t *testing.T) {
	// SETUP
	origin := prepareTree()

	// ACTION
	descendant := origin.NewZettel(zk.Address("1c"), "1c", "", "")
	err := origin.AddDescendent(origin, 1, descendant)

	// ASSERTIONS
	one, _ := origin.Children[zk.Address("1")]
	actual, ok := one.Children[zk.Address("1c")]

	assert.NoError(t, err)
	assert.True(t, ok)
	assert.Equal(t, descendant, actual)
}

func TestGetDescendentByAddress(t *testing.T) {
	// SETUP
	origin := prepareTree()

	// ACTION
	descendent, err := origin.GetDescendentByAddress(zk.Address("1b"))

	// ASSERTIONS
	assert.NoError(t, err)
	assert.Equal(t, zk.Address("1b"), descendent.Address)
	assert.Equal(t, "1b", descendent.Body)
}

func prepareOrigin() *zk.Zettel {
	return &zk.Zettel{
		Address:            zk.Address("0"),
		Body:               "0",
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		LatestChildAddress: zk.Address("0"),
	}
}

func prepareTree() *zk.Zettel {
	origin := prepareOrigin()
	origin.Children = map[zk.Address]*zk.Zettel{
		zk.Address("1"): origin.NewZettel(zk.Address("1"), "1", "", ""),
		zk.Address("2"): origin.NewZettel(zk.Address("2"), "2", "", ""),
	}
	origin.Children[zk.Address("1")].Children = map[zk.Address]*zk.Zettel{
		zk.Address("1a"): origin.NewZettel(zk.Address("1a"), "1a", "", ""),
		zk.Address("1b"): origin.NewZettel(zk.Address("1b"), "1b", "", ""),
	}
	return origin
}
