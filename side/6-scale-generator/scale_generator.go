// Package scale solves the Scale Generator side exercise at exercism.io
package scale

import (
	"strings"
)

// initNotes splits the given 'note' string into slice of notes.
// Just a util func to copy paste from instructions, instead of initing an array :-)
func initNotes(notes string) []string {
	return strings.Split(notes, ", ")
}

// indexOf returns the index of s in the slice. -1 if not found
func indexOf(s string, slice []string) int {
	for i, cur := range slice {
		if cur == s {
			return i
		}
	}
	return -1
}

// same as IndexOf but ignores case
func indexOfIgnoreCase(s string, slice []string) int {
	s = strings.ToLower(s)
	for i, cur := range slice {
		if strings.ToLower(cur) == s {
			return i
		}
	}
	return -1
}

func intervalToIncrement(i int, interval string) int {
	if interval == "" {
		return 1
	}
	i = i % len(interval)

	switch interval[i] {
	case 'm':
		return 1
	case 'M':
		return 2
	case 'A':
		return 3
	default:
		panic("Unknown interval")
	}
}

// contains returns true if the needle is found in the haystack when split by sep
func contains(haystack, sep, needle string) bool {
	toks := strings.Split(haystack, sep)
	return indexOf(needle, toks) != -1
}

// Scale generates the musical scale starting with the tonic
// and following the specified interval pattern.
func Scale(tonic string, interval string) []string {

	sharpScale := "A, A#, B, C, C#, D, D#, E, F, F#, G, G#"
	sharpNotes := "G, D, A, E, B, F#, e, b, f#, c#, g#, d#"

	flatScale := "A, Bb, B, C, Db, D, Eb, E, F, Gb, G, Ab"
	flatNotes := "F, Bb, Eb, Ab, Db, Gb, d, g, c, f, bb, eb"

	// find which scale to use
	var scale []string
	if contains(sharpNotes, ", ", tonic) {
		scale = initNotes(sharpScale)
	} else if contains(flatNotes, ", ", tonic) {
		scale = initNotes(flatScale)
	} else {
		scale = initNotes(sharpScale)
	}
	scaleCount := len(scale)

	// how many notes in the output. Use the interval.
	noteCount := len(strings.TrimSpace(interval))
	if noteCount == 0 {
		noteCount = len(scale)
	}

	var out []string
	startIndex := indexOfIgnoreCase(tonic, scale)
	if startIndex == -1 {
		panic("tonic not found in notes:" + string(tonic))
	}

	// now use the scale and collect the notes based on the interval
	// starting from the given tonic, if not from the beginning
	index := startIndex
	for i := 0; i < noteCount; i++ {
		out = append(out, scale[index])
		increment := intervalToIncrement(i, interval)
		index = (index + increment) % scaleCount
	}

	return out
}
