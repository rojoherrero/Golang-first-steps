package functions

// Divide two numbers
func Divide(dividend, divider int) (division, modulo int) {
	division = dividend / divider
	modulo = dividend % divider
	return division, modulo
}
