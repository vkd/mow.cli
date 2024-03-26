package values

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsBool(t *testing.T) {
	require.True(t, IsBool(NewBool(new(bool), false)))

	require.False(t, IsBool(NewString(new(string), "")))
	require.False(t, IsBool(NewInt(new(int), 0)))
	require.False(t, IsBool(NewStrings(new([]string), nil)))
	require.False(t, IsBool(NewInts(new([]int), nil)))
}
