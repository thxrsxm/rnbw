package rnbw

import "fmt"

// Color represents a terminal color code for ANSI-colored output.
type Color int

// Color constants define ANSI color codes for terminal output.
const (
	Reset  Color = iota // Resets the terminal color.
	Red                 // Red text color.
	Green               // Green text color.
	Yellow              // Yellow text color.
	Blue                // Blue text color.
	Purple              // Purple text color.
	Cyan                // Cyan text color.
	White               // White text color.
	Gray                // Gray text color.
)

const resetColorCode = "\033[0m"

// getFgColorCode returns the ANSI escape sequence for the given foreground color.
func getFgColorCode(c Color) string {
	switch c {
	case Reset:
		return "\033[0m"
	case Red:
		return "\033[31m"
	case Green:
		return "\033[32m"
	case Yellow:
		return "\033[33m"
	case Blue:
		return "\033[34m"
	case Purple:
		return "\033[35m"
	case Cyan:
		return "\033[36m"
	case White:
		return "\033[37m"
	case Gray:
		return "\033[90m"
	default:
		// Default to reset code for unknown colors.
		return resetColorCode
	}
}

func getBgColorCode(c Color) string {
	switch c {
	case Reset:
		return "\033[0m"
	case Red:
		return "\033[41m"
	case Green:
		return "\033[42m"
	case Yellow:
		return "\033[43m"
	case Blue:
		return "\033[44m"
	case Purple:
		return "\033[45m"
	case Cyan:
		return "\033[46m"
	case White:
		return "\033[47m"
	case Gray:
		return "\033[100m"
	default:
		// Default to reset code for unknown colors.
		return resetColorCode
	}
}

// String returns the given string wrapped with ANSI color codes.
func String(c Color, s string) string {
	return getFgColorCode(c) + s + resetColorCode
}

// Stringf returns a formatted string wrapped with ANSI color codes.
// It behaves like fmt.Sprintf but adds color to the result.
func Stringf(c Color, f string, a ...any) string {
	return String(c, fmt.Sprintf(f, a...))
}

// Print prints a colored string to standard output.
// It returns the number of bytes written and any write error encountered.
func Print(c Color, s string) (n int, err error) {
	return fmt.Print(String(c, s))
}

// Println prints a colored string to standard output followed by a newline.
// It returns the number of bytes written and any write error encountered.
func Println(c Color, s string) (n int, err error) {
	return fmt.Println(String(c, s))
}

// Printf prints a formatted colored string to standard output.
// It behaves like fmt.Printf but adds color to the output.
// It returns the number of bytes written and any write error encountered.
func Printf(c Color, f string, a ...any) (n int, err error) {
	return fmt.Print(Stringf(c, f, a...))
}

// ForegroundColor sets the terminal's foreground color to the specified color.
// The color persists until ResetColor is called or the terminal is reset.
func ForgroundColor(c Color) {
	fmt.Print(getFgColorCode(c))
}

// BgString returns the given string with a colored background.
func BgString(c Color, s string) string {
	return getBgColorCode(c) + s + resetColorCode
}

// BgStringf returns a formatted string with a colored background.
// It behaves like fmt.Sprintf but adds a background color to the result.
func BgStringf(c Color, f string, a ...any) string {
	return BgString(c, fmt.Sprintf(f, a...))
}

// BgPrint prints a string with a colored background to standard output.
// It returns the number of bytes written and any write error encountered.
func BgPrint(c Color, s string) (n int, err error) {
	return fmt.Print(BgString(c, s))
}

// BgPrintln prints a string with a colored background to standard output followed by a newline.
// It returns the number of bytes written and any write error encountered.
func BgPrintln(c Color, s string) (n int, err error) {
	return fmt.Println(BgString(c, s))
}

// BgPrintf prints a formatted string with a colored background to standard output.
// It behaves like fmt.Printf but adds a background color to the output.
// It returns the number of bytes written and any write error encountered.
func BgPrintf(c Color, f string, a ...any) (n int, err error) {
	return fmt.Print(BgStringf(c, f, a...))
}

// BackgroundColor sets the terminal's background color to the specified color.
// The color persists until ResetColor is called or the terminal is reset.
func BackgroundColor(c Color) {
	fmt.Print(getBgColorCode(c))
}

// StyledString returns a string with both foreground and background colors.
func StyledString(fg Color, bg Color, s string) string {
	return getFgColorCode(fg) + getBgColorCode(bg) + s + resetColorCode
}

// StyledStringf returns a formatted string with both foreground and background colors.
func StyledStringf(fg Color, bg Color, f string, a ...any) string {
	return StyledString(fg, bg, fmt.Sprintf(f, a...))
}

// StyledPrint prints a string with both foreground and background colors.
func StyledPrint(fg Color, bg Color, s string) (n int, err error) {
	return fmt.Print(StyledString(fg, bg, s))
}

// StyledPrintln prints a string with both foreground and background colors followed by a newline.
func StyledPrintln(fg Color, bg Color, s string) (n int, err error) {
	return fmt.Println(StyledString(fg, bg, s))
}

// StyledPrintf prints a formatted string with both foreground and background colors.
func StyledPrintf(fg Color, bg Color, f string, a ...any) (n int, err error) {
	return fmt.Print(StyledStringf(fg, bg, f, a...))
}

// ResetColor resets the terminal's foreground and background color to the default.
func ResetColor() {
	fmt.Print(resetColorCode)
}
