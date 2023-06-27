package analyze

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetBarNotes(t *testing.T) {
	t.Parallel()

	t.Run("get notes in bar", func(t *testing.T) {
		notes := NoteList{
			{Track: 1, Pitch: 64, Start: 0, End: 960},
			{Track: 2, Pitch: 100, Start: 0, End: 960},
			{Track: 1, Pitch: 100, Start: 960, End: 1920},
		}
		want := NoteList{
			{Track: 1, Pitch: 64, Start: 0, End: 960},
			{Track: 1, Pitch: 100, Start: 960, End: 1920},
		}
		require.Equal(t, want, notes.GetBarNotes(1, 0, 2840))
	})
}
