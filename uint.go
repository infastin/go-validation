package validation

import (
	"slices"

	"github.com/infastin/go-validation/constraints"
)

type uintValidationData[T constraints.Uint] struct {
	value T
	name  string
}

type UintValidation[T constraints.Uint] struct {
	data  *uintValidationData[T]
	rules []UintRule[T]
	skip  bool
}

func Uint[T constraints.Uint](i T, name string) UintValidation[T] {
	return UintValidation[T]{
		data: &uintValidationData[T]{
			value: i,
			name:  name,
		},
		rules: make([]UintRule[T], 0),
		skip:  false,
	}
}

func UintV[T constraints.Uint]() UintValidation[T] {
	return UintValidation[T]{
		data:  nil,
		rules: make([]UintRule[T], 0),
		skip:  false,
	}
}

func (uv UintValidation[T]) Required(condition bool) UintValidation[T] {
	if !uv.skip {
		uv.rules = append(uv.rules, Required[T](condition))
	}
	return uv
}

func (uv UintValidation[T]) Skip(condition bool) UintValidation[T] {
	if !uv.skip && condition {
		uv.skip = true
	}
	return uv
}

func (uv UintValidation[T]) In(elements ...T) UintValidation[T] {
	if !uv.skip {
		uv.rules = append(uv.rules, In(elements...))
	}
	return uv
}

func (uv UintValidation[T]) NotIn(elements ...T) UintValidation[T] {
	if !uv.skip {
		uv.rules = append(uv.rules, NotIn(elements...))
	}
	return uv
}

func (uv UintValidation[T]) Equal(v T) UintValidation[T] {
	if !uv.skip {
		uv.rules = append(uv.rules, Equal(v))
	}
	return uv
}

func (uv UintValidation[T]) Less(v T) UintValidation[T] {
	if !uv.skip {
		uv.rules = append(uv.rules, Less(v))
	}
	return uv
}

func (uv UintValidation[T]) LessEqual(v T) UintValidation[T] {
	if !uv.skip {
		uv.rules = append(uv.rules, LessEqual(v))
	}
	return uv
}

func (uv UintValidation[T]) Greater(v T) UintValidation[T] {
	if !uv.skip {
		uv.rules = append(uv.rules, Greater(v))
	}
	return uv
}

func (uv UintValidation[T]) GreaterEqual(v T) UintValidation[T] {
	if !uv.skip {
		uv.rules = append(uv.rules, GreaterEqual(v))
	}
	return uv
}

func (uv UintValidation[T]) Between(a, b T) UintValidation[T] {
	if !uv.skip {
		uv.rules = append(uv.rules, Between(a, b))
	}
	return uv
}

func (uv UintValidation[T]) BetweenEqual(a, b T) UintValidation[T] {
	if !uv.skip {
		uv.rules = append(uv.rules, BetweenEqual(a, b))
	}
	return uv
}

func (uv UintValidation[T]) When(condition bool, ok UintRule[T], otherwise UintRule[T]) UintValidation[T] {
	if !uv.skip {
		if condition {
			uv.rules = append(uv.rules, ok)
		} else if otherwise != nil {
			uv.rules = append(uv.rules, otherwise)
		}
	}
	return uv
}

func (uv UintValidation[T]) With(fns ...func(i T) error) UintValidation[T] {
	if !uv.skip {
		slices.Grow(uv.rules, len(fns))
		for _, fn := range fns {
			uv.rules = append(uv.rules, UintRuleFunc[T](fn))
		}
	}
	return uv
}

func (uv UintValidation[T]) By(rules ...UintRule[T]) UintValidation[T] {
	if !uv.skip {
		uv.rules = append(uv.rules, rules...)
	}
	return uv
}

func (uv UintValidation[T]) Valid() error {
	for _, rule := range uv.rules {
		if err := rule.Validate(uv.data.value); err != nil {
			return NewValueError(uv.data.name, err)
		}
	}
	return nil
}

func (uv UintValidation[T]) Validate(v T) error {
	for _, rule := range uv.rules {
		if err := rule.Validate(v); err != nil {
			return err
		}
	}
	return nil
}
