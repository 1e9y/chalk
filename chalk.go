package chalk

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// TODO: Explore performance impact of using `uint8`
type Parameter int

const (
	Reset Parameter = iota
	Bold
	Faint
	Italic
	Underline
	Blink
	_
	Reverse
	Conceal
	Strike
)

const (
	FGBlack Parameter = iota + 30
	FGRed
	FGGreen
	FGYellow
	FGBlue
	FGMagenta
	FGCyan
	FGWhite
)

const (
	FGBrightBlack Parameter = iota + 90
	FGBrightRed
	FGBrightGreen
	FGBrightYellow
	FGBrightBlue
	FGBrightMagenta
	FGBrightCyan
	FGBrightWhite
)

const FGGray = FGBrightBlack

const (
	BGBlack Parameter = iota + 40
	BGRed
	BGGreen
	BGYellow
	BGBlue
	BGMagenta
	BGCyan
	BGWhite
)

const (
	BGBrightBlack = iota + 100
	BGBrightRed
	BGBrightGreen
	BGBrightYellow
	BGBrightBlue
	BGBrightMagenta
	BGBrightCyan
	BGBrightWhite
)

const BGGray = BGBrightBlack

type Chalk struct {
	params []Parameter
}

const csi = "\x1b["

func (c *Chalk) sequence() string {
	seq := make([]string, len(c.params))
	for i, p := range c.params {
		seq[i] = strconv.Itoa(int(p))
	}
	return strings.Join(seq, ";")
}

func (c *Chalk) set(w io.Writer) {
	fmt.Fprint(w, csi+c.sequence()+"m")
}

func (c *Chalk) unset(w io.Writer) {
	fmt.Fprint(w, csi+"0m")
}

func (c *Chalk) wrap(s string) string {
	return fmt.Sprintf("%s%s%s%s", csi, c.sequence(), s, csi+"0m")
}

func NewChalk(p ...Parameter) *Chalk {
	c := Chalk{params: p}
	return &c
}

func (c *Chalk) Add(color Parameter) *Chalk {
	c.params = append(c.params, color)
	return c
}

func (c *Chalk) Print(a ...interface{}) (n int, err error) {
	c.set(os.Stdout)
	defer c.unset(os.Stdout)
	return fmt.Print(a...)
}

func (c *Chalk) Println(a ...interface{}) (n int, err error) {
	c.set(os.Stdout)
	defer c.unset(os.Stdout)
	return fmt.Println(a...)
}

// TODO: Force printf-like function linting
func (c *Chalk) Printf(format string, a ...interface{}) (n int, err error) {
	c.set(os.Stdout)
	defer c.unset(os.Stdout)
	return fmt.Printf(format, a...)
}

func (c *Chalk) Sprint(a ...interface{}) string {
	return c.wrap(fmt.Sprint(a...))
}

func (c *Chalk) Sprintln(a ...interface{}) string {
	return c.wrap(fmt.Sprintln(a...))
}

func (c *Chalk) Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	c.set(w)
	defer c.unset(w)
	return fmt.Fprint(w, a...)
}

func (c *Chalk) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	c.set(w)
	defer c.unset(w)
	return fmt.Fprintf(w, format, a...)
}

func (c *Chalk) Sprintf(format string, a ...interface{}) string {
	return c.wrap(fmt.Sprintf(format, a...))
}

func Red(format string, a ...interface{}) (int, error) {
	c := NewChalk(FGRed)
	return c.Printf(format, a...)
}

func Green(format string, a ...interface{}) (int, error) {
	c := NewChalk(FGGreen)
	return c.Printf(format, a...)
}
