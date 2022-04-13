package brotli

import (
	"context"
	"io"
)

func NewFactory(option *Options) func(context.Context, io.Writer) (io.WriteCloser, error) {
	return func(ctx context.Context, w io.Writer) (io.WriteCloser, error) {
		return NewWriter(w, option), nil
	}
}
