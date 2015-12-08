package packago

import "testing"

func TestConcat(t *testing.T) {
	concated := Concat("This is ", "a pen.")

	if concated != "This is a pen." {
		t.Errorf("concated want (%s)", "This is a pen.")
	}
}
