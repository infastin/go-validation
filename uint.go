package validation

import (
	"slices"

	"github.com/infastin/go-validation/constraints"
)

type uintValidatorData[T constraints.Uint] struct {
	value T
	name  string
}

type UintValidator[T constraints.Uint] struct {
	data  *uintValidatorData[T]
	rules []UintRule[T]
	skip  bool
}

func Uint[T constraints.Uint](i T, name string) UintValidator[T] {
	return UintValidator[T]{
		data: &uintValidatorData[T]{
			value: i,
			name:  name,
		},
		rules: make([]UintRule[T], 0),
		skip:  false,
	}
}

func UintV[T constraints.Uint]() UintValidator[T] {
	return UintValidator[T]{
		data:  nil,
		rules: make([]UintRule[T], 0),
		skip:  false,
	}
}

func (uv UintValidator[T]) Required(condition bool) UintValidator[T] {
	if !uv.skip {
		uv.rules = append(uv.rules, Required[T](condition))
	}
	return uv
}

func (uv UintValidator[T]) Skip(condition bool) UintValidator[T] {
	if !uv.skip && condition {
		uv.skip = true
	}
	return uv
}

func (uv UintValidator[T]) In(elements ...T) UintValidator[T] {
	if !uv.skip {
		uv.rules = append(uv.rules, In(elements...))
	}
	return uv
}

func (uv UintValidator[T]) NotIn(elements ...T) UintValidator[T] {
	if !uv.skip {
		uv.rules = append(uv.rules, NotIn(elements...))
	}
	return uv
}

func (uv UintValidator[T]) Equal(v T) UintValidator[T] {
	if !uv.skip {
		uv.rules = append(uv.rules, Equal(v))
	}
	return uv
}

func (uv UintValidator[T]) Less(v T) UintValidator[T] {
	if !uv.skip {
		uv.rules = append(uv.rules, Less(v))
	}
	return uv
}

func (uv UintValidator[T]) LessEqual(v T) UintValidator[T] {
	if !uv.skip {
		uv.rules = append(uv.rules, LessEqual(v))
	}
	return uv
}

func (uv UintValidator[T]) Greater(v T) UintValidator[T] {
	if !uv.skip {
		uv.rules = append(uv.rules, Greater(v))
	}
	return uv
}

func (uv UintValidator[T]) GreaterEqual(v T) UintValidator[T] {
	if !uv.skip {
		uv.rules = append(uv.rules, GreaterEqual(v))
	}
	return uv
}

func (uv UintValidator[T]) Between(a, b T) UintValidator[T] {
	if !uv.skip {
		uv.rules = append(uv.rules, Between(a, b))
	}
	return uv
}

func (uv UintValidator[T]) BetweenEqual(a, b T) UintValidator[T] {
	if !uv.skip {
		uv.rules = append(uv.rules, BetweenEqual(a, b))
	}
	return uv
}

func (uv UintValidator[T]) When(condition bool, ok UintRule[T], otherwise UintRule[T]) UintValidator[T] {
	if !uv.skip {
		if condition {
			uv.rules = append(uv.rules, ok)
		} else if otherwise != nil {
			uv.rules = append(uv.rules, otherwise)
		}
	}
	return uv
}

func (uv UintValidator[T]) With(fns ...func(i T) error) UintValidator[T] {
	if !uv.skip {
		slices.Grow(uv.rules, len(fns))
		for _, fn := range fns {
			uv.rules = append(uv.rules, UintRuleFunc[T](fn))
		}
	}
	return uv
}

func (uv UintValidator[T]) By(rules ...UintRule[T]) UintValidator[T] {
	if !uv.skip {
		uv.rules = append(uv.rules, rules...)
	}
	return uv
}

func (uv UintValidator[T]) Valid() error {
	for _, rule := range uv.rules {
		if err := rule.Validate(uv.data.value); err != nil {
			return NewValueError(uv.data.name, err)
		}
	}
	return nil
}

func (uv UintValidator[T]) Validate(v T) error {
	for _, rule := range uv.rules {
		if err := rule.Validate(v); err != nil {
			return err
		}
	}
	return nil
}
