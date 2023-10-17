package service

import (
	"testing"
)

// based on config.MaxResultLines = 2_000

func Test_maxLines(t *testing.T) {
	tests := []struct {
		name   string
		params *GetLogParams
		want   int
	}{
		{
			name:   "none",
			params: &GetLogParams{},
			want:   2000,
		},
		{
			name:   "zero",
			params: &GetLogParams{MaxLines: 0},
			want:   2000,
		},
		{
			name:   "negative",
			params: &GetLogParams{MaxLines: -1},
			want:   2000,
		},
		{
			name:   "one",
			params: &GetLogParams{MaxLines: 1},
			want:   1,
		},
		{
			name:   "two",
			params: &GetLogParams{MaxLines: 2},
			want:   2,
		},
		{
			name:   "nineteen_ninety_nine",
			params: &GetLogParams{MaxLines: 1_999},
			want:   1999,
		},
		{
			name:   "two_thousand",
			params: &GetLogParams{MaxLines: 2_000},
			want:   2000,
		},
		{
			name:   "two_thousand_one",
			params: &GetLogParams{MaxLines: 2_001},
			want:   2000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxLines(tt.params); got != tt.want {
				t.Errorf("maxLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
