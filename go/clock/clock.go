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
	m := (hour*minutesPerHour + minutes) % minutesPerDay
	if m < 0 {
		m += minutesPerDay
	}
	return Clock(m)
}

// String function formats the Clock object in a readable manner.
func (m Clock) String() string {
	return fmt.Sprintf("%02d:%02d", m/minutesPerHour, m%minutesPerHour)
}

// Add function adds the given number of minutes to the Clock.
func (m Clock) Add(minutes int) Clock {
	minutes = int(m) + minutes
	return New(0, minutes)
}
