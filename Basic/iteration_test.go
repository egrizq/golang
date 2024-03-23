package basic

import (
	"strings"
	"testing"
)

func sumMaxNumber(maxNumber int) int {
	counter := 0
	for current := 0; current < maxNumber; current++ {
		counter += current
	}
	return counter
}

func sumArray(array [3]int) int {
	counter := 0
	for index := range array {
		counter += array[index]
	}
	return counter
}

func sumSlice(array []int) int {
	counter := 0
	for _, value := range array {
		counter += value
	}
	return counter
}

// In Go there are no while, do, until keywords, you can only use for.
func TestIteration(t *testing.T) {

	t.Run("basic looping", func(t *testing.T) {
		maxNumber := 10
		counter := sumMaxNumber(maxNumber)
		want := 45

		AssertCheckEqual(t, counter, want)
	})

	t.Run("looping with array", func(t *testing.T) {
		array := [3]int{10, 12, 10}
		counter := sumArray(array)
		want := 32

		AssertCheckEqual(t, counter, want)
	})

	t.Run("looping with slice", func(t *testing.T) {
		slice := []int{10, 12, 10}
		counter := sumSlice(slice)
		want := 32

		AssertCheckEqual(t, counter, want)
	})

	t.Run("looping with condition", func(t *testing.T) {
		word := "leetCode"
		cutWord := []string{""}

		for i := range word {
			if string(word[i]) != "e" {
				cutWord = append(cutWord, string(word[i]))
			}
		}

		expect := strings.Join(cutWord, "")
		want := "ltCod"
		AssertCheckEqual(t, expect, want)
	})
}

func AssertCheckEqual(t testing.TB, expect, want interface{}) {
	if expect != want {
		t.Errorf("expected: %v but got -> %v", expect, want)
	}
}
