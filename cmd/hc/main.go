package main

import (
	"flag"
	"fmt"
	"os"

	hc "github.com/catalinc/hashcash"
)

var (
	mint  = flag.String("mint", "", "Mint a new stamp")
	check = flag.String("check", "", "Check a stamp for validity")
	bits  = flag.Int("bits", 20, "Specify required collision bits")
)

func main() {
	flag.Parse()

	if *bits <= 0 {
		fmt.Fprintln(os.Stderr, "Invalid value for collision bits")
		os.Exit(1)
	}
	h := hc.New(uint(*bits))

	if *mint != "" {
		stamp, err := h.Mint(*mint)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(stamp)
		os.Exit(0)
	}

	if *check != "" {
		if h.Check(*check) {
			fmt.Println("Valid")
		} else {
			fmt.Println("Invalid")
		}
	}
}
