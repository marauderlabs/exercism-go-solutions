package clock

import "fmt"

// Clock represents the clock time in hours and minutes
type Clock int

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

	return Clock(mins)
}

func (c Clock) String() string {
	var h int
	var m int
	i := int(c)

	h = i / minsInHour
	m = i % minsInHour

	return fmt.Sprintf("%02v:%02v", h, m)
}

// Add adds a minutes to c
func (c Clock) Add(a int) Clock {
	return New(0, int(c)+a)
}

// Subtract subtracts s minutes from c
func (c Clock) Subtract(s int) Clock {
	return New(0, int(c)-s)
}
