package utils

// If is a generic ternary helper function
// T can be any data type (string, int, bool, etc.)
func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}
