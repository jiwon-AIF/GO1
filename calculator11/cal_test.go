package calc

import (
	"math/rand"
	"testing"
)

func TestCount(t *testing.T) {
	total := Count(5)
	if total != 6 {
		t.Errorf("Count was incorrect, got: %d, want: %d.", total, 6)
	}
}

func TestRandom(t *testing.T) {
	ran := rand.Intn(45)
	if ran != rand.Intn(45) {
		t.Errorf("Count was incorrect, got: %d, want: %d.", ran, rand.Intn(45))
	}
}
