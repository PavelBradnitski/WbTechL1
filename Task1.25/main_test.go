package main

import (
	"testing"
	"time"
)

func TestSleepDuration(t *testing.T) {
	start := time.Now()
	Sleep(200 * time.Millisecond)
	elapsed := time.Since(start)

	if elapsed < 200*time.Millisecond {
		t.Errorf("Sleep was too short: %v", elapsed)
	}

	if elapsed > 300*time.Millisecond {
		t.Errorf("Sleep was too long: %v", elapsed)
	}
}
