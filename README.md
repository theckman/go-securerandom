# go-securerandom
[![TravisCI Build Status :: theckman/go-securerandom](https://img.shields.io/travis/theckman/go-securerandom/master.svg?style=flat)](https://travis-ci.org/theckman/go-securerandom)
[![GoDoc :: github.com/theckman/go-securerandom](https://img.shields.io/badge/go--securerandom-GoDoc-blue.svg?style=flat)](https://godoc.org/github.com/theckman/go-securerandom)
[![License](https://img.shields.io/badge/License-BSD_3--Clause-brightgreen.svg?style=flat)](https://github.com/theckman/go-securerandom/blob/master/LICENSE)

Package securerandom is a set of utilities for generating random data from the `crypto/rand` package. The package has the ability to generate bytes from secure-random data, as well as other data from those bytes. That includes numbers and Base64-encoded strings.

## License
go-securerandom is released under the BSD 3-Clause License. See the LICENSE file for the full contents of the license.

## Installation
```
go get -u github.com/theckman/go-securerandom
```

## Usage
Full usage information can be found on the [GoDoc](https://godoc.org/github.com/theckman/go-securerandom) page, but here is a short example:

```Go
import (
	"fmt"
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

You can also generate random base64 encoded strings from this package:

```Go
// generate string from 32 random bytes
rStr, err := securerandom.Base64OfBytes(32)

// secure-random data is unavailable
if err != nil { /* handle err */ }

fmt.Println(rStr)

// assume your random string can only be 32 bytes long
rStr, _ = securerandom.Base64InBytes(32)

fmt.Println(rStr) // would print Base64 string with a length of 32
```
