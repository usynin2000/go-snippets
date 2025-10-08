package testing_utils

import "testing"


func TestAbs(t *testing.T) {
	if val := Abs(-3.133); val != 3.133 {
		t.Errorf("Abs expected to be 3.133, go %f", val)
	}
}