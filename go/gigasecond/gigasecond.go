// Package gigasecond has the general stuff for the package.
package gigasecond

import (
	"math"
	"time"
)

// testVersion here must match `targetTestVersion` in the file gigasecond_test.go.
const testVersion = 4

// AddGigasecond adds a gigasecond (10**9 seconds) to the given time.
func AddGigasecond(m time.Time) time.Time {
	return m.Add(time.Duration(math.Pow(10, 9) * math.Pow(10, 9)))
}
