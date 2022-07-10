package ch2

// Explicit type conversion
var x int = 10
var y float64 = 30.2
var z float64 = float64(x) + y
var d int = x + int(y)

// Declaration list
var (
	a int
	b = 20
)

func example() int {
	// := is short-hand for declaring a variable. Can only be used in functions.
	c := 30
	var d = 10
	return c + d
}

// For declaraing a variable is immutable
// - Only works for static values (literals, strings, etc.)
const foobar = 10
