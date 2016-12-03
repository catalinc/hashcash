// Package hashcash provides an implementation of hashcash version 1 proof-of-work algorithm.
package hashcash

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
	"math"
	"time"
)

const (
	// Bits is the default value for the number of collision bits.
	Bits = 20
	// SaltSize is default value for random salt size.
	SaltSize = 8
	// DateFormat is default date format.
	DateFormat = "060102"
)

// Hash provides an implementation of hashcash v1 algorithm.
type Hash struct {
	hasher hash.Hash
	bits   uint
	zeros  uint
	salt   uint
}

// New creates a new Hash with 20 bits of collision.
func New() *Hash {
	zeros := uint(math.Ceil(float64(defaultBits) / 4.0))
	return &Hash{hasher: sha1.New(), bits: defaultBits, zeros: zeros, salt: defaultSalt}
}

// Mint a new hashcash stamp for resource.
func (h *Hash) Mint(resource string) (string, error) {
	salt, err := h.computeSalt()
	if err != nil {
		return "", err
	}
	date := time.Now().Format("060102")
	counter := 0
	var stamp string
	for {
		stamp = fmt.Sprintf("1:%d:%s:%s::%s:%x", h.bits, date, resource, salt, counter)
		if h.check(stamp) {
			return stamp, nil
		}
		counter++
	}
}

// Check whether a hashcash stamp is valid.
func (h *Hash) Check(stamp string) bool {
	return h.check(stamp)
}

func (h *Hash) computeSalt() (string, error) {
	buf := make([]byte, h.salt)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}
	salt := base64.StdEncoding.EncodeToString(buf)
	return salt, nil
}

func (h *Hash) check(stamp string) bool {
	h.hasher.Reset()
	h.hasher.Write([]byte(stamp))
	digest := hex.EncodeToString(h.hasher.Sum(nil))
	for i := uint(0); i < h.zeros; i++ {
		if digest[i] != '0' {
			return false
		}
	}
	return true
}
