package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 54, Capacity: 71, Latency: 17, Risk: 22, Weight: 13}
	if got := Score(signal); got != 56 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 68, Capacity: 75, Latency: 13, Risk: 9, Weight: 10}
	if got := Score(signal); got != 172 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 67, Capacity: 107, Latency: 26, Risk: 22, Weight: 12}
	if got := Score(signal); got != 95 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
}
