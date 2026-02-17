# rnbw 🌈

A lightweight, simple to use Go library for adding ANSI colors to your terminal output with support for both foreground and background colors.

## Features

- Simple and intuitive
- Support for 8 standard terminal colors
- Both foreground and background color support
- Drop-in replacements for `fmt.Print`, `fmt.Println`, and `fmt.Printf`
- Zero dependencies (only standard library)
- Lightweight and fast
- Combine foreground and background colors

## Installation

```bash
go get github.com/thxrsxm/rnbw
```

## Quick Start

```go
package main

import (
	"fmt"

	"github.com/thxrsxm/rnbw"
)

func main() {
	// Foreground colors
	rnbw.Println(rnbw.Green, "Success!")
	rnbw.Println(rnbw.Red, "Error!")
	rnbw.Printf(rnbw.Blue, "Hello, %s!\n", "World")
	fmt.Println()

	// Background colors
	rnbw.BgPrintln(rnbw.Yellow, "Warning background")
	fmt.Println()

	// Combined foreground and background
	rnbw.StyledPrintln(rnbw.Gray, rnbw.Red, "Gray text on red background")
	fmt.Println()
	rnbw.StyledPrintf(rnbw.Yellow, rnbw.Blue, "Score: %d", 100)
	fmt.Println()
}
```

## API

### String Builders

Return colored strings without printing.

| Function | Description |
|---|---|
| `String(c, s)` | Foreground colored string |
| `Stringf(c, f, a...)` | Formatted foreground colored string |
| `BgString(c, s)` | Background colored string |
| `BgStringf(c, f, a...)` | Formatted background colored string |
| `StyledString(fg, bg, s)` | Foreground + background colored string |
| `StyledStringf(fg, bg, f, a...)` | Formatted foreground + background colored string |

### Print to stdout

| Function | Description |
|---|---|
| `Print(c, s)` | Print with foreground color |
| `Println(c, s)` | Print with foreground color + newline |
| `Printf(c, f, a...)` | Formatted print with foreground color |
| `BgPrint(c, s)` | Print with background color |
| `BgPrintln(c, s)` | Print with background color + newline |
| `BgPrintf(c, f, a...)` | Formatted print with background color |
| `StyledPrint(fg, bg, s)` | Print with foreground + background |
| `StyledPrintln(fg, bg, s)` | Print with foreground + background + newline |
| `StyledPrintf(fg, bg, f, a...)` | Formatted print with foreground + background |

### Print to io.Writer

| Function | Description |
|---|---|
| `Fprint(w, c, s)` | Write with foreground color |
| `Fprintln(w, c, s)` | Write with foreground color + newline |
| `Fprintf(w, c, f, a...)` | Formatted write with foreground color |
| `BgFprint(w, c, s)` | Write with background color |
| `BgFprintln(w, c, s)` | Write with background color + newline |
| `BgFprintf(w, c, f, a...)` | Formatted write with background color |
| `StyledFprint(w, fg, bg, s)` | Write with foreground + background |
| `StyledFprintln(w, fg, bg, s)` | Write with foreground + background + newline |
| `StyledFprintf(w, fg, bg, f, a...)` | Formatted write with foreground + background |

### Persistent Colors

| Function | Description |
|---|---|
| `ForegroundColor(c)` | Set persistent foreground color |
| `BackgroundColor(c)` | Set persistent background color |
| `ResetColor()` | Reset to default terminal colors |

## Available Colors

| Color | Constant |
|---|---|
| Reset | `rnbw.Reset` |
| Red | `rnbw.Red` |
| Green | `rnbw.Green` |
| Yellow | `rnbw.Yellow` |
| Blue | `rnbw.Blue` |
| Purple | `rnbw.Purple` |
| Cyan | `rnbw.Cyan` |
| White | `rnbw.White` |
| Gray | `rnbw.Gray` |

## License

MIT License
