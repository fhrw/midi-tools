package analyze

import (
	"gitlab.com/gomidi/midi/v2/smf"
)

type TimeSig struct {
	num   uint8
	denom uint8
}

func Foo(f *smf.SMF) []TimeSig {
	tracks := f.Tracks

	var timeSigs []TimeSig

	for _, event := range tracks[0] {
		msg := event.Message

		var (
			num   uint8
			denom uint8
		)

		if msg.Type() == smf.MetaEndOfTrackMsg {
			// ignore
			continue
		}

		switch {
		case msg.GetMetaMeter(&num, &denom):
		default:
			newT := TimeSig{
				num:   num,
				denom: denom,
			}
			timeSigs = append(timeSigs, newT)
		}

	}
	return timeSigs
}
