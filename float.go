package validation

import (
	"slices"

	"github.com/infastin/go-validation/constraints"
)

type floatValidationData[T constraints.Float] struct {
	value T
	name  string
}

type FloatValidation[T constraints.Float] struct {
	data  *floatValidationData[T]
	rules []FloatRule[T]
	skip  bool
}

func Float[T constraints.Float](i T, name string) FloatValidation[T] {
	return FloatValidation[T]{
		data: &floatValidationData[T]{
			value: i,
			name:  name,
		},
		rules: make([]FloatRule[T], 0),
		skip:  false,
	}
}

func FloatV[T constraints.Float]() FloatValidation[T] {
	return FloatValidation[T]{
		data:  nil,
		rules: make([]FloatRule[T], 0),
		skip:  false,
	}
}

func (fv FloatValidation[T]) Required(condition bool) FloatValidation[T] {
	if !fv.skip {
		fv.rules = append(fv.rules, Required[T](condition))
	}
	return fv
}

func (fv FloatValidation[T]) Skip(condition bool) FloatValidation[T] {
	if !fv.skip && condition {
		fv.skip = true
	}
	return fv
}

func (fv FloatValidation[T]) In(elements ...T) FloatValidation[T] {
	if !fv.skip {
		fv.rules = append(fv.rules, In(elements...))
	}
	return fv
}

func (fv FloatValidation[T]) NotIn(elements ...T) FloatValidation[T] {
	if !fv.skip {
		fv.rules = append(fv.rules, NotIn(elements...))
	}
	return fv
}

func (fv FloatValidation[T]) Equal(v T) FloatValidation[T] {
	if !fv.skip {
		fv.rules = append(fv.rules, Equal(v))
	}
	return fv
}

func (fv FloatValidation[T]) Less(v T) FloatValidation[T] {
	if !fv.skip {
		fv.rules = append(fv.rules, Less(v))
	}
	return fv
}

func (fv FloatValidation[T]) LessEqual(v T) FloatValidation[T] {
	if !fv.skip {
		fv.rules = append(fv.rules, LessEqual(v))
	}
	return fv
}

func (fv FloatValidation[T]) Greater(v T) FloatValidation[T] {
	if !fv.skip {
		fv.rules = append(fv.rules, Greater(v))
	}
	return fv
}

func (fv FloatValidation[T]) GreaterEqual(v T) FloatValidation[T] {
	if !fv.skip {
		fv.rules = append(fv.rules, GreaterEqual(v))
	}
	return fv
}

func (fv FloatValidation[T]) Between(a, b T) FloatValidation[T] {
	if !fv.skip {
		fv.rules = append(fv.rules, Between(a, b))
	}
	return fv
}

func (fv FloatValidation[T]) BetweenEqual(a, b T) FloatValidation[T] {
	if !fv.skip {
		fv.rules = append(fv.rules, BetweenEqual(a, b))
	}
	return fv
}

func (fv FloatValidation[T]) When(condition bool, ok FloatRule[T], otherwise FloatRule[T]) FloatValidation[T] {
	if !fv.skip {
		if condition {
			fv.rules = append(fv.rules, ok)
		} else if otherwise != nil {
			fv.rules = append(fv.rules, otherwise)
		}
	}
	return fv
}

func (fv FloatValidation[T]) With(fns ...func(i T) error) FloatValidation[T] {
	if !fv.skip {
		slices.Grow(fv.rules, len(fns))
		for _, fn := range fns {
			fv.rules = append(fv.rules, FloatRuleFunc[T](fn))
		}
	}
	return fv
}

func (fv FloatValidation[T]) By(rules ...FloatRule[T]) FloatValidation[T] {
	if !fv.skip {
		fv.rules = append(fv.rules, rules...)
	}
	return fv
}

func (fv FloatValidation[T]) Valid() error {
	for _, rule := range fv.rules {
		if err := rule.Validate(fv.data.value); err != nil {
			return NewValueError(fv.data.name, err)
		}
	}
	return nil
}

func (fv FloatValidation[T]) Validate(v T) error {
	for _, rule := range fv.rules {
		if err := rule.Validate(v); err != nil {
			return err
		}
	}
	return nil
}
