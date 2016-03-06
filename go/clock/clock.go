// Package clock documents general stuff about the package here
package clock

import "fmt"

const (
	// testVersion here must match `targetTestVersion` in the file clock_test.go.
	testVersion = 3
	// minutesPerDay is used to convert between minutes and hours
	minutesPerDay = 1440
	// minutesPerHour is used to convert between minutes and hours
	minutesPerHour = 60
)

type Clock int

// New creates a new value of type Clock using the hour and minutes parameters
func New(hour, minutes int) Clock {
	var m int = (hour*minutesPerHour + minutes) % minutesPerDay
	if m < 0 {
		m += minutesPerDay
	}
	return Clock(m)
}

// String function formats the Clock object in a readable manner
func (m Clock) String() string {
	return fmt.Sprintf("%02d:%02d", m/minutesPerHour, m%minutesPerHour)
}

// Add function adds the given number of minutes to the Clock
func (m Clock) Add(minutes int) Clock {
	m = (m + Clock(minutes)) % minutesPerDay
	if m < 0 {
		m += minutesPerDay
	}
	return m
}
