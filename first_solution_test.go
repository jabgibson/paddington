package paddington

import "testing"

// This run of tests will test the crafty, somewhat unwieldy (in Go) recursive solution
func TestFirstSolution(t *testing.T) {
	for _, tt := range wnpTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstSolution(tt.args.item, tt.args.pad); got != tt.want {
				t.Errorf("FirstSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
