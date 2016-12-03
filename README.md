# hashcash

Implementation of [Hashcash](https://en.wikipedia.org/wiki/Hashcash) version 1 in Go.

# Installation

`go get github.com/catalinc/hashcash`

To install the `hc` command line tool:

`cd $GOPATH/github.com/catalin/hashcash/cmd/hc && go install`

# Usage

## API

```
package main

import (
    "fmt"

    hc "github.com/catalinc/hashcash"
) 
 
func main() {
    h := hc.NewStd() // or .New(bits, saltLength, extra)
    
    // Mint a new stamp
    stamp := hc.Mint("something")
    fmt.Println(t)

    // Check a stamp
    valid := hc.Check("1:20:161203:something::+YO19qNZKRs=:a31a2")
    if valid {
        fmt.Println("Valid")
    } else {
        fmt.Println("Invalid")
    }
}
```

## Command line

```
[cata:...ithub.com/catalinc/hashcash]$ hc -help 
Usage of hc:
  -bits uint
    	Specify required collision bits (default 20)
  -check string
    	Check a stamp for validity
  -extra string
    	Extra extension to a minted stamp
  -mint string
    	Mint a new stamp
  -salt uint
    	Salt length (default 8)
[cata:...ithub.com/catalinc/hashcash]$ hc -mint=something
1:20:161203:something::CIyEo8VUcVY=:3ed98b
```

# License

MIT

