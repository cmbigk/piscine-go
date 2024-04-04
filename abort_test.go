package piscine

import "testing"

var tests = []struct {
	a, b, c, d, e, median int
}{
	{2, 3, 8, 5, 7, 5},
}

func TestAbortTableDriven(t *testing.T) {
	for _, test := range tests {
		if out := Abort(test.a, test.b, test.c, test.d, test.e); out != test.median {
			t.Errorf("Abort(%v, %v, %v, %v, %v) = %v, want %v\n", test.a, test.b, test.c, test.d, test.e, out, test.median)
		} else {
			t.Logf("Abort(%v, %v, %v, %v, %v) = %v, want %v\n", test.a, test.b, test.c, test.d, test.e, out, test.median)
		}
	}
}

func BenchmarkAbort(b *testing.B) {
	for _, test := range tests {
		Abort(test.a, test.b, test.c, test.d, test.e)
	}
}
