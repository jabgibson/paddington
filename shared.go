package paddington

const (
	DecimalPoint = 46
	Zero         = 48
	One          = 49
	Two          = 50
	Three        = 51
	Four         = 52
	Five         = 53
	Six          = 54
	Seven        = 55
	Eight        = 56
	Nine         = 57
)

type NumBuilder []byte

func (n NumBuilder) is() bool {
	return len(n) > 0
}

func (n NumBuilder) isDecimal() bool {
	for _, c := range n {
		if c == DecimalPoint {
			return true
		}
	}
	return false
}
