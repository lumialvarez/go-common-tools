package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lumialvarez/go-common-tools/http/apierrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// HandlerFunc is the func type for the custom handlers.
type HandlerFunc func(c *gin.Context) *apierrors.APIError

// ErrorWrapper if handlerFunc return a error,then response will be composed from error's information.
func ErrorWrapper(handlerFunc HandlerFunc, c *gin.Context) {
	err := handlerFunc(c)
	if err != nil {
		c.JSON(err.Status, err)
	}
}

func ToAPIError(grpcError error) *apierrors.APIError {
	if e, ok := status.FromError(grpcError); ok {
		switch e.Code() {
		case codes.Canceled:
			return apierrors.NewInternalServerError("Internal Server Error", "Status gRPC Canceled", e.Message())
		case codes.Unknown:
			return apierrors.NewInternalServerError("Internal Server Error", "Status gRPC Unknown", e.Message())
		case codes.InvalidArgument:
			return apierrors.NewBadRequestError("Bad Request Error", "Status gRPC InvalidArgument", e.Message())
		case codes.DeadlineExceeded:
			return apierrors.NewInternalServerError("Internal Server Error", "Status gRPC DeadlineExceeded", e.Message())
		case codes.NotFound:
			return apierrors.NewNotFoundError("Not Found Error", "Status gRPC NotFound", e.Message())
		case codes.AlreadyExists:
			return apierrors.NewConflictError("Conflict Error", "Status gRPC AlreadyExists", e.Message())
		case codes.PermissionDenied:
			return apierrors.NewForbiddenError("Forbidden Error", "Status gRPC PermissionDenied", e.Message())
		case codes.ResourceExhausted:
			return apierrors.NewBadGatewayError("Bad Gateway", "Status gRPC ResourceExhausted", e.Message())
		case codes.FailedPrecondition:
			return apierrors.NewBadGatewayError("Bad Gateway", "Status gRPC FailedPrecondition", e.Message())
		case codes.Aborted:
			return apierrors.NewBadGatewayError("Bad Gateway", "Status gRPC Aborted", e.Message())
		case codes.OutOfRange:
			return apierrors.NewBadGatewayError("Bad Gateway", "Status gRPC OutOfRange", e.Message())
		case codes.Unimplemented:
			return apierrors.NewBadGatewayError("Bad Gateway", "Status gRPC Unimplemented", e.Message())
		case codes.Internal:
			return apierrors.NewBadGatewayError("Bad Gateway", "Status gRPC Internal", e.Message())
		case codes.Unavailable:
			return apierrors.NewBadGatewayError("Bad Gateway", "Status gRPC Unavailable", e.Message())
		case codes.DataLoss:
			return apierrors.NewBadGatewayError("Bad Gateway", "Status gRPC DataLoss", e.Message())
		case codes.Unauthenticated:
			return apierrors.NewUnauthorizedError("Unauthorized Error", "Status gRPC Unauthenticated", e.Message())

		default:
			return apierrors.NewBadGatewayError("Bad Gateway", "Unknown Authorization Service response", e.Message())
		}
	}
	return apierrors.NewBadGatewayError("Bad Gateway", "Unknown Authorization Service response", "Unknown gRPC status code", grpcError.Error())
}
