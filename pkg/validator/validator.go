package validator

import (
	"context"

	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
)

var v = validator.New()

func RegisterCustomTypeFunc(fn validator.CustomTypeFunc, types ...interface{}) {
	v.RegisterCustomTypeFunc(fn, types...)
}

func RegisterAlias(alias, tags string) {
	v.RegisterAlias(alias, tags)
}

func RegisterStructValidation(fn validator.StructLevelFunc, types ...interface{}) {
	v.RegisterStructValidation(fn, types...)
}

func RegisterStructValidationCtx(fn validator.StructLevelFuncCtx, types ...interface{}) {
	v.RegisterStructValidationCtx(fn, types...)
}

func RegisterTranslation(tag string, trans ut.Translator, registerFn validator.RegisterTranslationsFunc, translationFn validator.TranslationFunc) error {
	return v.RegisterTranslation(tag, trans, registerFn, translationFn)
}

func RegisterTagNameFunc(fn validator.TagNameFunc) {
	v.RegisterTagNameFunc(fn)
}

func RegisterValidation(tag string, fn validator.Func, callValidationEvenIfNull ...bool) error {
	return v.RegisterValidation(tag, fn, callValidationEvenIfNull...)
}

func RegisterValidationCtx(tag string, fn validator.FuncCtx, callValidationEvenIfNull ...bool) error {
	return v.RegisterValidationCtx(tag, fn, callValidationEvenIfNull...)
}

func Struct(s interface{}) error {
	return v.Struct(s)
}

func StructCtx(ctx context.Context, s interface{}) error {
	return v.StructCtx(ctx, s)
}

func StructFiltered(s interface{}, fn validator.FilterFunc) error {
	return v.StructFiltered(s, fn)
}

func StructFilteredCtx(ctx context.Context, s interface{}, fn validator.FilterFunc) error {
	return v.StructFilteredCtx(ctx, s, fn)
}

func StructExcept(s interface{}, fields ...string) error {
	return v.StructExcept(s, fields...)
}

func StructExceptCtx(ctx context.Context, s interface{}, fields ...string) (err error) {
	return v.StructExceptCtx(ctx, s, fields...)
}

func StructPartial(s interface{}, fields ...string) error {
	return v.StructPartial(s, fields...)
}

func StructPartialCtx(ctx context.Context, s interface{}, fields ...string) (err error) {
	return v.StructPartialCtx(ctx, s, fields...)
}

func SetTagName(name string) {
	v.SetTagName(name)
}

func ValidateMap(data map[string]interface{}, rules map[string]interface{}) map[string]interface{} {
	return v.ValidateMap(data, rules)
}

func ValidateMapCtx(ctx context.Context, data map[string]interface{}, rules map[string]interface{}) map[string]interface{} {
	return v.ValidateMapCtx(ctx, data, rules)
}

func VarWithValue(field interface{}, other interface{}, tag string) error {
	return v.VarWithValue(field, other, tag)
}

func VarWithValueCtx(ctx context.Context, field interface{}, other interface{}, tag string) (err error) {
	return v.VarWithValueCtx(ctx, field, other, tag)
}

func Var(field interface{}, tag string) error {
	return v.Var(field, tag)
}

func VarCtx(ctx context.Context, field interface{}, tag string) error {
	return v.Var(field, tag)
}
