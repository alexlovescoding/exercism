// Package clock documents general stuff about the package here.
package clock

import "fmt"

const (
	// testVersion here must match `targetTestVersion` in the file clock_test.go.
	testVersion = 3
	// minutesPerDay is used to convert between minutes and hours.
	minutesPerDay = 1440
	// minutesPerHour is used to convert between minutes and hours.
	minutesPerHour = 60
)

type Clock struct {
	hour   int
	minute int
}

// New creates a new value of type Clock.
func New(hour, minutes int) Clock {
	m := (hour*minutesPerHour + minutes) % minutesPerDay
	if m < 0 {
		m += minutesPerDay
	}
	return Clock{m / minutesPerHour, m % minutesPerHour}
}

// String function formats the Clock object in a readable manner.
func (m Clock) String() string {
	return fmt.Sprintf("%02d:%02d", m.hour, m.minute)
}

// Add function adds the given number of minutes to the Clock.
func (m Clock) Add(minutes int) Clock {
	minutes = m.minute + minutes
	return New(m.hour, minutes)
}
