package basic

import (
	"fmt"
	"testing"
)

type Cat struct {
	Name  string
	Color string
	Age   int64
}

func myCat(profile Cat) string {
	return fmt.Sprintf(`my cat name is %v and have %v color, and %v years old`, profile.Name, profile.Color, profile.Age)
}

type EmployeeSalary struct {
	Base  int64
	Bonus int64
}

func (salary EmployeeSalary) MonthlySalary() float64 {
	return float64(salary.Base) + float64(salary.Bonus)
}

type Box struct {
	Left  int
	Right int
}

func (box Box) CalcBox() int {
	return box.Left * box.Right
}

func TestStruct(t *testing.T) {

	t.Run("func with struct", func(t *testing.T) {
		snowbell := Cat{
			Name:  "snowbell",
			Color: "white",
			Age:   12,
		}
		expect := myCat(snowbell)
		got := "my cat name is snowbell and have white color, and 12 years old"

		AssertCheckEqual(t, expect, got)
	})

	t.Run("method struct", func(t *testing.T) {
		jamesSalary := EmployeeSalary{1000, 1000}

		expect := jamesSalary.MonthlySalary()
		var got float64 = 2000

		AssertCheckEqual(t, expect, got)
	})

	t.Run("struct with slice", func(t *testing.T) {
		sdnBekasi := []Box{{5, 5}, {2, 4}}
		result := 0

		for index := range sdnBekasi {
			result += sdnBekasi[index].CalcBox()
		}
		got := 33

		AssertCheckEqual(t, result, got)
	})
}
