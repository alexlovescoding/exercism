package clock

import "fmt"

const (
	testVersion    = 3
	minutesPerDay  = 1440
	minutesPerHour = 60
)

type Clock int

// New creates a new value of type Clock.
func New(hour, minutes int) Clock {
	minutes = (hour*minutesPerHour + minutes) % minutesPerDay
	if minutes < 0 {
		minutes += minutesPerDay
	}
	return Clock(minutes)
}

// String function formats the Clock object in a readable manner.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/minutesPerHour, c%minutesPerHour)
}

// Add function adds the given number of minutes to the Clock.
func (c Clock) Add(minutes int) Clock {
	minutes = int(c) + minutes
	return New(0, minutes)
}
