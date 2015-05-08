package main

import (
	"testing"
	"github.com/williamsodell/goBallClockgo/helpers"
)

func BenchmarkMode1(b *testing.B) {
	for i := 0; i < b.N; i++ {
  	helpers.Mode1(123)
  }
}
