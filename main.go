package main

import (
	"log"
	"midi-mangler/internal/analyze"

	"gitlab.com/gomidi/midi/v2/smf"
)

func main() {
	file, err := smf.ReadFile("./1m1.mid")
	if err != nil {
		log.Println("something went wrong")
	}

	noteOn, noteOff := analyze.GetAllNote(file)

	notes := analyze.MatchOnOffs(noteOn, noteOff)
	_ = notes

}
