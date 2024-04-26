package validation

import "slices"

type anyValidatorData[T any] struct {
	value T
	name  string
}

type AnyValidator[T any] struct {
	data  *anyValidatorData[T]
	rules []AnyRule[T]
	skip  bool
}

func Any[T any](v T, name string) AnyValidator[T] {
	return AnyValidator[T]{
		data: &anyValidatorData[T]{
			value: v,
			name:  name,
		},
		rules: make([]AnyRule[T], 0),
		skip:  false,
	}
}

func AnyV[T any]() AnyValidator[T] {
	return AnyValidator[T]{
		data:  nil,
		rules: make([]AnyRule[T], 0),
		skip:  false,
	}
}

func (av AnyValidator[T]) Required(condition bool, isDefault func(v T) bool) AnyValidator[T] {
	if !av.skip {
		av.rules = append(av.rules, RequiredAny(condition, isDefault))
	}
	return av
}

func (av AnyValidator[T]) Skip(condition bool) AnyValidator[T] {
	if !av.skip && condition {
		av.skip = true
	}
	return av
}

func (av AnyValidator[T]) When(condition bool, ok, otherwise AnyRule[T]) AnyValidator[T] {
	if !av.skip {
		if condition {
			av.rules = append(av.rules, ok)
		} else {
			av.rules = append(av.rules, otherwise)
		}
	}
	return av
}

func (av AnyValidator[T]) With(fns ...func(v T) error) AnyValidator[T] {
	if !av.skip {
		slices.Grow(av.rules, len(fns))
		for _, fn := range fns {
			av.rules = append(av.rules, AnyRuleFunc[T](fn))
		}
	}
	return av
}

func (av AnyValidator[T]) By(rules ...AnyRule[T]) AnyValidator[T] {
	if !av.skip {
		av.rules = append(av.rules, rules...)
	}
	return av
}

func (av AnyValidator[T]) Valid() error {
	for _, rule := range av.rules {
		if err := rule.Validate(av.data.value); err != nil {
			return NewValueError(av.data.name, err)
		}
	}
	return nil
}

func (av AnyValidator[T]) Validate(v T) error {
	for _, rule := range av.rules {
		if err := rule.Validate(v); err != nil {
			return err
		}
	}
	return nil
}
