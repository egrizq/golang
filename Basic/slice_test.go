package basic

import "testing"

func TestSlice(t *testing.T) {

	company := []string{"AMD", "Nvidia", "TSMC"}
	t.Run("change value", func(t *testing.T) {
		company[0] = "MSFT"

		AssertCheckEqual(t, company[0], "MSFT")
	})

	t.Run("append data", func(t *testing.T) {
		company = append(company, "GOOGL", "TSLA")

		AssertCheckEqual(t, company[3], "GOOGL")
	})
}
