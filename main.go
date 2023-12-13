// Package logic provides basic logic operations.
package logic

// Xor performs exclusive or (exclusive disjunction) logical operation.
// It returns true only when inputs differ (one is true, the other is false).
func Xor(a, b bool) bool {
	return (a || b) && (!a || !b)
}

// Or performs inclusive or logical operation.
// It returns true when either both or either premise is true.
func Or(a, b bool) bool {
	return a || b
}

// Eq performs equality logical operation.
// It returns true when both premises are equal.
func Eq(a, b bool) bool {
	return a == b
}

// And performs logical AND operation.
// It returns true only when both premises are true.
func And(a, b bool) bool {
	return a && b
}
