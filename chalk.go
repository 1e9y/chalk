package chalk

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/mattn/go-isatty"
)

var (
	// IsTerminal defines whether the output is a terminal and supports colorized output.
	// A more accurate way to determine this would be to check if `tput color` is greater than 8.
	IsTerminal = os.Getenv("TERM") != "dumb" && isatty.IsTerminal(os.Stdout.Fd())

	// HasColors turns the colorized output on. If HasColors is not set explicitly, it will
	// default to terminal capability.
	HasColors = IsTerminal
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

func (p Parameter) isFGColor() bool {
	return (p >= 30 && p <= 37) || (p >= 90 && p <= 97)
}

func (p Parameter) isBGColor() bool {
	return (p >= 40 && p <= 47) || (p >= 100 && p <= 107)
}

type Chalk struct {
	enabled    *bool
	fgColor    Parameter
	bgColor    Parameter
	attributes []Parameter
}

const csi = "\x1b["

func (c *Chalk) Enable() {
	c.enabled = func() *bool { t := true; return &t }()
}

func (c *Chalk) Disable() {
	c.enabled = func() *bool { f := false; return &f }()
}

func (c *Chalk) isEnabled() bool {
	if c.enabled == nil {
		return HasColors
	}
	return *c.enabled
}

func (c *Chalk) sequence() string {
	seq := make([]string, len(c.attributes))
	for i, p := range c.attributes {
		seq[i] = strconv.Itoa(int(p))
	}
	if c.fgColor > 0 {
		seq = append(seq, strconv.Itoa(int(c.fgColor)))
	}
	if c.bgColor > 0 {
		seq = append(seq, strconv.Itoa(int(c.bgColor)))
	}
	return strings.Join(seq, ";") + "m"
}

//func (c *Chalk) sequence() string {
//	var out strings.Builder
//	for i, p := range c.attributes {
//		out.Write([]byte(strconv.Itoa(int(p))))
//		if i != 0 {
//			out.WriteRune(';')
//		}
//	}
//	if c.fgColor > 0 {
//		if out.Len() > 0 {
//			out.WriteRune(';')
//		}
//		out.WriteString(strconv.Itoa(int(c.fgColor)))
//	}
//	if c.bgColor > 0 {
//		if out.Len() > 0 {
//			out.WriteRune(';')
//		}
//		out.WriteString(strconv.Itoa(int(c.bgColor)))
//	}
//	out.WriteRune('m')
//	return out.String()
//}

func (c *Chalk) set(w io.Writer) {
	if !c.isEnabled() {
		return
	}
	fmt.Fprint(w, csi+c.sequence())
}

func (c *Chalk) unset(w io.Writer) {
	if !c.isEnabled() {
		return
	}
	fmt.Fprint(w, csi+"0m")
}

func (c *Chalk) wrap(s string) string {
	if !c.isEnabled() {
		return s
	}
	return csi + c.sequence() + s + csi + "0m"
}

func NewChalk(p ...Parameter) *Chalk {
	c := Chalk{attributes: p}
	return &c
}

func (c *Chalk) Add(color Parameter) *Chalk {
	if color.isFGColor() {
		c.fgColor = color
	} else if color.isBGColor() {
		c.bgColor = color
	} else {
		c.attributes = append(c.attributes, color)
	}
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
