package rot13

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestNewWriter(t *testing.T) {
	tests := []struct {
		name  string
		want  *Writer
		wantW string
	}{
		{
			"nil writer", &Writer{Writer: &bytes.Buffer{}}, "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if got := NewWriter(w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWriter() = %#+v, want %#+v", got, tt.want)
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("NewWriter() = %#+v, want %#+v", gotW, tt.wantW)
			}
		})
	}
}

func TestWriter_Write(t *testing.T) {
	type fields struct {
		Writer io.Writer
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
		{
			"nil", fields{Writer: nil}, args{[]byte("test")}, 0, true,
		},
		{
			"test", fields{Writer: &bytes.Buffer{}}, args{[]byte("test")}, 4, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				Writer: tt.fields.Writer,
			}
			got, err := w.Write(tt.args.buf)
			if (err != nil) != tt.wantErr {
				t.Errorf("Writer.Write() error = %#+v, wantErr %#+v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Writer.Write() = %#+v, want %#+v", got, tt.want)
			}
		})
	}
}

func TestRot13Slice(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"Gur dhvpx oebja sbk whzcf bire gur ynml qbt", args{[]byte("Gur dhvpx oebja sbk whzcf bire gur ynml qbt")}, []byte("The quick brown fox jumps over the lazy dog")},
		{"The quick brown fox jumps over the lazy dog", args{[]byte("The quick brown fox jumps over the lazy dog")}, []byte("Gur dhvpx oebja sbk whzcf bire gur ynml qbt")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rot13Slice(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rot13Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}
