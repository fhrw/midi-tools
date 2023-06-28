package main

import (
	"fmt"
	"log"
	"midi-mangler/internal/analyze"

	"gitlab.com/gomidi/midi/v2/smf"
)

func main() {
	// file is 480 ticks
	file, err := smf.ReadFile("./1m1.mid")
	if err != nil {
		log.Println("something went wrong")
	}

	conductorTrack := file.Tracks[0]
	timeSigs := analyze.TimeSigReader(conductorTrack)

	notes := analyze.NoteReader(file)

	s, e := timeSigs.GetBar(10)

	barnotes := notes.GetBarNotes(2, s, e)

	fmt.Println(analyze.ReadTrackName(file.Tracks[2]))
	fmt.Println(s, e)
	fmt.Println(barnotes)
}
