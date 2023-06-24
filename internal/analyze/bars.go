package analyze

import (
	"errors"

	"gitlab.com/gomidi/midi/v2/smf"
)

type TimeSig struct {
	Num      uint8
	Denom    uint8
	AbsTicks uint64
}

type SigLength struct {
	Sig  TimeSig
	Bars uint
}

type SigList []TimeSig

func foo(bn, s int, siglen SigLength) uint64 {
	barLen := siglen.Sig.GetBarEnd(0)
	startTime := siglen.Sig.AbsTicks
	startBar := s
	for i := 0; i < int(siglen.Bars); i++ {
		if startBar == bn {
			return startTime
		}
		startTime += barLen
		startBar++
	}
	return startTime
}

func GetSigChunk(bn, count, i int, sigs []SigLength) (int, int, SigLength) {
	// doesn't catch bn out of range
	if count >= bn {
		if i < 1 {
			return 0, (count - int(sigs[0].Bars)) + 1, sigs[0]
		}
		return i - 1, (count - int(sigs[i-1].Bars)) + 1, sigs[i-1]
	}

	return GetSigChunk(bn, (count + int(sigs[i].Bars)), (i + 1), sigs)
}

func (s SigList) CalcSigLengths(sl []SigLength) []SigLength {
	if len(sl) == len(s)-1 {
		// returns 120 for the last bar coz how else can know how long?
		sl = append(sl, SigLength{Sig: s[len(sl)], Bars: 120})
		return sl
	}

	end := s[len(sl)+1]
	current := s[len(sl)]

	bars := BarsBetween(current, end.AbsTicks)
	sl = append(sl, SigLength{Sig: current, Bars: bars})

	return s.CalcSigLengths(sl)
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

func GetTimeSigs(f *smf.SMF) SigList {
	tracks := f.Tracks

	if len(tracks) < 1 {
		return SigList{}
	}

	var timeSigs SigList
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
				Num:      num,
				Denom:    denom,
				AbsTicks: absTicks,
			}
			timeSigs = append(timeSigs, newT)
		}
	}

	return timeSigs
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
