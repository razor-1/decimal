package misc_test

import (
	"testing"

	"github.com/razor-1/decimal"
	"github.com/razor-1/decimal/internal/test"
	"github.com/razor-1/decimal/misc"
)

func TestBig_NextMinus(t *testing.T) { test.NextMinus.Test(t) }
func TestBig_NextPlus(t *testing.T)  { test.NextPlus.Test(t) }

// func TestBig_Shift(t *testing.T)     { test.Shift.Test(t) }

func TestCmpTotal(t *testing.T) {
	for i, test := range [...]struct {
		x, y string
		r    int
	}{
		0: {"qnan", "snan", +1},
		1: {"400", "qnan", -1},
		2: {"snan", "snan", 0},
		3: {"snan", "1e+9", +1},
		4: {"qnan", "1e+12", +1},
		5: {"-qnan", "12", -1},
		6: {"12", "qnan", -1},
	} {
		x := new(decimal.Big)
		x.Context.OperatingMode = decimal.GDA
		x.SetString(test.x)

		y := new(decimal.Big)
		y.Context.OperatingMode = decimal.GDA
		y.SetString(test.y)

		if r := misc.CmpTotal(x, y); r != test.r {
			t.Fatalf("#%d: CmpTotal(%q, %q): got %d, wanted %d",
				i, x, y, r, test.r)
		}
	}
}
