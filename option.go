package brotli

const (
	BestSpeed          = 0
	BestCompression    = 11
	DefaultCompression = 6
)

// Options configures Writer.
type Options struct {
	// Quality controls the compression-speed vs compression-density trade-offs.
	// The higher the quality, the slower the compression. Range is 0 to 11.
	Quality int
	// LGWin is the base 2 logarithm of the sliding window size.
	// Range is 10 to 24. 0 indicates automatic configuration based on Quality.
	LGWin int
}
