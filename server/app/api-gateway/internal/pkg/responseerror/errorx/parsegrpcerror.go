package errorx

import "google.golang.org/grpc/status"

func ParseGrpcError(err error) error {
	grpcerr, isRpcErr := status.FromError(err)
	if isRpcErr { //rpc错误
		return &Errorx{Code: ErrorCode(grpcerr.Proto().Code), Msg: "grpc错误", Cause: err}
	} else {
		return err
	}
}
