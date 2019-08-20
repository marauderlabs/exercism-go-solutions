// gigasecond provides implementation of only AddGigasecond for now
package gigasecond

import "time"

// AddGigasecond adds a gigasecond to the provided Time and returns the same type.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
