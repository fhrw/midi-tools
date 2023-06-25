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

func TestMatchOnOffs(t *testing.T) {
	t.Parallel()

	t.Run("does it work?", func(t *testing.T) {
		ons := []NoteOn{
			{Track: 1, Pitch: 64, Start: 960},
			{Track: 1, Pitch: 84, Start: 960},
		}
		offs := []NoteOff{
			{Track: 1, Pitch: 64, End: 1200},
			{Track: 1, Pitch: 84, End: 1505},
		}
		want := NoteList{
			{Track: 1, Pitch: 64, Start: 960, End: 1200},
			{Track: 1, Pitch: 84, Start: 960, End: 1505},
		}
		require.Equal(t, want, MatchOnOffs(ons, offs))
	})
}
