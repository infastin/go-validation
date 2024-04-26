package validation

import "time"

type requiredRule[T comparable] struct {
	condition bool
}

var ErrRequired = NewRuleError("required", "cannot be blank")

func Example[T comparable](v T) error {
	if v == *new(T) {
		return ErrRequired
	}
	return nil
}

func Required[T comparable](condition bool) requiredRule[T] {
	return requiredRule[T]{
		condition: condition,
	}
}

func (r requiredRule[T]) Validate(v T) error {
	if r.condition && v == *new(T) {
		return ErrRequired
	}
	return nil
}

type requiredTimeRule struct {
	condition bool
}

func RequiredTime(condition bool) requiredTimeRule {
	return requiredTimeRule{
		condition: condition,
	}
}

func (r requiredTimeRule) Validate(v time.Time) error {
	if r.condition && v.IsZero() {
		return ErrRequired
	}
	return nil
}

type requiredAnyRule[T any] struct {
	condition bool
	isDefault func(a T) bool
}

func RequiredAny[T any](condition bool, isDefault func(a T) bool) requiredAnyRule[T] {
	return requiredAnyRule[T]{
		condition: condition,
		isDefault: isDefault,
	}
}

func (r requiredAnyRule[T]) Validate(v T) error {
	if r.condition && r.isDefault(v) {
		return ErrRequired
	}
	return nil
}
