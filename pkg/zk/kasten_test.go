package zk_test

// func TestInsertIntoSemanticOrder(t *testing.T) {
// 	// SETUP
// 	k := zk.NewKasten("test")

// 	type InsertTestConfig struct {
// 		name          string
// 		semanticOrder []zk.Address
// 		insert        *zk.Zettel
// 		expected      []zk.Address
// 	}

// 	tests := []InsertTestConfig{
// 		{
// 			name:          "0 -> 0, 1",
// 			semanticOrder: []zk.Address{"0"},
// 			insert:        newSemanticInsertFriendlyZettel(zk.Address("1"), zk.Address("0")),
// 			expected:      []zk.Address{"0", "1"},
// 		},
// 		{
// 			name:          "0, 1 -> 0, 1, 2",
// 			semanticOrder: []zk.Address{"0", "1"},
// 			insert:        newSemanticInsertFriendlyZettel(zk.Address("2"), zk.Address("1")),
// 			expected:      []zk.Address{"0", "1", "2"},
// 		},
// 		{
// 			name:          "0, 1, 2 -> 0, 1, 1a, 2",
// 			semanticOrder: []zk.Address{"0", "1", "2"},
// 			insert:        newSemanticInsertFriendlyZettel(zk.Address("1a"), zk.Address("1")),
// 			expected:      []zk.Address{"0", "1", "1a", "2"},
// 		},
// 		{
// 			name:          "0, 1, 1a, 2, 2a, 3, 4, 5 -> 0, 1, 1a, 2, 2a, 2b, 3, 4, 5",
// 			semanticOrder: []zk.Address{"0", "1", "1a", "2", "2a", "3", "4", "5"},
// 			insert:        newSemanticInsertFriendlyZettel(zk.Address("2b"), zk.Address("2a")),
// 			expected:      []zk.Address{"0", "1", "1a", "2", "2a", "2b", "3", "4", "5"},
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			k.SemanticOrder = test.semanticOrder
// 			k.Zettels = test.zettels

// 			// ACTION
// 			err := k.InsertIntoSemanticOrder(test.insert)

// 			// ASSERTION
// 			assert.NoError(t, err)
// 			assert.Equal(t, test.expected, k.SemanticOrder)
// 		})
// 	}
// }

// func newSemanticInsertFriendlyZettel(a zk.Address, parentLast zk.Address) *zk.Zettel {
// 	if parent := "" {

// 	}
// 	child := &zk.Zettel{
// 		Address: a,
// 		Parent:  parent,
// 	}
// 	return child
// }
