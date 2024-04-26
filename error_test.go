package validation_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/infastin/go-validation"
)

func Test_ruleError_Error(t *testing.T) {
	type fields struct {
		code    string
		message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ab",
			fields: fields{
				code:    "a",
				message: "b",
			},
			want: "b",
		},
		{
			name: "hello",
			fields: fields{
				code:    "hello",
				message: "world",
			},
			want: "world",
		},
		{
			name: "foobar",
			fields: fields{
				code:    "foo",
				message: "bar",
			},
			want: "bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := validation.NewRuleError(tt.fields.code, tt.fields.message)
			if got := re.Error(); got != tt.want {
				t.Errorf("ruleError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ruleError_Code(t *testing.T) {
	type fields struct {
		code    string
		message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ab",
			fields: fields{
				code:    "a",
				message: "b",
			},
			want: "a",
		},
		{
			name: "hello",
			fields: fields{
				code:    "hello",
				message: "world",
			},
			want: "hello",
		},
		{
			name: "foobar",
			fields: fields{
				code:    "foo",
				message: "bar",
			},
			want: "foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := validation.NewRuleError(tt.fields.code, tt.fields.message)
			if got := re.Code(); got != tt.want {
				t.Errorf("ruleError.Code() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ruleError_Message(t *testing.T) {
	type fields struct {
		code    string
		message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ab",
			fields: fields{
				code:    "a",
				message: "b",
			},
			want: "b",
		},
		{
			name: "hello",
			fields: fields{
				code:    "hello",
				message: "world",
			},
			want: "world",
		},
		{
			name: "foobar",
			fields: fields{
				code:    "foo",
				message: "bar",
			},
			want: "bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := validation.NewRuleError(tt.fields.code, tt.fields.message)
			if got := re.Message(); got != tt.want {
				t.Errorf("ruleError.Message() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_valueError_Error(t *testing.T) {
	type fields struct {
		name   string
		nested error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ab",
			fields: fields{
				name:   "a",
				nested: errors.New("b"),
			},
			want: "a: b",
		},
		{
			name: "hello",
			fields: fields{
				name:   "hello",
				nested: errors.New("world"),
			},
			want: "hello: world",
		},
		{
			name: "foobar",
			fields: fields{
				name:   "foo",
				nested: errors.New("bar"),
			},
			want: "foo: bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ve := validation.NewValueError(tt.fields.name, tt.fields.nested)
			if got := ve.Error(); got != tt.want {
				t.Errorf("valueError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_valueError_Name(t *testing.T) {
	type fields struct {
		name   string
		nested error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ab",
			fields: fields{
				name:   "a",
				nested: errors.New("b"),
			},
			want: "a",
		},
		{
			name: "hello",
			fields: fields{
				name:   "hello",
				nested: errors.New("world"),
			},
			want: "hello",
		},
		{
			name: "foobar",
			fields: fields{
				name:   "foo",
				nested: errors.New("bar"),
			},
			want: "foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ve := validation.NewValueError(tt.fields.name, tt.fields.nested)
			if got := ve.Name(); got != tt.want {
				t.Errorf("valueError.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_indexError_Error(t *testing.T) {
	type fields struct {
		index  int
		nested error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "0b",
			fields: fields{
				index:  0,
				nested: errors.New("b"),
			},
			want: "[0]: b",
		},
		{
			name: "hello",
			fields: fields{
				index:  1337,
				nested: errors.New("hello world"),
			},
			want: "[1337]: hello world",
		},
		{
			name: "foobar",
			fields: fields{
				index:  0xdeadbeef,
				nested: errors.New("foobar"),
			},
			want: "[3735928559]: foobar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ie := validation.NewIndexError(tt.fields.index, tt.fields.nested)
			if got := ie.Error(); got != tt.want {
				t.Errorf("indexError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_indexError_Index(t *testing.T) {
	type fields struct {
		index  int
		nested error
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "0b",
			fields: fields{
				index:  0,
				nested: errors.New("b"),
			},
			want: 0,
		},
		{
			name: "hello",
			fields: fields{
				index:  1337,
				nested: errors.New("hello world"),
			},
			want: 1337,
		},
		{
			name: "foobar",
			fields: fields{
				index:  0xdeadbeef,
				nested: errors.New("foobar"),
			},
			want: 0xdeadbeef,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ie := validation.NewIndexError(tt.fields.index, tt.fields.nested)
			if got := ie.Index(); got != tt.want {
				t.Errorf("indexError.Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrors_Error(t *testing.T) {
	tests := []struct {
		name string
		es   validation.Errors
		want string
	}{
		{
			name: "foobarbaz",
			es: []error{
				errors.New("foo"),
				errors.New("bar"),
				errors.New("baz"),
			},
			want: "",
		},
		{
			name: "validation",
			es: []error{
				validation.NewValueError("foo", errors.New("bar")),
				validation.NewIndexError(13, errors.New("out of bounds")),
			},
			want: "foo: bar",
		},
		{
			name: "rules",
			es: []error{
				validation.NewValueError("type", validation.NewRuleError("foo", "bar")),
				validation.NewValueError("data", validation.NewRuleError("baz", "quux")),
			},
			want: "type: bar; data: quux",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.es.Error(); got != tt.want {
				t.Errorf("Errors.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrors_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		es      validation.Errors
		want    []byte
		wantErr bool
	}{
		{
			name: "nested",
			es: []error{
				validation.NewValueError("foo", validation.NewIndexError(0, errors.New("bar"))),
				errors.New("B"),
				validation.NewValueError("baz", errors.New("quux")),
				errors.New("A"),
			},
			want:    []byte(`{"foo":{"0":"bar"},"baz":"quux"}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.es.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Errors.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Errors.MarshalJSON() = %s, want %s", got, tt.want)
			}
		})
	}
}
