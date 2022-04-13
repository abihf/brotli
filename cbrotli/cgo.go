//go:build cbrotli && cgo

// Copyright 2017 Google Inc. All Rights Reserved.
//
// Distributed under MIT license.
// See https://github.com/google/brotli/blob/v1.0.9/LICENSE

package cbrotli

// Inform golang build system that it should link brotli libraries.

// #cgo LDFLAGS: -lbrotlicommon
// #cgo LDFLAGS: -lbrotlidec
// #cgo LDFLAGS: -lbrotlienc
import "C"
