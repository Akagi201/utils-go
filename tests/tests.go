// Package tests provides various methods which we commonly use when writing
// tests
package tests

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"
)

// RandStr returns a random alphanumeric string of arbitrary length. It's not
// necessary to do any seeding to use this method
func RandStr() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}

// RandInt64 returns a random integer. greater than or equal to 0, and less than
// 1 million. It's not necessary to do any seeding to use this method
func RandInt64() int64 {
	i, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		panic(err)
	}
	return i.Int64()
}
