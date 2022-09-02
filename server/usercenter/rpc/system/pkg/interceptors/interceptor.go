package interceptors

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/pkg/validate"
	"google.golang.org/grpc"
)

func GrpcLogAndValidate(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	//参数校验
	if err := validate.StructExceptCtx(ctx, req); err != nil {
		return nil, err
	}
	resp, err = handler(ctx, req)
	if err != nil {
		logx.WithContext(ctx).Error(err)
		return resp, err
	} else {
		logx.WithContext(ctx).Infof("登陆成功, 参数: %+v, 响应: %+v", req, resp)
		return resp, err
	}

}
