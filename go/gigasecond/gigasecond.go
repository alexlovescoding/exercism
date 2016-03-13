// Package gigasecond has the general stuff for the package.
package gigasecond

import "time"

// testVersion here must match `targetTestVersion` in the file gigasecond_test.go.
const testVersion = 4

// AddGigasecond adds a gigasecond (10**9 seconds) to the given time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(1000000000000000000))
}
