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

	sigs := analyze.GetTimeSigs(file)
	lengths := sigs.CalcSigLengths([]analyze.SigLength{})

	for _, ele := range lengths {
		fmt.Printf("%v, %v\n", ele.Sig, ele.Bars)
	}

	idx, startCount, segment := analyze.GetSigChunk(5, 0, 0, lengths)
	fmt.Println(idx, startCount, segment)
}
