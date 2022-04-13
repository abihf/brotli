//go:build cbrotli && cgo

package brotli

import (
	"io"

	"github.com/abihf/brotli/cbrotli"
)

func NewWriter(w io.Writer, option *Options) io.WriteCloser {
	return cbrotli.NewWriter(w, cbrotli.WriterOptions{Quality: option.Quality, LGWin: option.LGWin})
}

func NewReader(r io.Reader) io.ReadCloser {
	return cbrotli.NewReader(r)
}
