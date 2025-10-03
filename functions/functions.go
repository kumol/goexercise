package functions

import "fmt"

func Mani() {
	value := []int{1, 2, 3, 4, 5}

	doubledValue := doubleValues(&value)

	fmt.Println(doubledValue)
}
func doubleValues(vals *[]int) []int {
	var doubled []int
	for _, v := range *vals {
		doubled = append(doubled, v*2)
	}
	return doubled
}
