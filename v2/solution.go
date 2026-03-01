// Package chapter10 contains the solution to the exercises in chapter 10 of Learning Go, 2nd edition.
package chapter10

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Adds two numeric values and returns their sum as defined in [mathisfun].
//
// [mathisfun]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
