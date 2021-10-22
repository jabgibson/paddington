package paddington

import "testing"

// These tests will be testing the second solution... which is more of a procedural loop based solution. More traditional
// in today's software lens.
func TestSecondSolution(t *testing.T) {
	for _, tt := range wnpTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SecondSolution(tt.args.item, tt.args.pad); got != tt.want {
				t.Errorf("SecondSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
