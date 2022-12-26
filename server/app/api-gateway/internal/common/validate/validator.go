package validate

import (
	"context"
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/pkg/errors"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func StructExceptCtx(ctx context.Context, v interface{}, field ...string) error {
	t1 := time.Now()
	defer func() {
		t2 := time.Now()
		fmt.Println("参数校验时间(s): ", t2.Sub(t1).Seconds())
	}()
	validate := validator.New()
	//注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := field.Tag.Get("commen")
		return name
	})

	validate.RegisterValidationCtx("slice_c", func(ctx context.Context, fl validator.FieldLevel) bool {
		if fl.Field().Type().Kind() == reflect.Slice {
			param := fl.Param()
			paramSlics := strings.Split(param, " ")
			f := []bool{}
			for _, v := range paramSlics {
				vSlice := strings.Split(strings.Trim(strings.TrimSpace(v), "'"), "=")
				if len(vSlice) != 2 {
					return false
				}
				vv, err := strconv.Atoi(vSlice[1])
				if err != nil {
					return false
				}
				switch vSlice[0] {
				case "gt":
					f = append(f, fl.Field().Len() > vv)
				case "lt":
					f = append(f, fl.Field().Len() < vv)
				case "gte":
					f = append(f, fl.Field().Len() >= vv)
				case "lte":
					f = append(f, fl.Field().Len() <= vv)
				case "eq":
					f = append(f, fl.Field().Len() == vv)
				}
				//fmt.Println("===>", vSlice[0], vv, fl.Field().Len(), f)
			}
			for _, v := range f {
				if !v {
					return false
				}
			}
		}
		//fl.Field().Len()
		//b := fl.GetTag() //lengthc
		//b := fl.FieldName() //Parameters
		//b := fl.Param() //'aa' 'bb' 'cc'
		//fmt.Println("fl.Field().String()", a)
		//fmt.Printf("fl.Param(),%q\n", b)
		return true
	})

	zh_cn := zh.New()
	uni := ut.New(zh_cn)
	trans, ok := uni.GetTranslator("zh")
	if ok {
		return errorx.NewByCode(errors.New("validate未找到指定地区的翻译程序"), errorx.SERVER_COMMON_ERROR)
	}

	if err := zh_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		return errorx.NewByCode(err, errorx.Validate_RegisterDefaultTranslations_ERROR)
	}

	if err := validate.StructExceptCtx(ctx, v, field...); err != nil {
		e := errorx.NewByCode(err, errorx.REUQEST_PARAM_ERROR)
		for _, err := range err.(validator.ValidationErrors) {
			e.WithMeta("", "", err.Translate(trans))
		}
		return e
	}
	return nil
}
