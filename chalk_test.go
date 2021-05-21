package chalk

import (
	"fmt"
	"testing"
)

func init() {
	// Force tests to print escape sequence regardless of terminal capabilities.
	HasColors = true
}

var codes = []struct {
	p    Parameter
	code int
	name string
}{
	{FGBlack, 30, "black"},
	{FGRed, 31, "red"},
	{FGGreen, 32, "green"},
	{FGYellow, 33, "yellow"},
	{FGBlue, 34, "blue"},
	{FGMagenta, 35, "magenta"},
	{FGCyan, 36, "cyan"},
	{FGWhite, 37, "white"},

	{FGGray, 90, "gray"},
	{FGBrightBlack, 90, "brightblack"},
	{FGBrightRed, 91, "brightred"},
	{FGBrightGreen, 92, "brightgreen"},
	{FGBrightYellow, 93, "brightyellow"},
	{FGBrightBlue, 94, "brightblue"},
	{FGBrightMagenta, 95, "brightmagenta"},
	{FGBrightCyan, 96, "brightcyan"},
	{FGBrightWhite, 97, "brightwhite"},
}

func TestAddColors(t *testing.T) {
	// TODO: Use inner loop to add variable number of colors to the chalk
	for i := 2; i < len(codes); i += 3 {
		chalk := NewChalk()
		chalk.Add(codes[i-2].p).Add(codes[i-1].p).Add(codes[i].p)
		want := fmt.Sprintf("\x1b[32m%s\x1b[0m", codes[i].name)
		if s := chalk.Sprint(codes[i].name); s != want {
			t.Errorf("got:%q\nwant:%s", s, want)
		}
	}
}

func TestChalk_Sprint(t *testing.T) {
	for _, test := range codes {
		chalk := NewChalk()
		chalk.Add(test.p)
		want := fmt.Sprintf("\x1b[%dm%s\x1b[0m", test.code, test.name)
		if s := chalk.Sprint(test.name); s != want {
			t.Errorf("Sprint(%s) = %q, want %q", test.name, s, want)
		}
	}
}

func TestChalk_Sprintf(t *testing.T) {
	for _, test := range codes {
		chalk := NewChalk()
		chalk.Add(test.p)
		want := fmt.Sprintf("\x1b[%dmhello, %s!\x1b[0m", test.code, test.name)
		if s := chalk.Sprintf("hello, %s!", test.name); s != want {
			t.Errorf("Sprint(%s) = %q, want %q", test.name, s, want)
		}
	}
}

func TestChalk_Sprintln(t *testing.T) {
	for _, test := range codes {
		chalk := NewChalk()
		chalk.Add(test.p)
		want := fmt.Sprintf("\x1b[%dm%s\n\x1b[0m", test.code, test.name)
		if s := chalk.Sprintln(test.name); s != want {
			t.Errorf("Sprintln(%s) = %q, want %q", test.name, s, want)
		}
	}
}

func ExampleChalk_Print() {
	// TODO: Explore ways of capturing stdout
	for _, test := range codes[:8] {
		chalk := NewChalk()
		chalk.Add(test.p)
		chalk.Println(test.name)
	}
	// Output:
	// [30mblack
	// [0m[31mred
	// [0m[32mgreen
	// [0m[33myellow
	// [0m[34mblue
	// [0m[35mmagenta
	// [0m[36mcyan
	// [0m[37mwhite
	// [0m
}

func ExampleChalk_Printf() {
	for _, test := range codes[:8] {
		chalk := NewChalk()
		chalk.Add(test.p)
		chalk.Printf("hello, %s!", test.name)
		fmt.Println()
	}
	// Output:
	// [30mhello, black![0m
	// [31mhello, red![0m
	// [32mhello, green![0m
	// [33mhello, yellow![0m
	// [34mhello, blue![0m
	// [35mhello, magenta![0m
	// [36mhello, cyan![0m
	// [37mhello, white![0m
}

func ExampleRed() {
	Red("red")
	fmt.Println()
	Red("hello, %s!", "red")
	// Output:
	// [31mred[0m
	// [31mhello, red![0m
}

func ExampleGreen() {
	Green("green")
	fmt.Println()
	Green("hello, %s!", "green")
	// Output:
	// [32mgreen[0m
	// [32mhello, green![0m
}
