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
		if note.Start >= s &&
			note.Start < e &&
			note.Track == t {
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

func Asdf(f *smf.SMF) NoteList {
	var notes NoteList
	for i, track := range f.Tracks {
		var absTicks uint64
		for j, event := range track {
			absTicks += uint64(event.Delta)
			var (
				channel, key, velocity uint8
			)
			if event.Message.GetNoteOn(&channel, &key, &velocity) {
				// find the matching noteOff from the remainder
				remainder := track[j:]
				var end uint64
				for _, event := range remainder {
					var (
						offChannel, offKey, offVelo uint8
					)
					if event.Message.GetNoteOff(&offChannel, &offKey, &offVelo) &&
						offChannel == channel &&
						offKey == key {
						end = absTicks + uint64(event.Delta)
					}
				}
				notes = append(notes, Note{Track: i, Pitch: key, Start: absTicks, End: end})
			}
		}
	}
	return notes
}
