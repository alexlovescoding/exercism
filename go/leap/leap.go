// Package leap documents general stuff about the package here.
package leap

// TestVersion here must match `testVersion` in the file leap_test.go.
const TestVersion = 1

// IsLeapYear tells you whether the given year is a leap year or not.
func IsLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}
