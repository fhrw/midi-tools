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

func ReadTimeSigs(c smf.Track) SigList {
	var sigs SigList
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
