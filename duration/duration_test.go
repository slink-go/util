package duration

import (
	"encoding/json"
	"testing"
	"time"
)

func TestDurationMarshalling(t *testing.T) {
	var duration = Duration(5 * time.Second)
	b, err := json.Marshal(duration)
	if err != nil {
		t.Fatalf(err.Error())
	}
	res := string(b)
	if res != `"5s"` {
		t.Fatalf("expected '5s', got '%s'", res)
	}
}
func TestDurationPtrMarshalling(t *testing.T) {
	var duration = Duration(5 * time.Second)
	b, err := json.Marshal(&duration)
	if err != nil {
		t.Fatalf(err.Error())
	}
	res := string(b)
	if res != `"5s"` {
		t.Fatalf("expected '5s', got '%s'", res)
	}
}
func TestDurationUnmarshalling(t *testing.T) {
	var value = `"5s"`
	var duration Duration
	err := json.Unmarshal([]byte(value), &duration)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if duration != Duration(5*time.Second) {
		t.Fatalf("expected Duration('5s'), got '%v'", duration)
	}
}
func TestDurationPtrUnmarshalling(t *testing.T) {
	var value = `"5s"`
	duration := Duration(5 * time.Second)
	var durationPtr = &duration
	err := json.Unmarshal([]byte(value), durationPtr)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if *durationPtr != Duration(5*time.Second) {
		t.Fatalf("expected Duration('5s'), got '%v'", duration)
	}
}
