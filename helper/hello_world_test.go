package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello World"
	actual := HelloWorld()
	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestHelloWorldWithParameter(t *testing.T) {
	expected := "Hello Ayom"
	result := HelloWorldWithParamater("Ayom")
	if expected != result {
		t.Errorf("Expected: %s, Actual: %s", expected, result)
	}
}
