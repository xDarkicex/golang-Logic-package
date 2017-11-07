/*Package logic contains basic logic operations */
package logic

var (
	operation bool
	negation  bool
)

/*Xor exclusive or function
Exclusive or or exclusive disjunction is a logical operation that outputs
true only when inputs differ (one is true, the other is false)
*/
func Xor(a, b bool) bool {
	operation = (a || b)
	negation = (!a || !b)
	if And(operation, negation) {
		return true
	}
	return false
}

/*Or inclusive or function
basic || wrapped
Returns true when either both or either premise is true
*/
func Or(a, b bool) bool {
	if a || b {
		return true
	}
	return false
}

/*Eq equality function
Returns true when both premise are equal
*/
func Eq(a, b bool) bool {
	if a == b {
		return true
	}
	return false
}

//And function only returns true when both premise are true.
func And(a, b bool) bool {
	if a && b {
		return true
	}
	return false
}
