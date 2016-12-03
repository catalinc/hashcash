package hashcash_test

import (
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/catalinc/hashcash"
)

func TestNewDefault(t *testing.T) {
	h := hashcash.NewDefault()
	if h == nil {
		t.Error("Expected hashcash.Hash, got nil")
	}
}

func TestNew(t *testing.T) {
	h := hashcash.New(uint(20), uint(16), "")
	if h == nil {
		t.Error("Expected hashcash.Hash, got nil")
	}
}

var stampTests = []struct {
	bits     uint
	saltLen  uint
	extra    string
	resource string
}{
	{20, 8, "", "abc"},
	{10, 10, "asdf", "something"},
	{20, 10, "abc", "something"},
}

func TestStampFormat(t *testing.T) {
	expectedDate := time.Now().Format("060102")
	for _, tt := range stampTests {
		h := hashcash.New(tt.bits, tt.saltLen, tt.extra)
		stamp, err := h.Mint(tt.resource)
		if err != nil {
			t.Errorf("Mint failed for %s with error %v", tt.resource, err)
		}
		fields := strings.Split(stamp, ":")
		if len(fields) != 7 {
			t.Errorf("Expected 7 fields got %d", len(fields))
		}
		ver, err := strconv.Atoi(fields[0])
		if err != nil {
			t.Errorf("Expected version 1, got error %v", err)
		}
		if ver != 1 {
			t.Errorf("Expected version 1, got %d", ver)
		}
		bits, err := strconv.ParseUint(fields[1], 10, 32)
		if err != nil {
			t.Errorf("Expected %d bits, got error %v", tt.bits, err)
		}
		if uint(bits) != tt.bits {
			t.Errorf("Expected %d bits, got %d", tt.bits, bits)
		}
		date := fields[2]
		if date != expectedDate {
			t.Errorf("Expected %s date, got %s", expectedDate, date)
		}
		resource := fields[3]
		if resource != tt.resource {
			t.Errorf("Expected %s resource, got %s", tt.resource, resource)
		}
		extra := fields[4]
		if extra != tt.extra {
			t.Errorf("Expected %s extra, got %s", tt.extra, extra)
		}
		salt := fields[5]
		if uint(len(salt)) != tt.saltLen {
			t.Errorf("Expected %d salt chars, got %d", tt.saltLen, len(salt))
		}
		counter := fields[6]
		if counter == "" {
			t.Errorf("Counter field is empty")
		}
	}
}

var mintAndCheckTests = []string{"abc", "something", "someone@example.net"}

func TestMintAndCheck(t *testing.T) {
	h := hashcash.NewDefault()
	for _, r := range mintAndCheckTests {
		stamp, err := h.Mint(r)
		if err != nil {
			t.Errorf("Mint failed for %s with error %v", r, err)
		}
		if !h.Check(stamp) {
			t.Errorf("Check failed for %s", r)
		}
	}
}
