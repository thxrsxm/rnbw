package rnbw

import (
	"bytes"
	"fmt"
	"testing"
)

// --- Color code tests ---

func TestGetFgColorCode(t *testing.T) {
	tests := []struct {
		name  string
		color Color
		want  string
	}{
		{"reset", Reset, "\033[0m"},
		{"red", Red, "\033[31m"},
		{"green", Green, "\033[32m"},
		{"yellow", Yellow, "\033[33m"},
		{"blue", Blue, "\033[34m"},
		{"purple", Purple, "\033[35m"},
		{"cyan", Cyan, "\033[36m"},
		{"white", White, "\033[37m"},
		{"gray", Gray, "\033[90m"},
		{"unknown defaults to reset", Color(99), "\033[0m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getFgColorCode(tt.color)
			if got != tt.want {
				t.Errorf("getFgColorCode(%d) = %q, want %q", tt.color, got, tt.want)
			}
		})
	}
}

func TestGetBgColorCode(t *testing.T) {
	tests := []struct {
		name  string
		color Color
		want  string
	}{
		{"reset", Reset, "\033[0m"},
		{"red", Red, "\033[41m"},
		{"green", Green, "\033[42m"},
		{"yellow", Yellow, "\033[43m"},
		{"blue", Blue, "\033[44m"},
		{"purple", Purple, "\033[45m"},
		{"cyan", Cyan, "\033[46m"},
		{"white", White, "\033[47m"},
		{"gray", Gray, "\033[100m"},
		{"unknown defaults to reset", Color(99), "\033[0m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getBgColorCode(tt.color)
			if got != tt.want {
				t.Errorf("getBgColorCode(%d) = %q, want %q", tt.color, got, tt.want)
			}
		})
	}
}

// --- Foreground string builder tests ---

func TestString(t *testing.T) {
	got := String(Red, "hello")
	want := "\033[31mhello\033[0m"
	if got != want {
		t.Errorf("String(Red, \"hello\") = %q, want %q", got, want)
	}
}

func TestStringf(t *testing.T) {
	got := Stringf(Green, "count: %d", 42)
	want := "\033[32mcount: 42\033[0m"
	if got != want {
		t.Errorf("Stringf(Green, ...) = %q, want %q", got, want)
	}
}

// --- Background string builder tests ---

func TestBgString(t *testing.T) {
	got := BgString(Blue, "hello")
	want := "\033[44mhello\033[0m"
	if got != want {
		t.Errorf("BgString(Blue, \"hello\") = %q, want %q", got, want)
	}
}

func TestBgStringf(t *testing.T) {
	got := BgStringf(Yellow, "val: %s", "test")
	want := "\033[43mval: test\033[0m"
	if got != want {
		t.Errorf("BgStringf(Yellow, ...) = %q, want %q", got, want)
	}
}

// --- Styled string builder tests ---

func TestStyledString(t *testing.T) {
	got := StyledString(White, Red, "alert")
	want := "\033[37m\033[41malert\033[0m"
	if got != want {
		t.Errorf("StyledString(White, Red, \"alert\") = %q, want %q", got, want)
	}
}

func TestStyledStringf(t *testing.T) {
	got := StyledStringf(Cyan, Purple, "n=%d", 7)
	want := "\033[36m\033[45mn=7\033[0m"
	if got != want {
		t.Errorf("StyledStringf(Cyan, Purple, ...) = %q, want %q", got, want)
	}
}

// --- Foreground writer tests ---

func TestFprint(t *testing.T) {
	var buf bytes.Buffer
	n, err := Fprint(&buf, Red, "hello")
	if err != nil {
		t.Fatalf("Fprint returned error: %v", err)
	}
	want := "\033[31mhello\033[0m"
	if buf.String() != want {
		t.Errorf("Fprint wrote %q, want %q", buf.String(), want)
	}
	if n != len(want) {
		t.Errorf("Fprint returned n=%d, want %d", n, len(want))
	}
}

func TestFprintln(t *testing.T) {
	var buf bytes.Buffer
	_, err := Fprintln(&buf, Green, "hello")
	if err != nil {
		t.Fatalf("Fprintln returned error: %v", err)
	}
	want := "\033[32mhello\033[0m\n"
	if buf.String() != want {
		t.Errorf("Fprintln wrote %q, want %q", buf.String(), want)
	}
}

func TestFprintf(t *testing.T) {
	var buf bytes.Buffer
	_, err := Fprintf(&buf, Blue, "n=%d", 5)
	if err != nil {
		t.Fatalf("Fprintf returned error: %v", err)
	}
	want := "\033[34mn=5\033[0m"
	if buf.String() != want {
		t.Errorf("Fprintf wrote %q, want %q", buf.String(), want)
	}
}

// --- Background writer tests ---

func TestBgFprint(t *testing.T) {
	var buf bytes.Buffer
	_, err := BgFprint(&buf, Red, "hello")
	if err != nil {
		t.Fatalf("BgFprint returned error: %v", err)
	}
	want := "\033[41mhello\033[0m"
	if buf.String() != want {
		t.Errorf("BgFprint wrote %q, want %q", buf.String(), want)
	}
}

func TestBgFprintln(t *testing.T) {
	var buf bytes.Buffer
	_, err := BgFprintln(&buf, Green, "hello")
	if err != nil {
		t.Fatalf("BgFprintln returned error: %v", err)
	}
	want := "\033[42mhello\033[0m\n"
	if buf.String() != want {
		t.Errorf("BgFprintln wrote %q, want %q", buf.String(), want)
	}
}

func TestBgFprintf(t *testing.T) {
	var buf bytes.Buffer
	_, err := BgFprintf(&buf, Blue, "n=%d", 5)
	if err != nil {
		t.Fatalf("BgFprintf returned error: %v", err)
	}
	want := "\033[44mn=5\033[0m"
	if buf.String() != want {
		t.Errorf("BgFprintf wrote %q, want %q", buf.String(), want)
	}
}

// --- Styled writer tests ---

func TestStyledFprint(t *testing.T) {
	var buf bytes.Buffer
	_, err := StyledFprint(&buf, White, Red, "alert")
	if err != nil {
		t.Fatalf("StyledFprint returned error: %v", err)
	}
	want := "\033[37m\033[41malert\033[0m"
	if buf.String() != want {
		t.Errorf("StyledFprint wrote %q, want %q", buf.String(), want)
	}
}

func TestStyledFprintln(t *testing.T) {
	var buf bytes.Buffer
	_, err := StyledFprintln(&buf, White, Red, "alert")
	if err != nil {
		t.Fatalf("StyledFprintln returned error: %v", err)
	}
	want := "\033[37m\033[41malert\033[0m\n"
	if buf.String() != want {
		t.Errorf("StyledFprintln wrote %q, want %q", buf.String(), want)
	}
}

func TestStyledFprintf(t *testing.T) {
	var buf bytes.Buffer
	_, err := StyledFprintf(&buf, Cyan, Purple, "n=%d", 7)
	if err != nil {
		t.Fatalf("StyledFprintf returned error: %v", err)
	}
	want := "\033[36m\033[45mn=7\033[0m"
	if buf.String() != want {
		t.Errorf("StyledFprintf wrote %q, want %q", buf.String(), want)
	}
}

// --- Edge cases ---

func TestStringEmpty(t *testing.T) {
	got := String(Red, "")
	want := "\033[31m\033[0m"
	if got != want {
		t.Errorf("String(Red, \"\") = %q, want %q", got, want)
	}
}

func TestStringReset(t *testing.T) {
	got := String(Reset, "hello")
	want := "\033[0mhello\033[0m"
	if got != want {
		t.Errorf("String(Reset, \"hello\") = %q, want %q", got, want)
	}
}

func TestStyledStringBothReset(t *testing.T) {
	got := StyledString(Reset, Reset, "hello")
	want := "\033[0m\033[0mhello\033[0m"
	if got != want {
		t.Errorf("StyledString(Reset, Reset, \"hello\") = %q, want %q", got, want)
	}
}

// --- Error propagation test ---

type errWriter struct{}

func (e errWriter) Write([]byte) (int, error) {
	return 0, fmt.Errorf("write failed")
}

func TestFprintError(t *testing.T) {
	_, err := Fprint(errWriter{}, Red, "hello")
	if err == nil {
		t.Error("Fprint should propagate writer errors")
	}
}

func TestBgFprintError(t *testing.T) {
	_, err := BgFprint(errWriter{}, Red, "hello")
	if err == nil {
		t.Error("BgFprint should propagate writer errors")
	}
}

func TestStyledFprintError(t *testing.T) {
	_, err := StyledFprint(errWriter{}, Red, Blue, "hello")
	if err == nil {
		t.Error("StyledFprint should propagate writer errors")
	}
}
