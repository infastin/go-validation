package validation

import (
	"slices"

	"github.com/infastin/go-validation/constraints"
)

type intValidationData[T constraints.Int] struct {
	value T
	name  string
}

type IntValidation[T constraints.Int] struct {
	data  *intValidationData[T]
	rules []IntRule[T]
	skip  bool
}

func Int[T constraints.Int](i T, name string) IntValidation[T] {
	return IntValidation[T]{
		data: &intValidationData[T]{
			value: i,
			name:  name,
		},
		rules: make([]IntRule[T], 0),
		skip:  false,
	}
}

func IntV[T constraints.Int]() IntValidation[T] {
	return IntValidation[T]{
		data:  nil,
		rules: make([]IntRule[T], 0),
		skip:  false,
	}
}

func (iv IntValidation[T]) Required(condition bool) IntValidation[T] {
	if !iv.skip {
		iv.rules = append(iv.rules, Required[T](condition))
	}
	return iv
}

func (iv IntValidation[T]) Skip(condition bool) IntValidation[T] {
	if !iv.skip && condition {
		iv.skip = true
	}
	return iv
}

func (iv IntValidation[T]) In(elements ...T) IntValidation[T] {
	if !iv.skip {
		iv.rules = append(iv.rules, In(elements...))
	}
	return iv
}

func (iv IntValidation[T]) NotIn(elements ...T) IntValidation[T] {
	if !iv.skip {
		iv.rules = append(iv.rules, NotIn(elements...))
	}
	return iv
}

func (iv IntValidation[T]) Equal(v T) IntValidation[T] {
	if !iv.skip {
		iv.rules = append(iv.rules, Equal(v))
	}
	return iv
}

func (iv IntValidation[T]) Less(v T) IntValidation[T] {
	if !iv.skip {
		iv.rules = append(iv.rules, Less(v))
	}
	return iv
}

func (iv IntValidation[T]) LessEqual(v T) IntValidation[T] {
	if !iv.skip {
		iv.rules = append(iv.rules, LessEqual(v))
	}
	return iv
}

func (iv IntValidation[T]) Greater(v T) IntValidation[T] {
	if !iv.skip {
		iv.rules = append(iv.rules, Greater(v))
	}
	return iv
}

func (iv IntValidation[T]) GreaterEqual(v T) IntValidation[T] {
	if !iv.skip {
		iv.rules = append(iv.rules, GreaterEqual(v))
	}
	return iv
}

func (iv IntValidation[T]) Between(a, b T) IntValidation[T] {
	if !iv.skip {
		iv.rules = append(iv.rules, Between(a, b))
	}
	return iv
}

func (iv IntValidation[T]) BetweenEqual(a, b T) IntValidation[T] {
	if !iv.skip {
		iv.rules = append(iv.rules, BetweenEqual(a, b))
	}
	return iv
}

func (iv IntValidation[T]) When(condition bool, ok IntRule[T], otherwise IntRule[T]) IntValidation[T] {
	if !iv.skip {
		if condition {
			iv.rules = append(iv.rules, ok)
		} else if otherwise != nil {
			iv.rules = append(iv.rules, otherwise)
		}
	}
	return iv
}

func (iv IntValidation[T]) With(fns ...func(i T) error) IntValidation[T] {
	if !iv.skip {
		slices.Grow(iv.rules, len(fns))
		for _, fn := range fns {
			iv.rules = append(iv.rules, IntRuleFunc[T](fn))
		}
	}
	return iv
}

func (iv IntValidation[T]) By(rules ...IntRule[T]) IntValidation[T] {
	if !iv.skip {
		iv.rules = append(iv.rules, rules...)
	}
	return iv
}

func (iv IntValidation[T]) Valid() error {
	for _, rule := range iv.rules {
		if err := rule.Validate(iv.data.value); err != nil {
			return NewValueError(iv.data.name, err)
		}
	}
	return nil
}

func (iv IntValidation[T]) Validate(v T) error {
	for _, rule := range iv.rules {
		if err := rule.Validate(v); err != nil {
			return err
		}
	}
	return nil
}
