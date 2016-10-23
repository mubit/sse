// For go 1.5 and below bufio.Scanner.Buffer() did not exist
//+build !go1.6

package sse

import (
	"bufio"
	"bytes"
	"io"
)

// NewDecoder builds an SSE decoder with a growing buffer.
// Lines are limited to bufio.MaxScanTokenSize - 1.
func NewDecoder(in io.Reader) Decoder {
	d := &decoder{bufio.NewScanner(in), new(bytes.Buffer)}
	d.scanner.Split(scanLinesCR) // See scanlines.go
	return d
}