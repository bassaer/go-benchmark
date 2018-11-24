package main

import (
	"testing"

	"github.com/bassaer/randgen"
)

const times = 100000

func create(c int) []Param {
	p := []Param{}
	for i := 0; i < c; i++ {
		n := randgen.RandNum(1000000000, 9999999999)
		m := randgen.RandStr(10)
		p = append(p, Param{n, m})
	}
	return p
}

func BenchmarkFmt(b *testing.B) {
	p := create(times)
	b.ResetTimer()
	for i := 0; i < times; i++ {
		_ = JoinFmt(&p[i])
		//time.Sleep(1 * time.Millisecond)
	}
}

func BenchmarkPlus(b *testing.B) {
	p := create(times)
	b.ResetTimer()
	for i := 0; i < times; i++ {
		_ = JoinPlus(&p[i])
	}
}

func BenchmarkBuffer(b *testing.B) {
	p := create(times)
	b.ResetTimer()
	for i := 0; i < times; i++ {
		_ = JoinBuffer(&p[i])
	}
}
