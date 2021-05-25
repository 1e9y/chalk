package chalk

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
