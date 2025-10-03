package product

type Product struct {
	Id    int
	Name  string
	Price float64
}

func New(id int, name string, price float64) Product {
	return Product{
		Id:    id,
		Name:  name,
		Price: price,
	}
}
