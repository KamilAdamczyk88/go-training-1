// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Tests to see how each algorithm compare.
package main

import (
	"bytes"
	"testing"
)

var output bytes.Buffer
var in = assembleInputStream()
var find = []byte("elvis")
var repl = []byte("Elvis")

// Capture the time it takes to execute algorithm one.
func BenchmarkAlgorithmOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		output.Reset()
		algOne(in, find, repl, &output)
	}
}

// Capture the time it takes to execute algorithm two.
func BenchmarkAlgorithmTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		output.Reset()
		algTwo(in, find, repl, &output)
	}
}
