// Clock stub file

// To use the right term, this is the package *clause*.
// You can document general stuff about the package here if you like.
package clock

import "fmt"

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.

const (
  testVersion = 3
  minutesPerDay = 1440
  minutesPerHour = 60
)

type Clock int

func New(hour, minutes int) Clock {
  var m int = (hour * minutesPerHour + minutes) % minutesPerDay
  if m < 0 {
    m += minutesPerDay
  }
  return Clock(m)
}

func (m Clock) String() string {
  return fmt.Sprintf("%02d:%02d", m/minutesPerHour, m%minutesPerHour)
}

func (m Clock) Add(minutes int) Clock {
  m = (m + Clock(minutes)) % minutesPerDay
  if m < 0 {
    m += minutesPerDay
  }
  return m
}
