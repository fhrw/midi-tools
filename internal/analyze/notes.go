package analyze

import (
	"gitlab.com/gomidi/midi/v2/smf"
)

type Note struct {
	Track int
	Pitch uint8
	Start uint64
	End   uint64
}

type NoteOn struct {
	Track int
	Pitch uint8
	Start uint64
}

type NoteOff struct {
	Track int
	Pitch uint8
	End   uint64
}

type NoteList []Note

func (n NoteList) GetBarNotes(t int, s, e uint64) NoteList {
	var notes []Note
	for _, note := range n {
		if note.Start >= s && note.Start < e && note.Track == t {
			notes = append(notes, note)
		}
	}
	return notes
}

func MatchOnOffs(on []NoteOn, off []NoteOff) NoteList {
	var notes []Note

	for _, n := range on {
		var end uint64
		for _, f := range off {
			if n.Track == f.Track &&
				n.Pitch == f.Pitch &&
				n.Start < f.End {
				end = f.End
			}
		}
		notes = append(notes, Note{
			Track: n.Track,
			Pitch: n.Pitch,
			Start: n.Start,
			End:   end,
		})
	}

	return notes
}

func GetAllNote(f *smf.SMF) ([]NoteOn, []NoteOff) {
	var noteOns []NoteOn
	var noteOffs []NoteOff

	for i, track := range f.Tracks {
		var absTicks uint64

		for _, event := range track {
			absTicks += uint64(event.Delta)
			msg := event.Message

			var (
				key     uint8
				channel uint8
				velo    uint8
			)

			if msg.GetNoteOn(&channel, &key, &velo) {
				note := NoteOn{
					Track: i,
					Pitch: key,
					Start: absTicks,
				}
				noteOns = append(noteOns, note)
			}

			if msg.GetNoteOff(&channel, &key, &velo) {
				note := NoteOff{
					Track: i,
					Pitch: key,
					End:   absTicks,
				}
				noteOffs = append(noteOffs, note)
			}
		}

	}

	return noteOns, noteOffs
}
