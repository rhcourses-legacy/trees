package testhelpers

import "testing"

// AssertListsEqual checks if the two given lists are equal.
// If they are not equal, it calls t.Errorf with a message.
func AssertListsEqual[T comparable](t testing.TB, expected, actual []T) {
	if len(expected) != len(actual) {
		t.Errorf("Expected list of length %d, got %d", len(expected), len(actual))
		return
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("expected %v, got %v", expected, actual)
			return
		}
	}
}

// AssertTrue checks if the given condition is true.
// If it is not true, it calls t.Errorf with a message stating that the
// condition described by description is not true.
func AssertTrue(t testing.TB, condition bool, description string) {
	if !condition {
		t.Errorf("Expected %s to be true, but it was false", description)
	}
}

// AssertFalse checks if the given condition is false.
// If it is not false, it calls t.Errorf with a message stating that the
// condition described by description is not false.
func AssertFalse(t testing.TB, condition bool, description string) {
	if condition {
		t.Errorf("Expected %s to be false, but it was true", description)
	}
}

// AssertEqual checks if the two given values are equal.
// If they are not equal, it calls t.Errorf with a message.
func AssertEqual(t testing.TB, expected, actual any) {
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
