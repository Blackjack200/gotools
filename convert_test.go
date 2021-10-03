package gotools

import (
	"fmt"
	"testing"

	"github.com/blackjack200/gotools/convert"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		return
	}
	t.Fatal(fmt.Sprintf("%v != %v", a, b))
}

func TestWrap(t *testing.T) {
	orig := uint8(1)
	val := convert.Wrap(orig)
	assertEqual(t, val.Uint8(), orig)
}
