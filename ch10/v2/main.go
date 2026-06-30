// Package ch10 includes functions to do simple mathematical operations
package ch10

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add takes two Numbers and returns their sum.
// Addition rules are defined in [mathsisfun addition guide]
//
// [mathsisfun addition guide]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
