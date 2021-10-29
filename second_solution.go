package paddington

func SecondSolution(item string, pad int) string {
	// aggregator and temp num variable
	var agg string
	var num NumBuilder

	// pad logic
	padFn := func(n []byte) (padResult string) {
		for i := 0; i < pad-len(n); i++ {
			padResult += "0"
		}
		return padResult + string(num)
	}

	for _, char := range item {
		switch char {
		case Zero, One, Two, Three, Four, Five, Six, Seven, Eight, Nine:
			num = append(num, byte(char))
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
			agg += string(char)
		}
		//
		//if char >= 48 && char <= 57 {
		//	num = append(num, byte(char))
		//	continue
		//}
		//if len(num) != 0 {
		//	agg += padFn(num)
		//	num = nil
		//	num = []byte{}
		//}
		//agg += string(char)

	}
	if len(num) != 0 {
		agg += padFn(num)
	}
	return agg
}
