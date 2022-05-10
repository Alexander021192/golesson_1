package product

var allProducts = []Product{
	{Title: "one"},
	{Title: "two"},
	{Title: "three"},
	{Title: "one1"},
	{Title: "two2"},
	{Title: "three3"},
}

type Product struct {
	Title string
}