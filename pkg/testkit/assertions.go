package testkit

import (
	"reflect"
	"testing"
)

func AssertNoError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("Error is not nil: %v", err)
	}
}

func AssertEqual(t *testing.T, actual any, expected any) {
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("%v is not equal to %v", actual, expected)
	}
}
