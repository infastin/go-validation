package validation

import "slices"

type customValidatableData[T Validatable] struct {
	value T
	name  string
}

type CustomValidation[T Validatable] struct {
	data  *customValidatableData[T]
	rules []AnyRule[T]
	skip  bool
	wrap  bool
}

func Custom[T Validatable](v T, name string) CustomValidation[T] {
	return CustomValidation[T]{
		data: &customValidatableData[T]{
			value: v,
			name:  name,
		},
		rules: make([]AnyRule[T], 0),
		skip:  false,
		wrap:  false,
	}
}

func CustomV[T Validatable]() CustomValidation[T] {
	return CustomValidation[T]{
		data:  nil,
		rules: make([]AnyRule[T], 0),
		skip:  false,
		wrap:  false,
	}
}

func (cv CustomValidation[T]) Required(condition bool, isDefault func(v T) bool) CustomValidation[T] {
	if !cv.skip {
		cv.rules = append(cv.rules, RequiredAny[T](condition, isDefault))
	}
	return cv
}

func (cv CustomValidation[T]) Skip(condition bool) CustomValidation[T] {
	if !cv.skip && condition {
		cv.skip = true
	}
	return cv
}

func (cv CustomValidation[T]) Wrap() CustomValidation[T] {
	cv.wrap = true
	return cv
}

func (cv CustomValidation[T]) When(condition bool, ok AnyRule[T], otherwise AnyRule[T]) CustomValidation[T] {
	if !cv.skip {
		if condition {
			cv.rules = append(cv.rules, ok)
		} else {
			cv.rules = append(cv.rules, otherwise)
		}
	}
	return cv
}

func (cv CustomValidation[T]) With(fns ...func(v T) error) CustomValidation[T] {
	if !cv.skip {
		slices.Grow(cv.rules, len(fns))
		for _, fn := range fns {
			cv.rules = append(cv.rules, AnyRuleFunc[T](fn))
		}
	}
	return cv
}

func (cv CustomValidation[T]) By(rules ...AnyRule[T]) CustomValidation[T] {
	if !cv.skip {
		cv.rules = append(cv.rules, rules...)
	}
	return cv
}

func (cv CustomValidation[T]) Valid() error {
	for _, rule := range cv.rules {
		if err := rule.Validate(cv.data.value); err != nil {
			return NewValueError(cv.data.name, err)
		}
	}
	if err := cv.data.value.Validate(); err != nil {
		if cv.wrap {
			err = NewWrapError(err)
		}
		return NewValueError(cv.data.name, err)
	}
	return nil
}

func (cv CustomValidation[T]) Validate(v T) error {
	for _, rule := range cv.rules {
		if err := rule.Validate(v); err != nil {
			return err
		}
	}
	return v.Validate()
}
