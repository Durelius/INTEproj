package random

import (
	"math/rand"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const baseIDLength = 16

var seed int64 = 0

func SetSeed(seedParam int64) {
	seed = seedParam
}

func new() *rand.Rand {
	if seed == 0 {
		return rand.New(rand.NewSource(int64(rand.Int())))
	}
	return rand.New(rand.NewSource(seed)) // default
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
	random := new()
	return random.Intn(max-min+1) + min
}
func IntList(max int) int {
	return new().Intn(max - 1)
}
