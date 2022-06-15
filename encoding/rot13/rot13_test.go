package rot13

import (
	"reflect"
	"testing"
)

func TestRot13Byte(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{"a", args{'a'}, 'n'},
		{"N", args{'N'}, 'A'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rot13(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rot13() = %#+v, want %#+v", got, tt.want)
			}
		})
	}
}

func TestRot13Rune(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{"a", args{'a'}, 'n'},
		{"N", args{'N'}, 'A'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rot13(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rot13() = %#+v, want %#+v", got, tt.want)
			}
		})
	}
}

func TestRot13Int(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"a", args{'a'}, 'n'},
		{"n", args{'N'}, 'A'}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rot13(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rot13() = %v, want %v", got, tt.want)
			}
		})
	}
}
