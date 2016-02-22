# go-securerandom
[![TravisCI Build Status :: theckman/go-securerandom](https://img.shields.io/travis/theckman/go-securerandom/master.svg?style=flat)](https://travis-ci.org/theckman/go-securerandom)
[![GoDoc :: github.com/theckman/go-securerandom](https://img.shields.io/badge/go--securerandom-GoDoc-blue.svg?style=flat)](https://godoc.org/github.com/theckman/go-securerandom)

Package securerandom is a set of utilities for generating random data from the `crypto/rand` package. The package has the ability to generate bytes from secure-random data, as well as other data from those bytes. That includes numbers and Base64-encoded strings.

## Installation
```
go get -u github.com/theckman/go-securerandom
```

## Usage
Full usage information can be found on the [GoDoc](https://godoc.org/github.com/theckman/go-securerandom) page, but here is a short example:

```Go
import (
	"math/rand"
	"github.com/theckman/go-securerandom"
)

ri64, err := securerandom.GenerateRandomInt64()

// secure-random data is unavailable
if err != nil {
	// handle err
}

rand.Seed(ri64)
```
