package responseerror

import (
	"database/sql"
	"github.com/pkg/errors"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"

	"google.golang.org/grpc/status"
)

type ErrorResponse struct {
	Code     int32  `json:"code"`
	Msg      string `json:"msg"`
	CauseErr string `json:"cause_err"`
}

/*
	code: 1-17   ==>   rpc错误
	code: 9999   ==>   其他错误
*/
func ErrorHandle(err error) (int, interface{}) {
	causeErr := errors.Cause(err)
	switch e := causeErr.(type) {
	case *errorx.Errorx: //自定义错误
		if e.Code == errorx.UNAUTHORIZATION {
			return 401, ErrorResponse{Code: int32(e.Code), Msg: e.Msg, CauseErr: e.Cause.Error()}
		}
		return 200, ErrorResponse{Code: int32(e.Code), Msg: e.Msg, CauseErr: e.Cause.Error()}
	default:
		//fmt.Printf("=====>  default  %T %v\n", err, err)
		//=====>  default  *status.Error      rpc error: code = Unavailable desc = last resolver error: produced zero addresses
		//s, o := status.FromError(e)
		//fmt.Println(o)           // true -> rpc错误， false 其他错误
		//fmt.Println(s.Err())     // rpc error: code = Unavailable desc = last resolver error: produced zero addresses
		//fmt.Println(s.Message()) //last resolver error: produced zero addresses
		//fmt.Println(s.Code())    //Unavailable
		//fmt.Println(s.Proto())   //code:14  message:"last resolver error: produced zero addresses"
		//fmt.Println(s.Details()) //[]
		//fmt.Println(s.String())  //rpc error: code = Unavailable desc = last resolver error: produced zero addresses
		e2, isRpcErr := status.FromError(e)
		if isRpcErr { // rpc错误
			if e2.Message() == sql.ErrNoRows.Error() { //rpc查询为空
				return 200, ErrorResponse{Code: e2.Proto().Code, Msg: "grpc错误", CauseErr: e2.String()}
			}
			return 200, ErrorResponse{Code: e2.Proto().Code, Msg: "grpc错误", CauseErr: e2.String()}
		}
		return 200, ErrorResponse{Code: 9999, Msg: "XXX", CauseErr: e.Error()}
	}
}
