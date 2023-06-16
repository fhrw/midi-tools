package analyze

import (
	"gitlab.com/gomidi/midi/v2/smf"
)

type TimeSig struct {
	num      uint8
	denom    uint8
	absTicks uint64
}

func GetTimeSigs(f *smf.SMF) []TimeSig {
	tracks := f.Tracks

	if len(tracks) < 1 {
		return []TimeSig{}
	}

	var timeSigs []TimeSig
	var absTicks uint64

	for _, event := range tracks[0] {
		msg := event.Message
		absTicks += uint64(event.Delta)

		var (
			num   uint8
			denom uint8
		)

		if msg.Type() == smf.MetaEndOfTrackMsg {
			// ignore
			continue
		}

		if msg.GetMetaMeter(&num, &denom) {
			newT := TimeSig{
				num:      num,
				denom:    denom,
				absTicks: absTicks,
			}
			timeSigs = append(timeSigs, newT)
		}
	}

	return timeSigs
}
