package validation

import (
	"errors"
	"reflect"
	"testing"
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
			re := &ruleError{
				code:    tt.fields.code,
				message: tt.fields.message,
			}
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
			re := &ruleError{
				code:    tt.fields.code,
				message: tt.fields.message,
			}
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
			re := &ruleError{
				code:    tt.fields.code,
				message: tt.fields.message,
			}
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
			ve := &valueError{
				name:   tt.fields.name,
				nested: tt.fields.nested,
			}
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
			ve := &valueError{
				name:   tt.fields.name,
				nested: tt.fields.nested,
			}
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
			want: "([0]: b)",
		},
		{
			name: "hello",
			fields: fields{
				index:  1337,
				nested: errors.New("hello world"),
			},
			want: "([1337]: hello world)",
		},
		{
			name: "foobar",
			fields: fields{
				index:  0xdeadbeef,
				nested: errors.New("foobar"),
			},
			want: "([3735928559]: foobar)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ie := &indexError{
				index:  tt.fields.index,
				nested: tt.fields.nested,
			}
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
			ie := &indexError{
				index:  tt.fields.index,
				nested: tt.fields.nested,
			}
			if got := ie.Index(); got != tt.want {
				t.Errorf("indexError.Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wrapError_Error(t *testing.T) {
	type fields struct {
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
				nested: errors.New("b"),
			},
			want: "(b)",
		},
		{
			name: "hello",
			fields: fields{
				nested: errors.New("world"),
			},
			want: "(world)",
		},
		{
			name: "foobar",
			fields: fields{
				nested: errors.New("bar"),
			},
			want: "(bar)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			we := &wrapError{
				nested: tt.fields.nested,
			}
			if got := we.Error(); got != tt.want {
				t.Errorf("wrapError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrors_Error(t *testing.T) {
	tests := []struct {
		name string
		es   Errors
		want string
	}{
		{
			name: "foobarbaz",
			es: []error{
				errors.New("foo"),
				errors.New("bar"),
				errors.New("baz"),
			},
			want: "foo; bar; baz",
		},
		{
			name: "validation",
			es: []error{
				&valueError{
					name:   "foo",
					nested: errors.New("bar"),
				},
				&wrapError{
					nested: errors.New("all your codebase are belong to us"),
				},
				&indexError{
					index:  13,
					nested: errors.New("out of bounds"),
				},
			},
			want: "foo: bar; (all your codebase are belong to us); ([13]: out of bounds)",
		},
		{
			name: "rules",
			es: []error{
				&valueError{
					name:   "type",
					nested: NewRuleError("foo", "bar"),
				},
				&valueError{
					name:   "data",
					nested: NewRuleError("baz", "quux"),
				},
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
		es      Errors
		want    []byte
		wantErr bool
	}{
		{
			name: "nested",
			es: []error{
				&valueError{
					name: "foo",
					nested: &indexError{
						index:  0,
						nested: errors.New("bar"),
					},
				},
				errors.New("B"),
				&valueError{
					name:   "baz",
					nested: errors.New("quux"),
				},
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
