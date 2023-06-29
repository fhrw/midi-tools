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

func (n NoteList) RemoveOverlaps() NoteList {
	var cleanedNotes NoteList
	for i, note := range n {
		if i == len(n)-1 {
			cleanedNotes = append(cleanedNotes, note)
		}
		if note.End > n[i+1].Start {
			cleanedNotes = append(cleanedNotes, Note{
				Track: note.Track,
				Pitch: note.Pitch,
				Start: note.Start,
				End:   n[i+1].Start,
			})
		} else {
			cleanedNotes = append(cleanedNotes, note)
		}
	}
	return cleanedNotes
}

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

func NoteReader(f *smf.SMF) NoteList {
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
