package paddington

func SecondSolution(item string, pad int) string {
	// aggregator and temp num variable
	var agg string
	var num []byte

	// pad logic
	padFn := func(n []byte) (padResult string) {
		for i := 0; i < pad-len(n); i++ {
			padResult += "0"
		}
		return padResult + string(num)
	}

	// rather than recursively iterating through each index of item, loop through each character
	for _, char := range item {
		if char >= 48 && char <= 57 {
			num = append(num, byte(char))
			continue
		}
		if len(num) != 0 {
			agg += padFn(num)
			num = nil
			num = []byte{}
		}
		agg += string(char)

	}
	if len(num) != 0 {
		agg += padFn(num)
	}
	return agg
}
