package random

import (
	"math/rand"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const baseIDLength = 16

var seed int64 = 0

// NEW: one shared RNG instance
var rng = rand.New(rand.NewSource(1))

func SetSeed(seedParam int64) {
	seed = seedParam
	rng.Seed(seed) // use the shared RNG
}

// REMOVE creation of fresh RNG each call.
// Keep function, but return shared rng.
func new() *rand.Rand {
	return rng
}

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[new().Intn(len(charset))]
	}
	return string(b)
}

func String() string {
	return stringWithCharset(baseIDLength, charset)
}

func Int(min, max int) int {
	return new().Intn(max-min+1) + min
}

func IntList(max int) int {
	return new().Intn(max - 1)
}
