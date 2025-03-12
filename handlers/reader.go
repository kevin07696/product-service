package handlers

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/kevin07696/produce-service/domain"
	"github.com/kevin07696/produce-service/generated"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/datatypes"
)

type ReadProductHandler struct {
	generated.UnimplementedProductReadServiceServer

	reader       Reader
	errorHandler []error
}

func NewReadProductHandler(reader Reader) ReadProductHandler {
	return ReadProductHandler{
		reader:       reader,
		errorHandler: initErrorHandler(),
	}
}

func (h ReadProductHandler) ListProducts(req *generated.ListSummariesRequest, stream grpc.ServerStream) error {
	// Read product summaries from the data source
	summaries, status := h.reader.ReadProductSummaries()
	if status > 0 {
		return h.errorHandler[status]
	}

	// Iterate over the summaries and send each summary as a response
	for _, summary := range summaries {
		// Create a response object for each product summary
		response := &generated.ListSummariesResponse{
			Products: []*generated.Summary{
				{
					Id:           summary.ID,
					Name:         summary.Name,
					ThumbnailUrl: summary.ThumbnailUrl,
					CategoryName: summary.CategoryName,
					CentPrice:    summary.CentPrice,
					InStock:      summary.InStock,
					Rating:       summary.Rating,
					AmountSold:   summary.AmountSold,
					CreatedAt:    timestamppb.New(summary.CreatedAt),
				},
			},
		}

		// Send the response object for the current summary
		if err := stream.SendMsg(response); err != nil {
			return err
		}
	}

	return nil
}


func (h ReadProductHandler) GetDetails(ctx context.Context, request *generated.GetDetailRequest) (*generated.Detail, error) {
	parsedUUID, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, h.errorHandler[domain.StatusUnauthorized]
	}

	dbId := datatypes.UUID(parsedUUID)

	productDetail, status := h.reader.ReadProductDetail(dbId)
	if status > 0 {
		return nil, h.errorHandler[status]
	}

	var attributes *structpb.Struct
	if err := protojson.Unmarshal(productDetail.Attributes, attributes); err != nil {
		log.Fatalf("Failed to decode attributes: %v", err)
	}

	var mainOption *generated.MainOption
	if err := protojson.Unmarshal(productDetail.MainOption, mainOption); err != nil {
		return nil, err
	}

	var optionList *generated.OptionList = &generated.OptionList{
		Options: []*generated.Option{},
	}
	if err := protojson.Unmarshal([]byte(productDetail.Options), optionList); err != nil {
		return nil, err
	}

	response := &generated.Detail{
		Id:           productDetail.ID,
		Name:         productDetail.Name,
		CategoryName: productDetail.CategoryName,
		Description:  productDetail.Description,
		CentPrice:    productDetail.CentPrice,
		Rating:       productDetail.Rating,
		Attributes:   attributes,
		MainOption:   mainOption,
		Options:      optionList.Options,
		CreatedAt:    timestamppb.New(productDetail.CreatedAt),
	}

	return response, nil
}

func (h ReadProductHandler) GetCategories(ctx context.Context, request *generated.ListCategoriesRequest) (*generated.ListCategoriesResponse, error) {
	categories, status := h.reader.ReadCategories()
	if status > 0 {
		return nil, h.errorHandler[status]
	}

	return &generated.ListCategoriesResponse{Categories: categories}, nil
}
