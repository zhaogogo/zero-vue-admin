package file

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/s3connManager"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/model/monitoring"
	"strings"
	"time"

	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/svc"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.FileListRequest) (resp *types.FileListResponse, err error) {
	conn := &monitoring.StoreConnectManager{Host: req.Host}
	if err = l.svcCtx.MonitoringDB().Take(conn).Error; err != nil {
		return nil, errorx.New(err, "获取连接失败")
	}
	s3Client, err := s3connManager.NewConnManager().GetS3Client(conn.Host, conn.AccessKey, conn.SecretKey)
	if err != nil {
		return nil, errorx.New(err, "创建客户端失败")
	}

	switch {
	case req.Path == "":
		ctx, _ := context.WithTimeout(l.ctx, time.Second*1)
		buckets, err := s3Client.ListBucketsWithContext(ctx, &s3.ListBucketsInput{})
		if err != nil {
			return nil, errorx.New(err, "列出bucket失败")
		}
		lists := []types.Files{}
		for _, bucket := range buckets.Buckets {
			list := types.Files{
				Name:         *bucket.Name + "/",
				LastModified: bucket.CreationDate.Local().UnixMilli(),
				Size:         0,
				IsFile:       false,
			}
			lists = append(lists, list)
		}
		return &types.FileListResponse{
			HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK"},
			List:               lists,
		}, nil
	case strings.HasPrefix(req.Path, "s3://"):
		fullpath := req.Path[5:]
		bucketIndex := strings.Index(fullpath, "/")
		if bucketIndex == -1 {
			return nil, errorx.New(errors.New("查找bucket失败"), "查找bucket失败")
		}
		ctx, _ := context.WithTimeout(l.ctx, time.Second*2)
		listObjectsOutput, err := s3Client.ListObjectsWithContext(ctx, &s3.ListObjectsInput{
			Bucket:    aws.String(fullpath[:bucketIndex]),
			Delimiter: aws.String("/"),
			Prefix:    aws.String(fullpath[bucketIndex+1:]),
		})
		if err != nil {
			return nil, errorx.New(err, fmt.Sprintf("列出文件失败, host: %q, bucket: %q, prefix: %q", req.Host, fullpath[:bucketIndex], fullpath[bucketIndex+1:]))
		}
		lists := []types.Files{}
		for _, c := range listObjectsOutput.CommonPrefixes {
			f := types.Files{
				Name: strings.Trim(*c.Prefix, *listObjectsOutput.Prefix),
			}
			lists = append(lists, f)
		}
		for _, c := range listObjectsOutput.Contents {
			f := types.Files{
				Name:         strings.Trim(*c.Key, *listObjectsOutput.Prefix),
				LastModified: c.LastModified.Local().Unix(),
				Size:         *c.Size,
				IsFile:       true,
			}
			lists = append(lists, f)
		}
		return &types.FileListResponse{
			HttpCommonResponse: types.HttpCommonResponse{Code: 200, Msg: "OK"},
			List:               lists,
		}, nil
	}
	return nil, errorx.New(errors.New("其他错误"), "其他错误")
}
