package calc

import (
	"testing"
	"time"
)

func TestWeekCount(t *testing.T) {
	start, _ := time.Parse("2006-01-02", "2016-01-01")
	end, _ := time.Parse("2006-01-02", "2016-05-22")
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		start time.Time
		end   time.Time
		// Expected results.
		want int
	}{
		{"Current week", start, end, 20},
	}
	for _, tt := range tests {
		if got := WeekCount(tt.start, tt.end); got != tt.want {
			t.Errorf("%q. WeekCount() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
