package main

import "testing"

func TestPredictPartyVictory(t *testing.T) {
	tests := []struct {
		senate string
		want   string
	}{
		{"RD", "Radiant"},
		{"RDD", "Dire"},
		{"R", "Radiant"},
		{"D", "Dire"},
		{"RDRDR", "Radiant"},
		{"DR", "Dire"},
	}

	for _, tt := range tests {
		got := predictPartyVictory(tt.senate)
		if got != tt.want {
			t.Fatalf("senate=%q: got %q, want %q", tt.senate, got, tt.want)
		}
	}
}
