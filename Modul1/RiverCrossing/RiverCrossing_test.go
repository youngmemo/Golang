package RiverCrossing

import "testing"

// Test function implemented in line with the Golang's testing tool
func TestViewState(t *testing.T) {
	iWantThis := "Nothing is in the boat"
	testThis := ViewState()
	if testThis != iWantThis {
		t.Errorf("Feil, fikk %q, ønsket %q.", testThis, iWantThis)
	}

}
