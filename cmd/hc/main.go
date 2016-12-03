package main

import (
	"flag"
	"fmt"
	"os"

	hc "github.com/catalinc/hashcash"
)

var (
	mint    = flag.String("mint", "", "Mint a new stamp")
	check   = flag.String("check", "", "Check a stamp for validity")
	bits    = flag.Uint("bits", 20, "Specify required collision bits")
	saltLen = flag.Uint("salt", 8, "Salt length")
	extra   = flag.String("extra", "", "Extra extension to a minted stamp")
)

func main() {
	flag.Parse()

	h := hc.New(*bits, *saltLen, *extra)

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
