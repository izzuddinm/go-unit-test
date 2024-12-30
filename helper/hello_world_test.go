package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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
	t.Log("TestHelloWorldWithParameter passed")
}

func TestSubTest(t *testing.T) {

	t.Run("HelloWorld", func(t *testing.T) {
		result := HelloWorld()
		require.Equal(t, "Hello World", result, "Result must be Hello World")
		t.Log("TestHelloWorld passed")
	})
}
