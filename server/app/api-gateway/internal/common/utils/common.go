package utils

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"strconv"
)

func GetUserIdWithJWT(ctx context.Context) (userid uint64, err error) {
	userIDstr := ctx.Value("userID").(json.Number).String()
	userid, err = strconv.ParseUint(userIDstr, 10, 64)
	if err != nil {
		return 0, errors.Wrapf(err, "userID: %s", userIDstr)
	}
	return
}
