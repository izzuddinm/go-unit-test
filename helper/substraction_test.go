package helper

import (
	"fmt"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	// Do something before running the test
	fmt.Println("Start test")
	m.Run()
	fmt.Println("End test")
	// Do something after running the test
}

func TestSubstration(t *testing.T) {
	expected := 2
	result := Substraction(4, 2)
	if runtime.GOOS == "windows" {
		if expected != result {
			t.Errorf("Expected: %d, Actual: %d", expected, result)
			t.Fail()
			t.FailNow()
			t.Skip("Skip this test")
			t.Log("TestSubstraction passed")
		}

		t.Skip("Can't run on windows")
	}
}
