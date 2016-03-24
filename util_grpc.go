package errors

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// GRPCError allows an error to be converted to a GRPC error
type GRPCError interface {
	error
	GRPCCode() codes.Code
	GRPCMessage() string
}

// AsGRPCError converts any error into a GRPC error.
func AsGRPCError(err error) error {
	if x, ok := err.(GRPCError); ok && x != nil {
		return grpc.Errorf(x.GRPCCode(), x.GRPCMessage())
	}

	// Will be turned into a codes.Unknown unless it is a GRPC error
	return err
}
