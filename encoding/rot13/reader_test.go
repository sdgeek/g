package rot13

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestNewReader(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want *Reader
	}{
		{"nil reader", args{nil}, &Reader{nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReader(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReader() = %#+v, want %#+v", got, tt.want)
			}
		})
	}
}

func TestReader_Read(t *testing.T) {
	type fields struct {
		Reader io.Reader
	}
	type args struct {
		buf []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{"nil reader", fields{nil}, args{[]byte("test")}, 0, true},
		{"test", fields{bytes.NewReader([]byte("test"))}, args{make([]byte, 10)}, 4, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Reader{
				Reader: tt.fields.Reader,
			}
			got, err := r.Read(tt.args.buf)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.Read() error = %#+v, wantErr %#+v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Reader.Read() = %#+v, want %#+v", got, tt.want)
			}
			// if !bytes.Compare(tt.args.buf, []byte("test")) {
			// 	t.Errorf("Reader.Read() = %#+v, want %#+v", tt.args.buf, []byte("test"))
			// }
		})
	}
}
