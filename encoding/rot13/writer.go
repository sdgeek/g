package rot13

import (
	"errors"
	"fmt"
	"io"

	"github.com/sdgeek/g/gen"
)

type Writer struct {
	io.Writer
}

var ErrNilWriter = errors.New("nil writer")

// NewWriter constructs a new rot13 Writer from an io.Writer.
func NewWriter(w io.Writer) *Writer {
	return &Writer{w}
}

// Rot13Slice performs a rot13 cipher on a char slice
func Rot13Slice[inp Char](b []inp) []inp {
	// ret := make([]inp, len(b))
	return gen.Map(b, func(c inp) inp {
		return Rot13(c)
	})
}

// Write implements io.Writer for a rot13 cipher.
func (w *Writer) Write(buf []byte) (int, error) {
	if w.Writer != nil {
		rot := Rot13Slice(buf)
		// use embedded io.Writer to write buf to file
		if n, err := w.Writer.Write(rot); err == nil {
			return n, err
		} else {
			return n, fmt.Errorf("write failed: %w", err)
		}
	}
	return 0, ErrNilWriter
}
