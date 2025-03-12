package domain

type ProductService struct {
	repo ProductWriter
}

func NewProductService(repo ProductWriter) ProductService {
	return ProductService{
		repo: repo,
	}
}