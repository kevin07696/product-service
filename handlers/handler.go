package handlers

import (
	"context"

	"github.com/kevin07696/produce-service/generated"
)

type ProductHandler struct {
	generated.UnimplementedProductServer

	servicer     ProductServicer
	reader       Reader
	errorHandler []error
}

func NewProductHandler(servicer ProductServicer, reader Reader) ProductHandler {
	return ProductHandler{
		servicer:     servicer,
		reader:       reader,
		errorHandler: initErrorHandler(),
	}
}

func (h ProductHandler) FilterProducts(ctx context.Context, request *generated.ProductSummaryRequest) (response *generated.ProductSummaryResponse, err error) {
	summaries, status := h.reader.ReadProductSummaries(request)
	if status > 0 {
		err = h.errorHandler[status]
		return
	}

	response = &generated.ProductSummaryResponse{
		Products: []*generated.ProductSummary{},
	}

	for _, summary := range summaries {
		response.Products = append(response.Products, &generated.ProductSummary{
			ID:           uint32(summary.ID),
			Name:         summary.Name,
			Price:        summary.Price,
			ThumbnailURL: summary.ThumbnailURL,
		})
	}

	return
}

func (h ProductHandler) GetDetails(ctx context.Context, request *generated.ProductDetailRequest) (response *generated.ProductDetailResponse, err error) {
	productDetail, status := h.servicer.GetProductDetail(request)
	if status > 0 {
		err = h.errorHandler[status]
		return
	}

	response = &generated.ProductDetailResponse{
		ID:          uint32(productDetail.ID),
		Category:    productDetail.Category,
		Price:       productDetail.Price,
		Name:        productDetail.Name,
		ImagesURLs:  productDetail.ImagesURLs,
		Description: productDetail.Description,
	}

	return
}

func (h ProductHandler) GetCategories(ctx context.Context, request *generated.CategoryRequest) (response *generated.CategoryResponse, err error) {
	categories, status := h.reader.ReadCategories()
	if status > 0 {
		err = h.errorHandler[status]
		return
	}

	response = &generated.CategoryResponse{Categories: categories}

	return
}
