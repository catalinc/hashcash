// Package hashcash provides an implementation of Hashcash version 1 algorithm.
package hashcash

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
	"math"
	"strings"
	"time"
)

// Hash provides an implementation of hashcash v1.
type Hash struct {
	hasher  hash.Hash // SHA-1
	bits    uint      // Number of zero bits
	zeros   uint      // Number of zero digits
	saltLen uint      // Random salt length
	extra   string    // Extension to add to the minted stamp
}

// New creates a new Hash with specified options.
func New(bits uint, saltLen uint, extra string) *Hash {
	h := &Hash{
		hasher:  sha1.New(),
		bits:    bits,
		saltLen: saltLen,
		extra:   extra}
	h.zeros = uint(math.Ceil(float64(h.bits) / 4.0))
	return h
}

// NewStd creates a new Hash with 20 bits of collision and 8 bytes of salt chars.
func NewStd() *Hash {
	return New(20, 8, "")
}

// Date field format
const dateFormat = "060102"

// Mint a new hashcash stamp for resource.
func (h *Hash) Mint(resource string) (string, error) {
	salt, err := h.getSalt()
	if err != nil {
		return "", err
	}
	date := time.Now().Format(dateFormat)
	counter := 0
	var stamp string
	for {
		stamp = fmt.Sprintf("1:%d:%s:%s:%s:%s:%x",
			h.bits, date, resource, h.extra, salt, counter)
		if h.checkZeros(stamp) {
			return stamp, nil
		}
		counter++
	}
}

// Check whether a hashcash stamp is valid.
func (h *Hash) Check(stamp string) bool {
	if h.checkDate(stamp) {
		return h.checkZeros(stamp)
	}
	return false
}

// CheckNoDate checks whether a hashcash stamp is valid ignoring date.
func (h *Hash) CheckNoDate(stamp string) bool {
	return h.checkZeros(stamp)
}

func (h *Hash) getSalt() (string, error) {
	buf := make([]byte, h.saltLen)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}
	salt := base64.StdEncoding.EncodeToString(buf)
	return salt[:h.saltLen], nil
}

func (h *Hash) checkZeros(stamp string) bool {
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

func (h *Hash) checkDate(stamp string) bool {
	fields := strings.Split(stamp, ":")
	if len(fields) != 7 {
		return false
	}
	then, err := time.Parse(dateFormat, fields[2])
	if err != nil {
		return false
	}
	duration := time.Since(then)
	return duration.Hours()*2 <= 48
}
