# go-randomizer
[![TravisCI Build Status :: theckman/go-randomizer](https://img.shields.io/travis/theckman/go-randomizer/master.svg?style=flat)](https://travis-ci.org/theckman/go-randomizer)
[![GoDoc :: github.com/theckman/go-randomizer](https://img.shields.io/badge/go--randomizer-GoDoc-blue.svg?style=flat)](https://godoc.org/github.com/theckman/go-randomizer)

Package randomizer is a set of utilities for generating random data from the `crypto/rand` package. The package has the ability to generate bytes from secure-random data, as well as other data from those bytes. That includes numbers and Base64-encoded strings.

## Installation
```
go get -u github.com/theckman/go-randomizer
```

## Usage
Full usage information can be found on the [GoDoc](https://godoc.org/github.com/theckman/go-randomizer) page, but here is a short example:

```Go
import (
	"math/rand"
	"github.com/theckman/randomizer"
)

ri64, err := randomizer.GenerateRandomInt64()

// secure-random data is unavailable
if err != nil {
	// handle err
}

rand.Seed(ri64)
```
