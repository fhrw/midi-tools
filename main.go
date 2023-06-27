package main

import (
	"fmt"
	"log"
	"midi-mangler/internal/analyze"

	"gitlab.com/gomidi/midi/v2/smf"
)

func main() {
	file, err := smf.ReadFile("./1m1.mid")
	if err != nil {
		log.Println("something went wrong")
	}

	conductorTrack := file.Tracks[0]
	timeSigs := analyze.TimeSigReader(conductorTrack)

	notes := analyze.Asdf(file)
	_ = notes

	for _, sig := range timeSigs {
		fmt.Println(sig)
	}
}
