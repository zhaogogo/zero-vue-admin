package logiccommon

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

func GetUserIdWithJWT(ctx context.Context) (userid uint64, err error) {
	userIDstr := ctx.Value("userID").(json.Number).String()
	userid, err = strconv.ParseUint(userIDstr, 10, 64)
	if err != nil {
		logx.Errorf("获取用户ID失败, error: %v", err)
	}
	return
}
