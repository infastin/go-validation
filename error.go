package validation

import (
	"strconv"
	"strings"
)

type RuleError interface {
	Error() string
	Code() string
	Message() string
}

type ruleError struct {
	code    string
	message string
}

func NewRuleError(code, message string) RuleError {
	return &ruleError{
		code:    code,
		message: message,
	}
}

func (re *ruleError) Error() string {
	return re.message
}

func (re *ruleError) Code() string {
	return re.code
}

func (re *ruleError) Message() string {
	return re.message
}

type ValueError interface {
	Error() string
	Unwrap() error
	Name() string
}

type valueError struct {
	name   string
	nested error
}

func NewValueError(name string, nested error) ValueError {
	return &valueError{
		name:   name,
		nested: nested,
	}
}

func (ve *valueError) Error() string {
	msg := ve.nested.Error()

	var b strings.Builder

	b.Grow(len(ve.name) + 2 + len(msg))
	b.WriteString(ve.name)
	b.WriteString(": ")
	b.WriteString(msg)

	return b.String()
}

func (ve *valueError) Name() string {
	return ve.name
}

func (ve *valueError) Unwrap() error {
	return ve.nested
}

type IndexError interface {
	Error() string
	Unwrap() error
	Index() int
}

type indexError struct {
	index  int
	nested error
}

func NewIndexError(index int, nested error) IndexError {
	return &indexError{
		index:  index,
		nested: nested,
	}
}

func (ie *indexError) Error() string {
	msg := ie.nested.Error()
	idx := strconv.Itoa(ie.index)

	var b strings.Builder

	b.Grow(2 + len(idx) + 3 + len(msg) + 1)
	b.WriteString("([")
	b.WriteString(idx)
	b.WriteString("]: ")
	b.WriteString(msg)
	b.WriteByte(')')

	return b.String()
}

func (ie *indexError) Index() int {
	return ie.index
}

func (ie *indexError) Unwrap() error {
	return ie.nested
}

type WrapError interface {
	Error() string
	Unwrap() error
}

type wrapError struct {
	nested error
}

func NewWrapError(nested error) WrapError {
	return &wrapError{
		nested: nested,
	}
}

func (we *wrapError) Error() string {
	msg := we.nested.Error()

	var b strings.Builder

	b.Grow(1 + len(msg) + 1)
	b.WriteByte('(')
	b.WriteString(msg)
	b.WriteByte(')')

	return b.String()
}

func (we *wrapError) Unwrap() error {
	return we.nested
}

type Errors []error

func (es Errors) Error() string {
	if len(es) == 0 {
		return ""
	}

	var b strings.Builder

	for i, e := range es {
		if i != 0 {
			b.WriteString("; ")
		}
		b.WriteString(e.Error())
	}

	return b.String()
}

func errorMarshalJSON(err error, b []byte) []byte {
	switch e := err.(type) {
	case Errors:
		b = e.marshalJSON(b)
	case IndexError:
		b = append(b, '{', '"')
		b = strconv.AppendInt(b, int64(e.Index()), 10)
		b = append(b, '"', ':')
		b = errorMarshalJSON(e.Unwrap(), b)
		b = append(b, '}')
	case WrapError:
		err = e.Unwrap()
		b = errorMarshalJSON(err, b)
	default:
		b = strconv.AppendQuote(b, e.Error())
	}
	return b
}

func (es Errors) marshalJSON(b []byte) []byte {
	b = append(b, '{')

	for i, err := range es {
		ve, ok := err.(ValueError)
		if !ok {
			continue
		}

		if i != 0 {
			b = append(b, ',')
		}

		b = strconv.AppendQuote(b, ve.Name())
		b = append(b, ':')
		b = errorMarshalJSON(ve.Unwrap(), b)
	}

	b = append(b, '}')

	return b
}

func (es Errors) MarshalJSON() ([]byte, error) {
	return es.marshalJSON(nil), nil
}
