package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 54, Capacity: 71, Latency: 17, Risk: 22, Weight: 13}, wantScore: 56, wantDecision: "review"},
		{name: "case_2", signal: Signal{Demand: 68, Capacity: 75, Latency: 13, Risk: 9, Weight: 10}, wantScore: 172, wantDecision: "review"},
		{name: "case_3", signal: Signal{Demand: 67, Capacity: 107, Latency: 26, Risk: 22, Weight: 12}, wantScore: 95, wantDecision: "review"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.signal); got != tc.wantScore {
				t.Fatalf("score = %d, want %d", got, tc.wantScore)
			}
			if got := Classify(tc.signal); got != tc.wantDecision {
				t.Fatalf("decision = %s, want %s", got, tc.wantDecision)
			}
		})
	}
}
