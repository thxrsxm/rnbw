package rnbw

import (
	"fmt"
	"io"
	"os"
)

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

// resetColorCode represents the terminal reset color code.
const resetColorCode = "\033[0m"

// getFgColorCode returns the ANSI escape sequence for the given foreground color.
func getFgColorCode(c Color) string {
	switch c {
	case Reset:
		return resetColorCode
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

// getBgColorCode returns the ANSI escape sequence for the given background color.
func getBgColorCode(c Color) string {
	switch c {
	case Reset:
		return resetColorCode
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

// --- Foreground string builders ---

// String returns the given string wrapped with foreground ANSI color codes.
func String(c Color, s string) string {
	return getFgColorCode(c) + s + resetColorCode
}

// Stringf returns a formatted string wrapped with foreground ANSI color codes.
func Stringf(c Color, f string, a ...any) string {
	return String(c, fmt.Sprintf(f, a...))
}

// --- Foreground print functions ---

// Fprint writes a colored string to the given writer.
func Fprint(w io.Writer, c Color, s string) (int, error) {
	return fmt.Fprint(w, String(c, s))
}

// Fprintln writes a colored string followed by a newline to the given writer.
func Fprintln(w io.Writer, c Color, s string) (int, error) {
	return fmt.Fprintln(w, String(c, s))
}

// Fprintf writes a formatted colored string to the given writer.
func Fprintf(w io.Writer, c Color, f string, a ...any) (int, error) {
	return fmt.Fprint(w, Stringf(c, f, a...))
}

// Print prints a colored string to standard output.
func Print(c Color, s string) (int, error) {
	return Fprint(os.Stdout, c, s)
}

// Println prints a colored string followed by a newline to standard output.
func Println(c Color, s string) (int, error) {
	return Fprintln(os.Stdout, c, s)
}

// Printf prints a formatted colored string to standard output.
func Printf(c Color, f string, a ...any) (int, error) {
	return Fprintf(os.Stdout, c, f, a...)
}

// --- Background string builders ---

// BgString returns the given string with a colored background.
func BgString(c Color, s string) string {
	return getBgColorCode(c) + s + resetColorCode
}

// BgStringf returns a formatted string with a colored background.
func BgStringf(c Color, f string, a ...any) string {
	return BgString(c, fmt.Sprintf(f, a...))
}

// --- Background print functions ---

// BgFprint writes a string with a colored background to the given writer.
func BgFprint(w io.Writer, c Color, s string) (int, error) {
	return fmt.Fprint(w, BgString(c, s))
}

// BgFprintln writes a string with a colored background followed by a newline to the given writer.
func BgFprintln(w io.Writer, c Color, s string) (int, error) {
	return fmt.Fprintln(w, BgString(c, s))
}

// BgFprintf writes a formatted string with a colored background to the given writer.
func BgFprintf(w io.Writer, c Color, f string, a ...any) (int, error) {
	return fmt.Fprint(w, BgStringf(c, f, a...))
}

// BgPrint prints a string with a colored background to standard output.
func BgPrint(c Color, s string) (int, error) {
	return BgFprint(os.Stdout, c, s)
}

// BgPrintln prints a string with a colored background followed by a newline to standard output.
func BgPrintln(c Color, s string) (int, error) {
	return BgFprintln(os.Stdout, c, s)
}

// BgPrintf prints a formatted string with a colored background to standard output.
func BgPrintf(c Color, f string, a ...any) (int, error) {
	return BgFprintf(os.Stdout, c, f, a...)
}

// --- Styled string builders ---

// StyledString returns a string with both foreground and background colors.
func StyledString(fg Color, bg Color, s string) string {
	return getFgColorCode(fg) + getBgColorCode(bg) + s + resetColorCode
}

// StyledStringf returns a formatted string with both foreground and background colors.
func StyledStringf(fg Color, bg Color, f string, a ...any) string {
	return StyledString(fg, bg, fmt.Sprintf(f, a...))
}

// --- Styled print functions ---

// StyledFprint writes a string with both foreground and background colors to the given writer.
func StyledFprint(w io.Writer, fg Color, bg Color, s string) (int, error) {
	return fmt.Fprint(w, StyledString(fg, bg, s))
}

// StyledFprintln writes a string with both foreground and background colors followed by a newline to the given writer.
func StyledFprintln(w io.Writer, fg Color, bg Color, s string) (int, error) {
	return fmt.Fprintln(w, StyledString(fg, bg, s))
}

// StyledFprintf writes a formatted string with both foreground and background colors to the given writer.
func StyledFprintf(w io.Writer, fg Color, bg Color, f string, a ...any) (int, error) {
	return fmt.Fprint(w, StyledStringf(fg, bg, f, a...))
}

// StyledPrint prints a string with both foreground and background colors to standard output.
func StyledPrint(fg Color, bg Color, s string) (int, error) {
	return StyledFprint(os.Stdout, fg, bg, s)
}

// StyledPrintln prints a string with both foreground and background colors followed by a newline to standard output.
func StyledPrintln(fg Color, bg Color, s string) (int, error) {
	return StyledFprintln(os.Stdout, fg, bg, s)
}

// StyledPrintf prints a formatted string with both foreground and background colors to standard output.
func StyledPrintf(fg Color, bg Color, f string, a ...any) (int, error) {
	return StyledFprintf(os.Stdout, fg, bg, f, a...)
}

// --- Persistent color functions ---

// ForegroundColor sets the terminal's foreground color to the specified color.
// The color persists until ResetColor is called or the terminal is reset.
func ForegroundColor(c Color) {
	fmt.Print(getFgColorCode(c))
}

// BackgroundColor sets the terminal's background color to the specified color.
// The color persists until ResetColor is called or the terminal is reset.
func BackgroundColor(c Color) {
	fmt.Print(getBgColorCode(c))
}

// ResetColor resets the terminal's foreground and background color to the default.
func ResetColor() {
	fmt.Print(resetColorCode)
}
