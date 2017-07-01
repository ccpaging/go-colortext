/*
ct package provides functions to change the color of console text.

Under windows platform, the Console API is used. Under other systems, ANSI text mode is used.
*/
package ct

// Term is the type of terminal type to be set.
type TermType int

const (
	DumbTerm = TermType(iota)
	AnsiTerm
	WinTerm
)

// Color is the type of color to be set.
type Color int

const (
	// No change of color
	None = Color(iota)
	Black
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

type ctInterface interface {
	resetColor()
	changeColor(fg Color, fgBright bool, bg Color, bgBright bool)
}

var (
	Global ctInterface
)

func init() {
	switch GetTerminal() {
		case AnsiTerm:
			Global = NewAnsi()
		case WinTerm:
			Global = NewWin()
		default:
			Global = NewDumb()
	}
}

// ResetColor resets the foreground and background to original colors
func ResetColor() {
	Global.resetColor()
}

// ChangeColor sets the foreground and background colors. If the value of the color is None,
// the corresponding color keeps unchanged.
// If fgBright or bgBright is set true, corresponding color use bright color. bgBright may be
// ignored in some OS environment.
func ChangeColor(fg Color, fgBright bool, bg Color, bgBright bool) {
	Global.changeColor(fg, fgBright, bg, bgBright)
}

// Foreground changes the foreground color.
func Foreground(cl Color, bright bool) {
	ChangeColor(cl, bright, None, false)
}

// Background changes the background color.
func Background(cl Color, bright bool) {
	ChangeColor(None, false, cl, bright)
}
