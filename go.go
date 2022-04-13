//go:build !cbrotli

package brotli

import (
	"io"

	gobrotli "github.com/andybalholm/brotli"
)

func NewWriter(w io.Writer, option *Options) io.WriteCloser {
	return gobrotli.NewWriterOptions(w, gobrotli.WriterOptions{
		LGWin:   option.LGWin,
		Quality: option.Quality,
	})
}

func NewReader(r io.Reader) io.ReadCloser {
	return io.NopCloser(gobrotli.NewReader(r))
}
