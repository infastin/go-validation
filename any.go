package validation

import "slices"

type anyValidatableData[T any] struct {
	value T
	name  string
}

type AnyValidation[T any] struct {
	data  *anyValidatableData[T]
	rules []AnyRule[T]
	skip  bool
}

func Any[T Validatable](v T, name string) AnyValidation[T] {
	return AnyValidation[T]{
		data: &anyValidatableData[T]{
			value: v,
			name:  name,
		},
		rules: make([]AnyRule[T], 0),
		skip:  false,
	}
}

func AnyV[T Validatable]() AnyValidation[T] {
	return AnyValidation[T]{
		data:  nil,
		rules: make([]AnyRule[T], 0),
		skip:  false,
	}
}

func (av AnyValidation[T]) Required(condition bool, isDefault func(v T) bool) AnyValidation[T] {
	if !av.skip {
		av.rules = append(av.rules, RequiredAny[T](condition, isDefault))
	}
	return av
}

func (av AnyValidation[T]) Skip(condition bool) AnyValidation[T] {
	if !av.skip && condition {
		av.skip = true
	}
	return av
}

func (av AnyValidation[T]) When(condition bool, ok AnyRule[T], otherwise AnyRule[T]) AnyValidation[T] {
	if !av.skip {
		if condition {
			av.rules = append(av.rules, ok)
		} else {
			av.rules = append(av.rules, otherwise)
		}
	}
	return av
}

func (av AnyValidation[T]) With(fns ...func(v T) error) AnyValidation[T] {
	if !av.skip {
		slices.Grow(av.rules, len(fns))
		for _, fn := range fns {
			av.rules = append(av.rules, AnyRuleFunc[T](fn))
		}
	}
	return av
}

func (av AnyValidation[T]) By(rules ...AnyRule[T]) AnyValidation[T] {
	if !av.skip {
		av.rules = append(av.rules, rules...)
	}
	return av
}

func (av AnyValidation[T]) Valid() error {
	for _, rule := range av.rules {
		if err := rule.Validate(av.data.value); err != nil {
			return NewValueError(av.data.name, err)
		}
	}
	return nil
}

func (av AnyValidation[T]) Validate(v T) error {
	for _, rule := range av.rules {
		if err := rule.Validate(v); err != nil {
			return err
		}
	}
	return nil
}
