package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpin(t *testing.T) {
	result := Spin()

	assert.Len(t, result.Symbols, 3, "Spin should return 3 symbols")
	assert.GreaterOrEqual(t, result.Win, 0.0, "Win amount should be non-negative")
}

func TestCalculateWin(t *testing.T) {
	testCases := []struct {
		name    string
		symbols []string
		want    float64
	}{
		{"All same symbols", []string{"7", "7", "7"}, 100.0},
		{"Different symbols", []string{"7", "BAR", "CHERRY"}, 0.0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := calculateWin(tc.symbols)
			assert.Equal(t, tc.want, got)
		})
	}
}
