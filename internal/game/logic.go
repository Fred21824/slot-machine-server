package game

import (
	"math/rand"
	"time"

	"slot-machine-server/pkg/models"

	"github.com/spf13/viper"
)

var symbols []string

func init() {
	rand.Seed(time.Now().UnixNano())
	symbols = viper.GetStringSlice("game.symbols")
}

func Spin() models.SpinResult {
	result := models.SpinResult{
		Symbols: make([]string, 3),
	}

	for i := 0; i < 3; i++ {
		result.Symbols[i] = symbols[rand.Intn(len(symbols))]
	}

	result.Win = calculateWin(result.Symbols)
	return result
}

func calculateWin(symbols []string) float64 {
	// Implement win calculation logic here
	// This is a simplified example
	if symbols[0] == symbols[1] && symbols[1] == symbols[2] {
		return 100.0
	}
	return 0.0
}
