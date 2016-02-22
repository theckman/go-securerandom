// Copyright 2016 Tim Heckman. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package randomizer_test

import (
	"testing"

	"github.com/theckman/go-randomizer"

	. "gopkg.in/check.v1"
)

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func Test(t *testing.T) { TestingT(t) }

func (*TestSuite) TestGenerateRandomBytes(c *C) {
	// I have no fucking idea how to test this shit...
	var b []byte
	var err error

	b, err = randomizer.GenerateRandomBytes(10)
	c.Assert(err, IsNil)
	c.Check(len(b), Equals, 10) // ¯\_(ツ)_/¯
}

func (t *TestSuite) BenchmarkGenerateRandomBytesBy1(c *C) {
	for i := 0; i < c.N; i++ {
		randomizer.GenerateRandomBytes(1)
		c.SetBytes(1)
	}
}

func (t *TestSuite) BenchmarkGenerateRandomBytesBy64(c *C) {
	for i := 0; i < c.N; i++ {
		randomizer.GenerateRandomBytes(64)
		c.SetBytes(64)
	}
}

func (t *TestSuite) BenchmarkGenerateRandomBytesBy128(c *C) {
	for i := 0; i < c.N; i++ {
		randomizer.GenerateRandomBytes(128)
		c.SetBytes(128)
	}
}

func (*TestSuite) TestGenerateRandomBase64(c *C) {
	var s string
	var err error

	s, err = randomizer.GenerateRandomBase64(32)
	c.Assert(err, IsNil)
	c.Check(len(s), Equals, 32)
}

func (*TestSuite) TestGenerateRandomBase64URL(c *C) {
	var s string
	var err error

	s, err = randomizer.GenerateRandomBase64URL(32)
	c.Assert(err, IsNil)
	c.Check(len(s), Equals, 32)
}

func (t *TestSuite) BenchmarkGenerateRandomBase64(c *C) {
	var s string
	for i := 0; i < c.N; i++ {
		s, _ = randomizer.GenerateRandomBase64(16)
		c.SetBytes(int64(len(s)))
	}
}

func (t *TestSuite) BenchmarkGenerateRandomBase64URL(c *C) {
	var s string
	for i := 0; i < c.N; i++ {
		s, _ = randomizer.GenerateRandomBase64URL(16)
		c.SetBytes(int64(len(s)))
	}
}

func (t *TestSuite) BenchmarkGenerateRandomUint16(c *C) {
	for i := 0; i < c.N; i++ {
		randomizer.GenerateRandomUint16()
		c.SetBytes(2)
	}
}

func (t *TestSuite) BenchmarkGenerateRandomUint32(c *C) {
	for i := 0; i < c.N; i++ {
		randomizer.GenerateRandomUint32()
		c.SetBytes(4)
	}
}

func (t *TestSuite) BenchmarkGenerateRandomUint64(c *C) {
	for i := 0; i < c.N; i++ {
		randomizer.GenerateRandomUint64()
		c.SetBytes(8)
	}
}
