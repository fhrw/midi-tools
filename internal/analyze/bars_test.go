package analyze

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalcSigLengths(t *testing.T) {
	t.Parallel()

	t.Run("gets correct siglengths", func(t *testing.T) {
		sigs := SigList{
			{Num: 4, Denom: 4, AbsTicks: 0},
			{Num: 3, Denom: 4, AbsTicks: 3840},
		}
		want := []SigLength{
			{Sig: TimeSig{Num: 4, Denom: 4, AbsTicks: 0}, Bars: 1},
			{Sig: TimeSig{Num: 3, Denom: 4, AbsTicks: 3840}, Bars: 120},
		}
		got := sigs.CalcSigLengths([]SigLength{})
		require.Equal(t, want, got)
	})
}

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
