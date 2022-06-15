package rot13

import (
	"errors"
	"fmt"
	"io"

	"github.com/sdgeek/g/gen"
)

var ErrNillReader = errors.New("nil reader")

type Reader struct {
	io.Reader
}

// NewReader constructs a new rot13 Reader from an io.Reader.
func NewReader(r io.Reader) *Reader {
	return &Reader{r}
}

// Read implements io.Reader for a rot13 cipher.
func (r *Reader) Read(buf []byte) (int, error) {
	if r.Reader != nil {
		// use embedded io.Reader to read file into buf
		if n, err := r.Reader.Read(buf); err == nil {
			// rot13 the buffer in-place
			gen.ForEach(buf[:n], func(b *byte) {
				*b = Rot13(*b)
			})
			return n, err
		} else {
			return n, fmt.Errorf("read failed: %w", err)
		}
	}
	return 0, ErrNillReader
}
