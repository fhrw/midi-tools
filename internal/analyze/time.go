package analyze

import (
	"errors"

	"gitlab.com/gomidi/midi/v2/smf"
)

type TimeSig struct {
	Num      uint8
	Denom    uint8
	AbsTicks uint64
	Length   uint
}

type SigList []TimeSig

func ReadTimeSigs(f *smf.SMF) SigList {
	var sigs SigList

	return sigs
}

func BarsBetween(s TimeSig, e uint64) uint {
	barTicks := s.GetBarEnd(0)
	return uint(e) / uint(barTicks)
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

func (t TimeSig) GetBarEnd(s uint64) uint64 {
	return s + uint64(t.Num)*GetDenomTicks(t.Denom)
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
