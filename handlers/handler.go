package handlers

import (
	"context"
	"fmt"

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
			Price:        fmt.Sprintf("%.2f", summary.Price),
			ThumbnailURL: summary.ThumbnailURL,
		})
	}

	return
}
