package chalk

import (
	"fmt"
	"testing"
)

func TestChalk(t *testing.T) {
	Red("Hello, Red!\n")
	Green("Hello, Green!\n")

	var c Chalk
	for i := 30; i <= 37; i++ {
		for j := 40; j <= 47; j++ {
			c = Chalk{}
			c.Add(Parameter(i)).Add(Parameter(j))
			c.Printf("%3d", i)
		}
		fmt.Println()
	}
}
