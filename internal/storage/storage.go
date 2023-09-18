package storage

import (
	"context"
	"errors"
)

var (
	EmptyArguments = errors.New("arguments should not be nil or empty")
)

type IStorage interface {
	Call(ctx context.Context, req *[]byte, resp *[]byte) error
}
