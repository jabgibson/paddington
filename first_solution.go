package paddington

// FirstSolution is merely the functional transport to the recursiveLogic. Since
// the logic for this is recursive in nature, and this function is really only responsible
// for setting the data up for the recursive function to be invoked.
func FirstSolution(item string, padding int) string {
	return recursiveLogic([]byte(item), "", padding, []byte{})
}

// If this wasn't Go.. I would embed innerPad within FirstSolution, or even recursiveLogic...
// but Go has a constraint against calling internally defined lambdas within the same scope it was
// defined, so for this reason... it is top level. I could also add one more argument to recursiveLogic
// as a func that contains the padding logic, but that gets messy and we should respect that Go doesn't
// support recursion as a first class citizen.
func innerPad(num []byte, pad int) (x string) {
	for i := 0; i < pad-len(num); i++ {
		x += "0"
	}
	return x + string(num)
}

// This recursive function uses an aggregator pattern. This is not as elegant as would be found
// in something like Haskell or F#, but for what it does I think its elegant and simple in its own Go
// scoped right.
func recursiveLogic(item []byte, agg string, pad int, num NumBuilder) string {
	if len(item) == 0 {
		if len(num) != 0 {
			agg += innerPad(num, pad)
		}
		return agg
	}

	switch item[0] {
	case Zero, One, Two, Three, Four, Five, Six, Seven, Eight, Nine:
		num = append(num, item[0])
	case DecimalPoint:
		// If decimal point, lets pad the existing whole segment of number, prepend the number to the decimal point
		num = append([]byte(innerPad(num, pad)), DecimalPoint)
	default:
		if num.is() {
			// if the num has already been padded, because its a decimal, dont pad again, just agg it
			if num.isDecimal() {
				agg += string(num)
			} else {
				agg += innerPad(num, pad)
			}
			num = nil
			num = []byte{}
		}
		agg += string(item[0])
	}

	return recursiveLogic(item[1:], agg, pad, num)
}
