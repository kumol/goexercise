package arrays

import (
	"fmt"

	"example.com/stucts/product"
)

func ArraySlice() {
	var products = make([]product.Product, 10, 10)
	products = append(products, product.New(1, "Laptop", 9999))
	i := 0
	for i < 10 {
		products[i] = product.New(i, `Product `+string(i), float64(i*100.00))
		// products = append(products, product.New(i, `Product `+string(i), float64(i*100.00)))
		i++
	}
	value := map[string]int{"kumol": 1, "sakura": 2}
	value["sakura"] = 3
	for key, value := range value {
		fmt.Println(key, value)
	}
	for _, product := range products {
		fmt.Println(product.Id)
	}
	fmt.Println(value)
	fmt.Println(products)
}
