package testkit

import (
	"reflect"
	"strings"
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

func AssertContians(t *testing.T, str string, substr string) {
	if !strings.Contains(str, substr) {
		t.Fatalf("String does not contain %s", substr)
	}
}
