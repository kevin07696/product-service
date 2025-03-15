package handlers

import (
	"github.com/kevin07696/produce-service/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func initErrorHandler() []error {
	handler := make([]error, 6)

	handler[domain.StatusOK] = status.Error(codes.OK, "Success!")
	handler[domain.StatusInvalidArgument] = status.Error(codes.InvalidArgument, "The login credentials are invalid.")
	handler[domain.StatusUnauthorized] = status.Error(codes.Unauthenticated, "The login credentials are invalid.")
	handler[domain.StatusAlreadyExists] = status.Error(codes.InvalidArgument, "The login credentials are taken. Please try again.")
	handler[domain.StatusNotFound] = status.Error(codes.NotFound, "")
	handler[domain.StatusInternalError] = status.Error(codes.Internal, "Something went wrong. Please try again.")

	return handler
}
