package basic

import (
	"reflect"
	"testing"
)

func TestVariable(t *testing.T) {

	t.Run("initialization var", func(t *testing.T) {
		var variable string = "this is a variable"
		want := "this is a variable"

		AssertCheckVariable(t, variable, want)

		var number = 10
		wantNumber := 10

		AssertCheckVariable(t, number, wantNumber)
	})

	t.Run("indexing string", func(t *testing.T) {
		var variable = "hello go!"
		indexVar := string(variable[2])
		want := "l"

		AssertCheckVariable(t, indexVar, want)
	})

	t.Run("check type data", func(t *testing.T) {
		integer := 10
		strings := "hello"
		want := []string{"int", "string"}

		AssertCheckDataType(t, integer, want[0])
		AssertCheckDataType(t, strings, want[1])
	})
}

func AssertCheckVariable(t testing.TB, variable, want interface{}) {
	if variable != want {
		t.Errorf("expected: %v but got -> %v", want, variable)
	}
}

func AssertCheckDataType(t testing.TB, data interface{}, want interface{}) {
	if reflect.TypeOf(data).String() != want {
		t.Errorf("expected: %q but got -> %q", want, data)
	}
}
