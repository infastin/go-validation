package validation_test

import (
	"testing"

	"github.com/infastin/go-validation"
)

func Test_RequiredSlice_Validate(t *testing.T) {
	tests := []struct {
		name  string
		slice []string
		want  error
	}{
		{
			name:  "nil",
			slice: nil,
			want:  validation.ErrRequired,
		},
		{
			name:  "empty",
			slice: []string{},
			want:  validation.ErrRequired,
		},
		{
			name:  "not empty",
			slice: make([]string, 10),
			want:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := validation.RequiredSlice[string](true)
			if got := r.Validate(tt.slice); got != tt.want {
				t.Errorf("RequiredSlice.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NilOrNotEmptySlice_Validate(t *testing.T) {
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
			want:  validation.ErrNilOrNotEmpty,
		},
		{
			name:  "not empty",
			slice: make([]string, 10),
			want:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := validation.NilOrNotEmptySlice[string](true)
			if got := r.Validate(tt.slice); got != tt.want {
				t.Errorf("NilOrNotEmptySlice.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_RequiredMap_Validate(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]string
		want error
	}{
		{
			name: "nil",
			m:    nil,
			want: validation.ErrRequired,
		},
		{
			name: "empty",
			m:    make(map[string]string),
			want: validation.ErrRequired,
		},
		{
			name: "not empty",
			m:    map[string]string{"foo": "bar"},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := validation.RequiredMap[string](true)
			if got := r.Validate(tt.m); got != tt.want {
				t.Errorf("RequiredMap.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NilOrNotEmptyMap_Validate(t *testing.T) {
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
			want: validation.ErrNilOrNotEmpty,
		},
		{
			name: "not empty",
			m:    map[string]string{"foo": "bar"},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := validation.NilOrNotEmptyMap[string](true)
			if got := r.Validate(tt.m); got != tt.want {
				t.Errorf("NilOrNotEmptyMap.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
