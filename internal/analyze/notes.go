package analyze

import (
	"gitlab.com/gomidi/midi/v2/smf"
)

type Note struct {
	Channel uint8
	Pitch   uint8
	Start   uint64
	End     uint64
}

func GetAllNotes(f *smf.SMF) []Note {
	var notes []Note

	for _, track := range f.Tracks {
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
				note := Note{
					Channel: channel,
					Pitch:   key,
					Start:   absTicks,
				}
				notes = append(notes, note)
			}
		}

	}

	return notes
}
