package readline

// ascii
const (
	// Enter move the line start
	Enter = "\r"
	// Backspace like backspace, move the cursor if it's the line enter,
	// example: "aaa\b\b" -> "aaa"; "aa\bc" -> "ac"
	Backspace = "\b"
)

// https://en.wikipedia.org/wiki/ANSI_escape_code
const (

	// EcsJ Clears part of the screen.
	// If n is 0 (or missing), clear from cursor to end of screen.
	// If n is 1, clear from cursor to beginning of the screen.
	// If n is 2, clear entire screen (and moves cursor to upper left on DOS ANSI.SYS).
	// If n is 3, clear entire screen and delete all lines saved in the scrollback buffer (this feature was added for xterm and is supported by other terminal applications).
	EcsJ = "\033[J"
	// EcsK Erases part of the line.
	// If n is 0 (or missing), clear from cursor to the end of the line.
	// If n is 1, clear from cursor to beginning of the line.
	// If n is 2, clear entire line. Cursor position does not change.
	EcsK = "\033[K"

	// EcsA Cursor Up
	// Moves the cursor n (default 1) cells in the given direction.
	// If the cursor is already at the edge of the screen, this has no effect.
	EcsA = "\033[A"
	// EcsB Cursor Down
	EcsB = "\033[B"
	// EcsC Cursor Forward
	EcsC = "\033[C"
	// EcsD Cursor Back
	EcsD = "\033[D"
)
