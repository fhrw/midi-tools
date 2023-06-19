package analyze

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetBarEnd(t *testing.T) {
	t.Parallel()

	t.Run("should return correct end value", func(t *testing.T) {
		ts := TimeSig{num: 4, denom: 4, absTicks: 0}
		require.Equal(t, uint64(3840), ts.GetBarEnd())
	})
}
