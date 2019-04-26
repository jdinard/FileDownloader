package downloader

import (
	"testing"
)

func TestCalculateByteRangeZero(t *testing.T) {
	got, got1 := calculateByteRange(0, 100)

	if got != 0 && got1 != 99 {
		t.Errorf("Invalid get calculateByteRange, expected 0 and 99, got %d and %d", got, got1)
	}
}

func TestCalculateByteRangeMiddle(t *testing.T) {
	got, got1 := calculateByteRange(1, 100)

	if got != 100 && got1 != 199 {
		t.Errorf("Invalid get calculateByteRange, expected 100 and 199, got %d and %d", got, got1)
	}
}

func TestCalculateByteStartZero(t *testing.T) {
	got := calculateByteStart(0, 100)

	if got != 0 {
		t.Errorf("Invalid get calculateByteStart, expected 0, got %d", got)
	}
}

func TestCalculateByteStartMiddle(t *testing.T) {
	got := calculateByteStart(1, 100)

	if got != 100 {
		t.Errorf("Invalid get calculateByteStart, expected 100, got %d", got)
	}
}

func TestCalculateByteEndStart(t *testing.T) {
	got := calculateByteEnd(0, 100)

	if got != 99 {
		t.Errorf("Invalid get calculateByteEnd, expected 99, got %d", got)
	}
}

func TestCalculateByteEndMiddle(t *testing.T) {
	got := calculateByteEnd(1, 100)

	if got != 199 {
		t.Errorf("Invalid get calculateByteEnd, expected 199, got %d", got)
	}
}

func TestGetTempFileName(t *testing.T) {
	got := getTempFileName(5, "myfilename")

	if got != "myfilename-5" {
		t.Errorf("Invalid get filename, expected \"myfilename-f\", got %v", got)
	}
}

func TestGetTempFileNameShouldBeConstant(t *testing.T) {
	got := getTempFileName(2, "myfilename")
	got1 := getTempFileName(2, "myfilename")

	if got != got1 {
		t.Errorf("Invalid get filename, expected %v and %v to be equal", got, got1)
	}
}
