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

// getColorCode returns the ANSI escape sequence for the given color.
func getColorCode(color Color) string {
	switch color {
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
		return "\033[0m"
	}
}

// String returns the given string wrapped with ANSI color codes.
func String(c Color, s string) string {
	return getColorCode(c) + s + getColorCode(Reset)
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
// The color persists until ResetForegroundColor is called or the terminal is reset.
func ForgroundColor(c Color) {
	fmt.Print(getColorCode(c))
}

// ResetForegroundColor resets the terminal's foreground color to the default.
func ResetForegroundColor() {
	fmt.Print(getColorCode(Reset))
}
