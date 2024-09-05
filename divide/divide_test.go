package divide

import (
	"testing"
)

func FuzzDivide(f *testing.F) {
	f.Add(int64(1), int64(1))
	f.Fuzz(func(t *testing.T, a int64, b int64) {
		Divide(a, b)
	})
}
