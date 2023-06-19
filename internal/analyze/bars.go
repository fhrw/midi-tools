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

	for _, event := range tracks[0] { // prefer not to use [0] index ideally
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

func (t TimeSig) GetBarEnd() uint64 {
	return t.absTicks + uint64(t.num)*GetDenomTicks(t.denom)
}

func GetDenomTicks(d uint8) uint64 {
	if d == 2 {
		return 1920
	}

	if d == 4 {
		return 960
	}

	if d == 8 {
		return 480
	}

	if d == 16 {
		return 240
	}

	return 0
}
