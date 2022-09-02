package validate

import (
	"context"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	errorx2 "github.com/zhaoqiang0201/zero-vue-admin/server/usercenter/api/pkg/errorx"
	"reflect"
)

func StructExceptCtx(ctx context.Context, v interface{}, field ...string) error {
	zh_cn := zh.New()
	vali := validator.New()
	//注册一个函数，获取struct tag里自定义的label作为字段名
	vali.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := field.Tag.Get("comment")
		return name
	})
	uni := ut.New(zh_cn)
	trans, _ := uni.GetTranslator("zh")
	zh_translations.RegisterDefaultTranslations(vali, trans)

	if err := vali.StructExceptCtx(ctx, v, field...); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errorx2.New(errorx2.REUQEST_PARAM_ERROR, err.Translate(trans))
		}
	}
	return nil
}
