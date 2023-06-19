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

	notes := analyze.GetAllNotes(file)

	for _, n := range notes {
		fmt.Println("--- Note ---")
		fmt.Printf("channel: %v\npitch:%v\nstart:%v\n", n.Channel, n.Pitch, n.Start)
	}

}
