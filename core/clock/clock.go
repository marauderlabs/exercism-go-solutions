package clock

import "fmt"

// Clock represents the clock time as minutes since midnight (0000 hours or 12:00AM)
type Clock struct {
	mins int
}

const (
	minsInHour = 60
	hoursInDay = 24
	minsInDay  = minsInHour * hoursInDay
)

// New returns an instance of Clock for the given hour and minute
func New(h, m int) Clock {

	mins := (h*minsInHour + m) % minsInDay
	if mins < 0 {
		mins += minsInDay
	}

	return Clock{mins}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.mins/minsInHour, c.mins%minsInHour)
}

// Add adds a minutes to c
func (c Clock) Add(a int) Clock {
	return New(0, c.mins+a)
}

// Subtract subtracts s minutes from c
func (c Clock) Subtract(s int) Clock {
	return New(0, c.mins-s)
}
