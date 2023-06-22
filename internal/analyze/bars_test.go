package analyze

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetBarEnd(t *testing.T) {
	t.Parallel()

	t.Run("should return correct end value", func(t *testing.T) {
		ts := TimeSig{Num: 4, Denom: 4, AbsTicks: 0}
		require.Equal(t, uint64(3840), ts.GetBarEnd(0))
	})
}

func TestGetCurrSig(t *testing.T) {
	t.Parallel()

	t.Run("Gets the correct time sig from the list", func(t *testing.T) {
		sigs := SigList{
			{Num: 4, Denom: 4, AbsTicks: 0},
			{Num: 3, Denom: 4, AbsTicks: 3800},
		}
		want := TimeSig{Num: 4, Denom: 4, AbsTicks: 0}
		got, err := sigs.GetCurrSig(200)
		require.NoError(t, err)
		require.Equal(t, want, got)
	})
}
