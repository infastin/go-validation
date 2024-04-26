package validation

import "slices"

type comparableValidatorData[T comparable] struct {
	value T
	name  string
}

type ComparableValidator[T comparable] struct {
	data  *comparableValidatorData[T]
	rules []ComparableRule[T]
	skip  bool
}

func Comparable[T comparable](v T, name string) ComparableValidator[T] {
	return ComparableValidator[T]{
		data: &comparableValidatorData[T]{
			value: v,
			name:  name,
		},
		rules: make([]ComparableRule[T], 0),
		skip:  false,
	}
}

func ComparableV[T comparable]() ComparableValidator[T] {
	return ComparableValidator[T]{
		data:  nil,
		rules: make([]ComparableRule[T], 0),
		skip:  false,
	}
}

func (cv ComparableValidator[T]) Required(condition bool) ComparableValidator[T] {
	if !cv.skip {
		cv.rules = append(cv.rules, Required[T](condition))
	}
	return cv
}

func (cv ComparableValidator[T]) Skip(condition bool) ComparableValidator[T] {
	if !cv.skip && condition {
		cv.skip = true
	}
	return cv
}

func (cv ComparableValidator[T]) In(elements ...T) ComparableValidator[T] {
	if !cv.skip {
		cv.rules = append(cv.rules, In(elements...))
	}
	return cv
}

func (cv ComparableValidator[T]) NotIn(elements ...T) ComparableValidator[T] {
	if !cv.skip {
		cv.rules = append(cv.rules, NotIn(elements...))
	}
	return cv
}

func (cv ComparableValidator[T]) Equal(v T) ComparableValidator[T] {
	if !cv.skip {
		cv.rules = append(cv.rules, Equal[T](v))
	}
	return cv
}

func (cv ComparableValidator[T]) When(condition bool, ok, otherwise ComparableRule[T]) ComparableValidator[T] {
	if !cv.skip {
		if condition {
			cv.rules = append(cv.rules, ok)
		} else if otherwise != nil {
			cv.rules = append(cv.rules, otherwise)
		}
	}
	return cv
}

func (cv ComparableValidator[T]) With(fns ...func(v T) error) ComparableValidator[T] {
	if !cv.skip {
		slices.Grow(cv.rules, len(fns))
		for _, fn := range fns {
			cv.rules = append(cv.rules, ComparableRuleFunc[T](fn))
		}
	}
	return cv
}

func (cv ComparableValidator[T]) By(rules ...ComparableRule[T]) ComparableValidator[T] {
	if !cv.skip {
		cv.rules = append(cv.rules, rules...)
	}
	return cv
}

func (cv ComparableValidator[T]) Valid() error {
	for _, rule := range cv.rules {
		if err := rule.Validate(cv.data.value); err != nil {
			return NewValueError(cv.data.name, err)
		}
	}
	return nil
}

func (cv ComparableValidator[T]) Validate(v T) error {
	for _, rule := range cv.rules {
		if err := rule.Validate(v); err != nil {
			return err
		}
	}
	return nil
}
