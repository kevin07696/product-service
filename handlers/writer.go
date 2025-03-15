package handlers

import (
	"context"

	"github.com/kevin07696/produce-service/domain"
	"github.com/kevin07696/produce-service/generated"
)

type WriteProductHandler struct {
	generated.UnimplementedProductWriteServiceServer

	service      Service
	errorHandler []error
}

func NewWriteProductHandler(service *domain.ProductService) WriteProductHandler {
	return WriteProductHandler{
		service:      service,
		errorHandler: initErrorHandler(),
	}
}

func (h WriteProductHandler) WriteProduct(ctx context.Context, request *generated.WriteProductRequest) (*generated.WriteProductResponse, error) {
	id, status := h.service.WriteProduct(ctx, request.Product)

	response := &generated.WriteProductResponse{Id: id}
	return response, h.errorHandler[status]
}

func (h WriteProductHandler) UpdateProduct(ctx context.Context, request *generated.UpdateProductRequest) (*generated.UpdateProductResponse, error) {
	status := h.service.UpdateProduct(ctx, request)

	response := &generated.UpdateProductResponse{}
	if status == 0 {
		response.Success = true
		return response, h.errorHandler[status]
	}
	return response, h.errorHandler[status]
}

func (h WriteProductHandler) DeleteProduct(ctx context.Context, request *generated.DeleteProductRequest) (*generated.DeleteProductResponse, error) {
	status := h.service.DeleteProduct(ctx, request)

	response := &generated.DeleteProductResponse{}
	if status == 0 {
		response.Success = true
		return response, h.errorHandler[status]
	}
	return response, h.errorHandler[status]
}

func (h WriteProductHandler) RecoverProduct(ctx context.Context, request *generated.RecoverProductRequest) (*generated.RecoverProductResponse, error) {
	status := h.service.RecoverProduct(ctx, request)

	response := &generated.RecoverProductResponse{}
	if status == 0 {
		response.Success = true
		return response, h.errorHandler[status]
	}
	return response, h.errorHandler[status]
}
