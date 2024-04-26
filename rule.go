package validation

import (
	"time"

	"github.com/infastin/go-validation/constraints"
)

type Validatable interface {
	Validate() error
}

type Rule[T any] interface {
	Validate(v T) error
}

type StringRule[T ~string] interface {
	Validate(s T) error
}

type StringRuleFunc[T ~string] func(s T) error

func (fn StringRuleFunc[T]) Validate(s T) error {
	return fn(s)
}

type IntRule[T constraints.Int] interface {
	Validate(i T) error
}

type IntRuleFunc[T constraints.Int] func(i T) error

func (fn IntRuleFunc[T]) Validate(i T) error {
	return fn(i)
}

type UintRule[T constraints.Uint] interface {
	Validate(u T) error
}

type UintRuleFunc[T constraints.Uint] func(u T) error

func (fn UintRuleFunc[T]) Validate(u T) error {
	return fn(u)
}

type FloatRule[T constraints.Float] interface {
	Validate(f T) error
}

type FloatRuleFunc[T constraints.Float] func(f T) error

func (fn FloatRuleFunc[T]) Validate(f T) error {
	return fn(f)
}

type TimeRule interface {
	Validate(t time.Time) error
}

type TimeRuleFunc func(t time.Time) error

func (fn TimeRuleFunc) Validate(t time.Time) error {
	return fn(t)
}

type PtrRule[T any] interface {
	Validate(p *T) error
}

type PtrRuleFunc[T any] func(p *T) error

func (fn PtrRuleFunc[T]) Validate(p *T) error {
	return fn(p)
}

type SliceRule[T any] interface {
	Validate(s []T) error
}

type SliceRuleFunc[T any] func(s []T) error

func (fn SliceRuleFunc[T]) Validate(s []T) error {
	return fn(s)
}

type MapRule[T any] interface {
	Validate(m map[string]T) error
}

type MapRuleFunc[T any] func(m map[string]T) error

func (fn MapRuleFunc[T]) Validate(m map[string]T) error {
	return fn(m)
}

type AnyRule[T any] interface {
	Validate(v T) error
}

type AnyRuleFunc[T any] func(v T) error

func (fn AnyRuleFunc[T]) Validate(v T) error {
	return fn(v)
}
