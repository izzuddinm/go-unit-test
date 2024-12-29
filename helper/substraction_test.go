package helper

import "testing"

func TestSubstration(t *testing.T) {
	expected := 2
	result := Substraction(4, 2)
	if expected != result {
		t.Errorf("Expected: %d, Actual: %d", expected, result)
		t.Fail()
		t.FailNow()
	}

	t.Log("TestSubstraction passed")
}
