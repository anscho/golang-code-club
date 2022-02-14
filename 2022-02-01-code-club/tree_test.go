package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTree(t *testing.T) {

	//region 0

	t.Run("should match leaves", func(t *testing.T) {
		nodeOne := Node{
			value: 3,
		}
		nodeTwo := Node{
			value: 3,
		}
		require.True(t, nodeOne.Equals(nodeTwo))
	})

	t.Run("should not match leaves", func(t *testing.T) {
		nodeOne := Node{
			value: 3,
		}
		nodeTwo := Node{
			value: 2,
		}
		require.False(t, nodeOne.Equals(nodeTwo))
	})

	//endregion

	//region single child

	t.Run("should match left", func(t *testing.T) {
		nodeOne := Node{
			value: 3,
			left: &Node{
				value: 2,
			},
		}
		nodeTwo := Node{
			value: 3,
			left: &Node{
				value: 2,
			},
		}
		require.True(t, nodeOne.Equals(nodeTwo))
	})

	t.Run("should not match left nil", func(t *testing.T) {
		nodeOne := Node{
			value: 3,
			left: &Node{
				value: 2,
			},
		}
		nodeTwo := Node{
			value: 3,
		}
		require.False(t, nodeOne.Equals(nodeTwo))
	})

	t.Run("should not match left value", func(t *testing.T) {
		nodeOne := Node{
			value: 3,
			left: &Node{
				value: 2,
			},
		}
		nodeTwo := Node{
			value: 3,
			left: &Node{
				value: 1,
			},
		}
		require.False(t, nodeOne.Equals(nodeTwo))
	})

	t.Run("should match right", func(t *testing.T) {
		nodeOne := Node{
			value: 3,
			right: &Node{
				value: 4,
			},
		}
		nodeTwo := Node{
			value: 3,
			right: &Node{
				value: 4,
			},
		}
		require.True(t, nodeOne.Equals(nodeTwo))
	})

	t.Run("should not match right nil", func(t *testing.T) {
		nodeOne := Node{
			value: 3,
			right: &Node{
				value: 4,
			},
		}
		nodeTwo := Node{
			value: 3,
		}
		require.False(t, nodeOne.Equals(nodeTwo))
	})

	t.Run("should not match right value", func(t *testing.T) {
		nodeOne := Node{
			value: 3,
			right: &Node{
				value: 4,
			},
		}
		nodeTwo := Node{
			value: 3,
			right: &Node{
				value: 5,
			},
		}
		require.False(t, nodeOne.Equals(nodeTwo))
	})

	//endregion

	//region two children

	t.Run("should match both", func(t *testing.T) {
		nodeOne := Node{
			value: 3,
			left: &Node{
				value: 2,
			},
			right: &Node{
				value: 4,
			},
		}
		nodeTwo := Node{
			value: 3,
			left: &Node{
				value: 2,
			},
			right: &Node{
				value: 4,
			},
		}
		require.True(t, nodeOne.Equals(nodeTwo))
	})

	t.Run("should not match left", func(t *testing.T) {
		nodeOne := Node{
			value: 3,
			left: &Node{
				value: 2,
			},
			right: &Node{
				value: 4,
			},
		}
		nodeTwo := Node{
			value: 3,
			left: &Node{
				value: 1,
			},
			right: &Node{
				value: 4,
			},
		}
		require.False(t, nodeOne.Equals(nodeTwo))
	})

	t.Run("should not match right", func(t *testing.T) {
		nodeOne := Node{
			value: 3,
			left: &Node{
				value: 2,
			},
			right: &Node{
				value: 4,
			},
		}
		nodeTwo := Node{
			value: 3,
			left: &Node{
				value: 2,
			},
			right: &Node{
				value: 5,
			},
		}
		require.False(t, nodeOne.Equals(nodeTwo))
	})

	t.Run("should not match either", func(t *testing.T) {
		nodeOne := Node{
			value: 3,
			left: &Node{
				value: 1,
			},
			right: &Node{
				value: 4,
			},
		}
		nodeTwo := Node{
			value: 3,
			left: &Node{
				value: 2,
			},
			right: &Node{
				value: 5,
			},
		}
		require.False(t, nodeOne.Equals(nodeTwo))
	})

	//endregion

	//region grandchildren

	t.Run("should match grandchildren", func(t *testing.T) {
		nodeOne := Node{
			value: 5,
			left: &Node{
				value: 3,
				left: &Node{
					value: 2,
				},
				right: &Node{
					value: 4,
				},
			},
			right: &Node{
				value: 7,
				left: &Node{
					value: 6,
				},
				right: &Node{
					value: 8,
				},
			},
		}
		nodeTwo := Node{
			value: 5,
			left: &Node{
				value: 3,
				left: &Node{
					value: 2,
				},
				right: &Node{
					value: 4,
				},
			},
			right: &Node{
				value: 7,
				left: &Node{
					value: 6,
				},
				right: &Node{
					value: 8,
				},
			},
		}
		require.True(t, nodeOne.Equals(nodeTwo))
	})

	t.Run("should not match left outer grandchild", func(t *testing.T) {
		nodeOne := Node{
			value: 5,
			left: &Node{
				value: 3,
				left: &Node{
					value: 2,
				},
				right: &Node{
					value: 4,
				},
			},
			right: &Node{
				value: 7,
				left: &Node{
					value: 6,
				},
				right: &Node{
					value: 8,
				},
			},
		}
		nodeTwo := Node{
			value: 5,
			left: &Node{
				value: 3,
				left: &Node{
					value: 0,
				},
				right: &Node{
					value: 4,
				},
			},
			right: &Node{
				value: 7,
				left: &Node{
					value: 6,
				},
				right: &Node{
					value: 8,
				},
			},
		}
		require.False(t, nodeOne.Equals(nodeTwo))
	})

	t.Run("should not match left inner grandchild", func(t *testing.T) {
		nodeOne := Node{
			value: 5,
			left: &Node{
				value: 2,
				left: &Node{
					value: 1,
				},
				right: &Node{
					value: 3,
				},
			},
			right: &Node{
				value: 7,
				left: &Node{
					value: 6,
				},
				right: &Node{
					value: 8,
				},
			},
		}
		nodeTwo := Node{
			value: 5,
			left: &Node{
				value: 2,
				left: &Node{
					value: 1,
				},
				right: &Node{
					value: 4,
				},
			},
			right: &Node{
				value: 7,
				left: &Node{
					value: 6,
				},
				right: &Node{
					value: 8,
				},
			},
		}
		require.False(t, nodeOne.Equals(nodeTwo))
	})

	t.Run("should not match right grandchild", func(t *testing.T) {
		nodeOne := Node{
			value: 5,
			left: &Node{
				value: 3,
				left: &Node{
					value: 2,
				},
				right: &Node{
					value: 4,
				},
			},
			right: &Node{
				value: 7,
				left: &Node{
					value: 6,
				},
				right: &Node{
					value: 8,
				},
			},
		}
		nodeTwo := Node{
			value: 5,
			left: &Node{
				value: 3,
				left: &Node{
					value: 2,
				},
				right: &Node{
					value: 4,
				},
			},
			right: &Node{
				value: 7,
				left: &Node{
					value: 6,
				},
				right: &Node{
					value: 9,
				},
			},
		}
		require.False(t, nodeOne.Equals(nodeTwo))
	})

	//endregion
}
