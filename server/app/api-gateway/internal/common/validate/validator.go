package validate

import (
	"context"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/pkg/errors"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/common/responseerror/errorx"

	"github.com/go-playground/validator/v10"
	"reflect"
)

func StructExceptCtx(ctx context.Context, v interface{}, field ...string) error {
	validate := validator.New()
	//注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := field.Tag.Get("commen")
		return name
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
