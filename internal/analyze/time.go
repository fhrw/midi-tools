package analyze

import (
	"errors"

	"gitlab.com/gomidi/midi/v2/smf"
)

type TimeSig struct {
	Num      uint8
	Denom    uint8
	AbsTicks uint64
	End      uint64
}

type SigList []TimeSig

func TimeSigReader(t smf.Track) SigList {
	var sigs SigList
	var absticks uint64
	for i, ev := range t {
		absticks += uint64(ev.Delta)
		msg := ev.Message
		if msg.Is(smf.MetaTimeSigMsg) {
			// find the next one
			var localDelta uint64
			rest := t[i+1:]
			for _, next := range rest {
				localDelta += uint64(next.Delta)
				nextMsg := next.Message
				if nextMsg.Is(smf.MetaTimeSigMsg) {
					// this means you've found the next one!
					// create new and break
					var (
						num, denom uint8
					)
					msg.GetMetaMeter(&num, &denom)
					sigs = append(sigs, TimeSig{Num: num, Denom: denom, AbsTicks: absticks, End: absticks + localDelta})
					break
				}
			}
		}
	}
	return sigs
}

func (s SigList) GetCurrSig(t uint64) (TimeSig, error) {
	reverseSigs := ReverseSigList(s)

	for _, sig := range reverseSigs {
		if t >= sig.AbsTicks {
			return sig, nil
		}
	}

	return TimeSig{}, errors.New("time outside valid range of sigs")
}

func ReverseSigList(s SigList) SigList {
	clone := make(SigList, len(s))
	copy(clone, s)
	for i, j := 0, len(clone)-1; i < j; i, j = i+1, j-1 {
		clone[i], clone[j] = clone[j], clone[i]
	}
	return clone
}
