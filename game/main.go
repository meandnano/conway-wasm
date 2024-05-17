package main

// This calls a JS function from Go.
func main() {
	println("mult two numbers:", multiply(2, 3))
}

// exports.multiply() in JavaScript.
//
//export multiply
func multiply(x, y int) int {
	return x * y
}
