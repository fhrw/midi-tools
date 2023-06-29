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

// Gets bar start and end given a siglist and a bar number
func (s SigList) GetBar(b uint) (start, end uint64) {
	var (
		st, en uint64
	)
	currStart := 1
	// walk over each time signature and determine if our bar is within the range
	for _, sig := range s {
		nextStart := currStart + int(sig.BarsLong(480))
		// if it is, we then break the sig into bars and starting from the currStart,
		if b < uint(nextStart) {
			// ticks per bar
			offset := (sig.End - sig.AbsTicks) / uint64(sig.BarsLong(480))
			counter := 0
			// locate the start and end
			for ((uint64(counter) * offset) + sig.AbsTicks) < sig.End {
				if uint(counter+currStart) == b {
					st = sig.AbsTicks + (uint64(counter) * offset)
					en = st + offset
				}
				counter++
			}
			break
		}
		// if not we increment the currStart and move on to next sig
		currStart += int(sig.BarsLong(480))
	}

	return st, en
}

// given ticks per quarter, calculates number of a bars between
// time sig start and end
func (t TimeSig) BarsLong(ticks uint32) uint {
	span := t.End - t.AbsTicks
	barTicks := (ticks * uint32(DenomMultiplier(t.Denom)) * uint32(t.Num))
	return uint(span) / uint(barTicks)
}

func DenomMultiplier(d uint8) float64 {
	switch d {
	case 2:
		return 2.0
	case 4:
		return 1.0
	case 8:
		return 0.5
	case 16:
		return 0.25
	}
	return 1.0
}

// reads sigs from time track
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
					sigs = append(sigs, TimeSig{
						Num:      num,
						Denom:    denom,
						AbsTicks: absticks,
						End:      absticks + localDelta})
					break
				}
				// add something to look for end of track message
			}
		}
	}
	return sigs
}

// given a ticks per quarter (t) and desired value (v) returns the value in ticks
func CalcNoteLength(t uint64, v uint) uint64 {
	return t / (uint64(v) / 4)
}

// probably outdated
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
