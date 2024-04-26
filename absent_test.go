package validation_test

import (
	"testing"

	"github.com/infastin/go-validation"
)

func Test_NilSlice_Validation(t *testing.T) {
	tests := []struct {
		name  string
		slice []string
		want  error
	}{
		{
			name:  "nil",
			slice: nil,
			want:  nil,
		},
		{
			name:  "empty",
			slice: []string{},
			want:  validation.ErrNil,
		},
		{
			name:  "not empty",
			slice: make([]string, 10),
			want:  validation.ErrNil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := validation.NilSlice[string](true)
			if got := r.Validate(tt.slice); got != tt.want {
				t.Errorf("NilSlice.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_EmptySlice_Validation(t *testing.T) {
	tests := []struct {
		name  string
		slice []string
		want  error
	}{
		{
			name:  "nil",
			slice: nil,
			want:  nil,
		},
		{
			name:  "empty",
			slice: []string{},
			want:  nil,
		},
		{
			name:  "not empty",
			slice: make([]string, 10),
			want:  validation.ErrEmpty,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := validation.EmptySlice[string](true)
			if got := r.Validate(tt.slice); got != tt.want {
				t.Errorf("EmptySlice.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NilMap_Validation(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]string
		want error
	}{
		{
			name: "nil",
			m:    nil,
			want: nil,
		},
		{
			name: "empty",
			m:    make(map[string]string),
			want: validation.ErrNil,
		},
		{
			name: "not empty",
			m:    map[string]string{"foo": "bar"},
			want: validation.ErrNil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := validation.NilMap[string](true)
			if got := r.Validate(tt.m); got != tt.want {
				t.Errorf("NilMap.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_EmptyMap_Validation(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]string
		want error
	}{
		{
			name: "nil",
			m:    nil,
			want: nil,
		},
		{
			name: "empty",
			m:    make(map[string]string),
			want: nil,
		},
		{
			name: "not empty",
			m:    map[string]string{"foo": "bar"},
			want: validation.ErrEmpty,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := validation.EmptyMap[string](true)
			if got := r.Validate(tt.m); got != tt.want {
				t.Errorf("EmptyMap.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
