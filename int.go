package validation

import (
	"slices"

	"github.com/infastin/go-validation/constraints"
)

type intValidatorData[T constraints.Int] struct {
	value T
	name  string
}

type IntValidator[T constraints.Int] struct {
	data  *intValidatorData[T]
	rules []IntRule[T]
	skip  bool
}

func Int[T constraints.Int](i T, name string) IntValidator[T] {
	return IntValidator[T]{
		data: &intValidatorData[T]{
			value: i,
			name:  name,
		},
		rules: make([]IntRule[T], 0),
		skip:  false,
	}
}

func IntV[T constraints.Int]() IntValidator[T] {
	return IntValidator[T]{
		data:  nil,
		rules: make([]IntRule[T], 0),
		skip:  false,
	}
}

func (iv IntValidator[T]) Required(condition bool) IntValidator[T] {
	if !iv.skip {
		iv.rules = append(iv.rules, Required[T](condition))
	}
	return iv
}

func (iv IntValidator[T]) Skip(condition bool) IntValidator[T] {
	if !iv.skip && condition {
		iv.skip = true
	}
	return iv
}

func (iv IntValidator[T]) In(elements ...T) IntValidator[T] {
	if !iv.skip {
		iv.rules = append(iv.rules, In(elements...))
	}
	return iv
}

func (iv IntValidator[T]) NotIn(elements ...T) IntValidator[T] {
	if !iv.skip {
		iv.rules = append(iv.rules, NotIn(elements...))
	}
	return iv
}

func (iv IntValidator[T]) Equal(v T) IntValidator[T] {
	if !iv.skip {
		iv.rules = append(iv.rules, Equal(v))
	}
	return iv
}

func (iv IntValidator[T]) Less(v T) IntValidator[T] {
	if !iv.skip {
		iv.rules = append(iv.rules, Less(v))
	}
	return iv
}

func (iv IntValidator[T]) LessEqual(v T) IntValidator[T] {
	if !iv.skip {
		iv.rules = append(iv.rules, LessEqual(v))
	}
	return iv
}

func (iv IntValidator[T]) Greater(v T) IntValidator[T] {
	if !iv.skip {
		iv.rules = append(iv.rules, Greater(v))
	}
	return iv
}

func (iv IntValidator[T]) GreaterEqual(v T) IntValidator[T] {
	if !iv.skip {
		iv.rules = append(iv.rules, GreaterEqual(v))
	}
	return iv
}

func (iv IntValidator[T]) Between(a, b T) IntValidator[T] {
	if !iv.skip {
		iv.rules = append(iv.rules, Between(a, b))
	}
	return iv
}

func (iv IntValidator[T]) BetweenEqual(a, b T) IntValidator[T] {
	if !iv.skip {
		iv.rules = append(iv.rules, BetweenEqual(a, b))
	}
	return iv
}

func (iv IntValidator[T]) When(condition bool, ok IntRule[T], otherwise IntRule[T]) IntValidator[T] {
	if !iv.skip {
		if condition {
			iv.rules = append(iv.rules, ok)
		} else if otherwise != nil {
			iv.rules = append(iv.rules, otherwise)
		}
	}
	return iv
}

func (iv IntValidator[T]) With(fns ...func(i T) error) IntValidator[T] {
	if !iv.skip {
		slices.Grow(iv.rules, len(fns))
		for _, fn := range fns {
			iv.rules = append(iv.rules, IntRuleFunc[T](fn))
		}
	}
	return iv
}

func (iv IntValidator[T]) By(rules ...IntRule[T]) IntValidator[T] {
	if !iv.skip {
		iv.rules = append(iv.rules, rules...)
	}
	return iv
}

func (iv IntValidator[T]) Valid() error {
	for _, rule := range iv.rules {
		if err := rule.Validate(iv.data.value); err != nil {
			return NewValueError(iv.data.name, err)
		}
	}
	return nil
}

func (iv IntValidator[T]) Validate(v T) error {
	for _, rule := range iv.rules {
		if err := rule.Validate(v); err != nil {
			return err
		}
	}
	return nil
}
