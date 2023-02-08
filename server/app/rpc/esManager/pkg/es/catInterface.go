package es

import "context"

type CatInterface interface {
	Health(ctx context.Context, param CatParam) (res interface{}, err error)
}
