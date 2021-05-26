package main

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/1e9y/chalk"
	"github.com/fatih/color"
	. "github.com/klauspost/cpuid/v2"
)

func main() {
	fmt.Printf("goos: %s\n", runtime.GOOS)
	fmt.Printf("goarch: %s\n", runtime.GOARCH)
	fmt.Printf("cpu: %s\n", CPU.BrandName)

	br := testing.Benchmark(func(b *testing.B) {
		c := color.New(color.BgRed, color.FgBlack)
		for i := 0; i < b.N; i++ {
			_ = c.Sprint("hello")
		}
		b.ReportAllocs()
	})
	println("fatih/color", br.String(), br.MemString())

	br = testing.Benchmark(func(b *testing.B) {
		c := chalk.NewChalk(chalk.BGRed, chalk.BGRed)
		for i := 0; i < b.N; i++ {
			_ = c.Sprint("hello")
		}
		b.ReportAllocs()
	})
	println("1e9y/chalk ", br.String(), br.MemString())

}
