package t

import (
	"testing"
	"time"
)

func TestTimeBetween(t *testing.T) {
	sp, _ := time.Parse("2006-01-02 15:04:05", "2023-01-01 12:00:00")
	ep, _ := time.Parse("2006-01-02 15:04:05", "2023-01-07 12:00:00")
	//2024-03-18 11:23:02
	st := time.Unix(1710732182, 0)
	//2024-03-20 11:23:02
	et := time.Unix(1710904982, 0)
	tests := []struct {
		name  string
		start time.Time
		end   time.Time
		diff  time.Duration
		want  bool
	}{
		{
			name:  "Determine if two times are within 6 day true",
			start: sp,
			end:   ep,
			diff:  6 * time.Hour * 24,
			want:  true,
		},
		{
			name:  "Determine if two times are within 5 day false",
			start: sp,
			end:   ep,
			diff:  5 * time.Hour * 24,
			want:  false,
		},
		{
			name:  "Determine if two times are within 2 day true",
			start: st,
			end:   et,
			diff:  2 * time.Hour * 24,
			want:  true,
		},
		{
			name:  "Determine if two times are within 1 day false",
			start: st,
			end:   et,
			diff:  1 * time.Hour * 24,
			want:  false,
		},
		{
			name:  "When the start time is less than the end time",
			start: et,
			end:   st,
			diff:  1 * time.Hour * 24,
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeBetween(tt.start, tt.end, tt.diff); got != tt.want {
				t.Errorf("TimeBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
