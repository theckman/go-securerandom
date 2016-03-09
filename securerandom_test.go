// Copyright 2016 Tim Heckman. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package securerandom_test

import (
	mrand "math/rand"
	"testing"

	"github.com/theckman/go-securerandom"

	. "gopkg.in/check.v1"
)

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func Test(t *testing.T) { TestingT(t) }

// noop is a function so that we can assert the type of a value
// returned from a function without causing the compiler to complain
// because we aren't using the variable in a meaningful way
func noop(interface{}) {}

func (*TestSuite) TestBytes(c *C) {
	// I have no fucking idea how to test this shit...
	var b []byte
	var err error

	b, err = securerandom.Bytes(10)
	c.Assert(err, IsNil)
	c.Check(len(b), Equals, 10) // ¯\_(ツ)_/¯
}

func (t *TestSuite) BenchmarkBytesBy1(c *C) {
	for i := 0; i < c.N; i++ {
		securerandom.Bytes(1)
		c.SetBytes(1)
	}
}

func (t *TestSuite) BenchmarkBytesBy16(c *C) {
	for i := 0; i < c.N; i++ {
		securerandom.Bytes(16)
		c.SetBytes(16)
	}
}

func (t *TestSuite) BenchmarkBytesBy32(c *C) {
	for i := 0; i < c.N; i++ {
		securerandom.Bytes(32)
		c.SetBytes(32)
	}
}

func (t *TestSuite) BenchmarkBytesBy64(c *C) {
	for i := 0; i < c.N; i++ {
		securerandom.Bytes(64)
		c.SetBytes(64)
	}
}

func (*TestSuite) TestBase64InBytes(c *C) {
	var s string
	var err error

	s, err = securerandom.Base64InBytes(32)
	c.Assert(err, IsNil)
	c.Check(len(s), Equals, 32)
}

func (t *TestSuite) BenchmarkBase64InBytes(c *C) {
	var s string
	for i := 0; i < c.N; i++ {
		s, _ = securerandom.Base64InBytes(16)
		c.SetBytes(int64(len(s)))
	}
}

func (*TestSuite) TestURLBase64InBytes(c *C) {
	var s string
	var err error

	s, err = securerandom.URLBase64InBytes(32)
	c.Assert(err, IsNil)
	c.Check(len(s), Equals, 32)
}

func (t *TestSuite) BenchmarkURLBase64In(c *C) {
	var s string
	for i := 0; i < c.N; i++ {
		s, _ = securerandom.URLBase64InBytes(16)
		c.SetBytes(int64(len(s)))
	}
}

func (*TestSuite) TestBase64OfBytes(c *C) {
	var s string
	var err error

	s, err = securerandom.Base64OfBytes(32)
	c.Assert(err, IsNil)
	c.Check(len(s), Equals, 44)
}

func (t *TestSuite) BenchmarkBase64OfBytes(c *C) {
	var s string
	for i := 0; i < c.N; i++ {
		s, _ = securerandom.Base64OfBytes(16)
		c.SetBytes(int64(len(s)))
	}
}

func (*TestSuite) TestURLBase64OfBytes(c *C) {
	var s string
	var err error

	s, err = securerandom.URLBase64OfBytes(32)
	c.Assert(err, IsNil)
	c.Check(len(s), Equals, 44)
}

func (t *TestSuite) BenchmarkURLBase64OfBytes(c *C) {
	var s string
	for i := 0; i < c.N; i++ {
		s, _ = securerandom.URLBase64OfBytes(16)
		c.SetBytes(int64(len(s)))
	}
}

func (t *TestSuite) TestUint16(c *C) {
	var u16 uint16
	var err error

	u16, err = securerandom.Uint16()
	c.Assert(err, IsNil)
	noop(u16)
}

func (t *TestSuite) BenchmarkUint16(c *C) {
	for i := 0; i < c.N; i++ {
		securerandom.Uint16()
		c.SetBytes(2)
	}
}

func (t *TestSuite) TestUint32(c *C) {
	var u32 uint32
	var err error

	u32, err = securerandom.Uint32()
	c.Assert(err, IsNil)
	noop(u32)
}

func (t *TestSuite) BenchmarUint32(c *C) {
	for i := 0; i < c.N; i++ {
		securerandom.Uint32()
		c.SetBytes(4)
	}
}

func (t *TestSuite) TestUint64(c *C) {
	var u64 uint64
	var err error

	u64, err = securerandom.Uint64()
	c.Assert(err, IsNil)
	noop(u64)
}

func (t *TestSuite) BenchmarkUint64(c *C) {
	for i := 0; i < c.N; i++ {
		securerandom.Uint64()
		c.SetBytes(8)
	}
}

func (t *TestSuite) TestInt16(c *C) {
	var i16 int16
	var err error

	i16, err = securerandom.Int16()
	c.Assert(err, IsNil)
	noop(i16)
}

func (t *TestSuite) BenchmarkInt16(c *C) {
	for i := 0; i < c.N; i++ {
		securerandom.Int16()
		c.SetBytes(2)
	}
}

func (t *TestSuite) TestInt32(c *C) {
	var i32 int32
	var err error

	i32, err = securerandom.Int32()
	c.Assert(err, IsNil)
	noop(i32)
}

func (t *TestSuite) BenchmarkInt32(c *C) {
	for i := 0; i < c.N; i++ {
		securerandom.Int32()
		c.SetBytes(4)
	}
}

func (t *TestSuite) TestInt64(c *C) {
	var i64 int64
	var err error

	i64, err = securerandom.Int64()
	c.Assert(err, IsNil)
	noop(i64)
}

func (t *TestSuite) BenchmarkInt64(c *C) {
	for i := 0; i < c.N; i++ {
		securerandom.Uint64()
		c.SetBytes(8)
	}
}

func (*TestSuite) TestRandSource(c *C) {
	var src mrand.Source
	var err error

	src, err = securerandom.RandSource()
	c.Assert(err, IsNil)
	noop(src)
}

func (*TestSuite) BenchmarkRandSource(c *C) {
	for i := 0; i < c.N; i++ {
		securerandom.RandSource()
	}
}
