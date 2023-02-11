package calculator_app

import (
	"testing"
)

func TestAdd(t *testing.T) {
	exp := 5
	res := Addition(2, 3)
	if res != exp {
		t.Errorf("%d was expect but got %d .\n", exp, res)
	}
}
