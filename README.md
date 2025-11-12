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

import "github.com/thxrsxm/rnbw"

func main() {
    // Print colored text (foreground)
    rnbw.Println(rnbw.Green, "Success!")
    rnbw.Println(rnbw.Red, "Error!")
    
    // Print with background color
    rnbw.BgPrintln(rnbw.Yellow, "Warning background")
    
    // Combine foreground and background
    rnbw.StyledPrintln(rnbw.White, rnbw.Red, "White text on red background")
    
    // Format and color
    rnbw.Printf(rnbw.Blue, "Hello, %s!\n", "World")
    rnbw.StyledPrintf(rnbw.Yellow, rnbw.Blue, "Score: %d\n", 100)
    
    // Get colored strings
    message := rnbw.String(rnbw.Yellow, "Warning")
    styled := rnbw.StyledString(rnbw.Green, rnbw.White, "Success!")
    
    // Set persistent colors
    rnbw.ForegroundColor(rnbw.Cyan)
    fmt.Println("This text is cyan")
    rnbw.BackgroundColor(rnbw.Blue)
    fmt.Println("This has a blue background")
    rnbw.resetColor()
}
```

## Available Colors

- Reset - Reset to default
- Red
- Green
- Yellow
- Blue
- Purple
- Cyan
- White
- Gray

## License

MIT License
