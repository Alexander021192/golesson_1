package product

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List(page int) []Product {
	start := 0
	end := len(allProducts)
	offset := (page - 1) * 5
	if end == 0 || offset > end {
		currrentPage = 0
		return []Product{{Title: "Finish list"}}
	}

	if offset < end {
		start = offset
	}
	if offset + 5 < end {
		end = offset + 5
	}
	return allProducts[start:end]
}

func (s * Service) Get(idx int) (*Product, error) {
	// need validation

	return &allProducts[idx], nil
}

func (s * Service) Len() (int, error) {
	return len(allProducts), nil
}

func (s * Service) CurrentPage() (*int, error) {
	return &currrentPage, nil
}