package analyze

import "gitlab.com/gomidi/midi/v2/smf"

func ReadTrackName(t smf.Track) string {
	for _, ev := range t {
		msg := ev.Message
		var name string
		if msg.GetMetaTrackName(&name) {
			return name
		}
		if msg.GetMetaText(&name) {
			return name
		}
	}
	return "unnamed track"
}
