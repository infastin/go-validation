package validation

import (
	"slices"

	"github.com/infastin/go-validation/constraints"
)

type floatValidatorData[T constraints.Float] struct {
	value T
	name  string
}

type FloatValidator[T constraints.Float] struct {
	data  *floatValidatorData[T]
	rules []FloatRule[T]
	skip  bool
}

func Float[T constraints.Float](i T, name string) FloatValidator[T] {
	return FloatValidator[T]{
		data: &floatValidatorData[T]{
			value: i,
			name:  name,
		},
		rules: make([]FloatRule[T], 0),
		skip:  false,
	}
}

func FloatV[T constraints.Float]() FloatValidator[T] {
	return FloatValidator[T]{
		data:  nil,
		rules: make([]FloatRule[T], 0),
		skip:  false,
	}
}

func (fv FloatValidator[T]) Required(condition bool) FloatValidator[T] {
	if !fv.skip {
		fv.rules = append(fv.rules, Required[T](condition))
	}
	return fv
}

func (fv FloatValidator[T]) Skip(condition bool) FloatValidator[T] {
	if !fv.skip && condition {
		fv.skip = true
	}
	return fv
}

func (fv FloatValidator[T]) In(elements ...T) FloatValidator[T] {
	if !fv.skip {
		fv.rules = append(fv.rules, In(elements...))
	}
	return fv
}

func (fv FloatValidator[T]) NotIn(elements ...T) FloatValidator[T] {
	if !fv.skip {
		fv.rules = append(fv.rules, NotIn(elements...))
	}
	return fv
}

func (fv FloatValidator[T]) Equal(v T) FloatValidator[T] {
	if !fv.skip {
		fv.rules = append(fv.rules, Equal(v))
	}
	return fv
}

func (fv FloatValidator[T]) Less(v T) FloatValidator[T] {
	if !fv.skip {
		fv.rules = append(fv.rules, Less(v))
	}
	return fv
}

func (fv FloatValidator[T]) LessEqual(v T) FloatValidator[T] {
	if !fv.skip {
		fv.rules = append(fv.rules, LessEqual(v))
	}
	return fv
}

func (fv FloatValidator[T]) Greater(v T) FloatValidator[T] {
	if !fv.skip {
		fv.rules = append(fv.rules, Greater(v))
	}
	return fv
}

func (fv FloatValidator[T]) GreaterEqual(v T) FloatValidator[T] {
	if !fv.skip {
		fv.rules = append(fv.rules, GreaterEqual(v))
	}
	return fv
}

func (fv FloatValidator[T]) Between(a, b T) FloatValidator[T] {
	if !fv.skip {
		fv.rules = append(fv.rules, Between(a, b))
	}
	return fv
}

func (fv FloatValidator[T]) BetweenEqual(a, b T) FloatValidator[T] {
	if !fv.skip {
		fv.rules = append(fv.rules, BetweenEqual(a, b))
	}
	return fv
}

func (fv FloatValidator[T]) When(condition bool, ok FloatRule[T], otherwise FloatRule[T]) FloatValidator[T] {
	if !fv.skip {
		if condition {
			fv.rules = append(fv.rules, ok)
		} else if otherwise != nil {
			fv.rules = append(fv.rules, otherwise)
		}
	}
	return fv
}

func (fv FloatValidator[T]) With(fns ...func(i T) error) FloatValidator[T] {
	if !fv.skip {
		slices.Grow(fv.rules, len(fns))
		for _, fn := range fns {
			fv.rules = append(fv.rules, FloatRuleFunc[T](fn))
		}
	}
	return fv
}

func (fv FloatValidator[T]) By(rules ...FloatRule[T]) FloatValidator[T] {
	if !fv.skip {
		fv.rules = append(fv.rules, rules...)
	}
	return fv
}

func (fv FloatValidator[T]) Valid() error {
	for _, rule := range fv.rules {
		if err := rule.Validate(fv.data.value); err != nil {
			return NewValueError(fv.data.name, err)
		}
	}
	return nil
}

func (fv FloatValidator[T]) Validate(v T) error {
	for _, rule := range fv.rules {
		if err := rule.Validate(v); err != nil {
			return err
		}
	}
	return nil
}
