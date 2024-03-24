package basic

import (
	"strconv"
	"strings"
	"testing"
)

func TestStrings(t *testing.T) {

	t.Run("convert int to string", func(t *testing.T) {
		var strNumber string = "100"

		convert, err := strconv.Atoi(strNumber)
		if err != nil {
			panic(err)
		}
		expect := 100

		AssertCheckEqual(t, convert, expect)
	})

	t.Run("convert string to int", func(t *testing.T) {
		var number int = 100

		convert := strconv.Itoa(number)
		expect := "100"

		AssertCheckEqual(t, convert, expect)
	})

	t.Run("convert to lower & upper case", func(t *testing.T) {
		upperCase := "This is a text"
		upper := strings.ToUpper(upperCase)

		AssertCheckEqual(t, upper, "THIS IS A TEXT")

		lowerCase := "Passing a lower TEXT"
		lower := strings.ToLower(lowerCase)

		AssertCheckEqual(t, lower, "passing a lower text")
	})

	t.Run("cut a word", func(t *testing.T) {
		text := "12. hello world!"

		before, after, found := strings.Cut(text, "12. ")

		AssertCheckEqual(t, before, "")
		AssertCheckEqual(t, after, "hello world!")
		AssertCheckEqual(t, found, true)
	})

	t.Run("trim a string", func(t *testing.T) {
		text := "12. hello world!"

		trim := strings.Trim(text, "12.!")
		leftTrim := strings.TrimLeft(text, "12.")
		rightTrim := strings.TrimRight(text, "!")

		AssertCheckEqual(t, trim, " hello world")
		AssertCheckEqual(t, leftTrim, " hello world!")
		AssertCheckEqual(t, rightTrim, "12. hello world")
	})

	t.Run("replace string", func(t *testing.T) {
		text := "oink oink oink"

		replaceFirstTwoWords := strings.Replace(text, "o", "bo", 2)
		replaceWithNoLimit := strings.Replace(text, "o", "bo", -1)

		AssertCheckEqual(t, replaceFirstTwoWords, "boink boink oink")
		AssertCheckEqual(t, replaceWithNoLimit, "boink boink boink")

		replaceAllWords := strings.ReplaceAll(text, "oink", "hello")
		AssertCheckEqual(t, replaceAllWords, "hello hello hello")
	})

	t.Run("count a word in text", func(t *testing.T) {
		text := "oink oink oink"

		countO := strings.Count(text, "i")

		AssertCheckEqual(t, countO, 3)
	})

	t.Run("check contains word in text", func(t *testing.T) {
		text := "he's the greatest developer"

		containsWord := strings.Contains(text, "the")

		AssertCheckEqual(t, containsWord, true)
	})
}
