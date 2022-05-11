package product

import (
	"bufio"
	"log"
	"os"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

// func (s *Service) List(page int) []Product {
// 	start := 0
// 	end := len(allProducts)
// 	offset := (page - 1) * 5
// 	if end == 0 || offset > end {
// 		currrentPage = 0
// 		return []Product{{Title: "Finish list"}}
// 	}

// 	if offset < end {
// 		start = offset
// 	}
// 	if offset+5 < end {
// 		end = offset + 5
// 	}
// 	return allProducts[start:end]
// }

// func (s *Service) Get(idx int) (*Product, error) {
// 	// need validation
// 	return &allProducts[idx], nil
// }

func (s *Service) Count() (int, error) {
	file, err := os.Open("products.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return count, nil
}

func (s *Service) CurrentPage() (*int) {
	return &currrentPage
}

func (s *Service) ReadProducts(start, end int) []Product {
	file, err := os.Open("products.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var products []Product
	for i := 0; i < end; i++ {
		scanner.Scan()
		// fmt.Println(scanner.Text())
		if i >= start && i <= end {
			products = append(products, Product{Title: scanner.Text()})
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return products
}
