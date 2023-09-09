package status

import (
	"context"
	"google.golang.org/genproto/googleapis/rpc/status"

	"beyond/pkg/xcode"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
)

func FromError(err error) *grpcstatus.Status {
	err = errors.Cause(err)
	if code, ok := err.(xcode.XCode); ok {
		grpcStatus, e := gRPCStatusFromXCode(code)
		if e == nil {
			return grpcStatus
		}
	}

	var grpcStatus *grpcstatus.Status
	switch err {
	case context.Canceled:
		grpcStatus, _ = gRPCStatusFromXCode(xcode.Canceled)
	case context.DeadlineExceeded:
		grpcStatus, _ = gRPCStatusFromXCode(xcode.Deadline)
	default:
		grpcStatus, _ = grpcstatus.FromError(err)
	}

	return grpcStatus
}

func gRPCStatusFromXCode(code xcode.XCode) (*grpcstatus.Status, error) {
	switch v := code.(type) {
	case xcode.Code:
	default:

	}

	stas := grpcstatus.New(codes.Unknown)
}

func GrpcStatusToXCode(gstatus *status.Status) xcode.XCode {
	details := gstatus.Details()

}
