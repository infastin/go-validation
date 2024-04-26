package validation

import (
	"slices"

	"github.com/infastin/go-validation/constraints"
)

type numberValidatorData[T constraints.Number] struct {
	value T
	name  string
}

type NumberValidator[T constraints.Number] struct {
	data  *numberValidatorData[T]
	rules []NumberRule[T]
	skip  bool
}

func Number[T constraints.Number](n T, name string) NumberValidator[T] {
	return NumberValidator[T]{
		data: &numberValidatorData[T]{
			value: n,
			name:  name,
		},
		rules: make([]NumberRule[T], 0),
		skip:  false,
	}
}

func NumberV[T constraints.Number]() NumberValidator[T] {
	return NumberValidator[T]{
		data:  nil,
		rules: make([]NumberRule[T], 0),
		skip:  false,
	}
}

func (nv NumberValidator[T]) Required(condition bool) NumberValidator[T] {
	if !nv.skip {
		nv.rules = append(nv.rules, Required[T](condition))
	}
	return nv
}

func (nv NumberValidator[T]) Skip(condition bool) NumberValidator[T] {
	if !nv.skip && condition {
		nv.skip = true
	}
	return nv
}

func (nv NumberValidator[T]) In(elements ...T) NumberValidator[T] {
	if !nv.skip {
		nv.rules = append(nv.rules, In(elements...))
	}
	return nv
}

func (nv NumberValidator[T]) NotIn(elements ...T) NumberValidator[T] {
	if !nv.skip {
		nv.rules = append(nv.rules, NotIn(elements...))
	}
	return nv
}

func (nv NumberValidator[T]) Equal(v T) NumberValidator[T] {
	if !nv.skip {
		nv.rules = append(nv.rules, Equal(v))
	}
	return nv
}

func (nv NumberValidator[T]) Less(v T) NumberValidator[T] {
	if !nv.skip {
		nv.rules = append(nv.rules, Less(v))
	}
	return nv
}

func (nv NumberValidator[T]) LessEqual(v T) NumberValidator[T] {
	if !nv.skip {
		nv.rules = append(nv.rules, LessEqual(v))
	}
	return nv
}

func (nv NumberValidator[T]) Greater(v T) NumberValidator[T] {
	if !nv.skip {
		nv.rules = append(nv.rules, Greater(v))
	}
	return nv
}

func (nv NumberValidator[T]) GreaterEqual(v T) NumberValidator[T] {
	if !nv.skip {
		nv.rules = append(nv.rules, GreaterEqual(v))
	}
	return nv
}

func (nv NumberValidator[T]) Between(a, b T) NumberValidator[T] {
	if !nv.skip {
		nv.rules = append(nv.rules, Between(a, b))
	}
	return nv
}

func (nv NumberValidator[T]) BetweenEqual(a, b T) NumberValidator[T] {
	if !nv.skip {
		nv.rules = append(nv.rules, BetweenEqual(a, b))
	}
	return nv
}

func (nv NumberValidator[T]) When(condition bool, ok, otherwise NumberRule[T]) NumberValidator[T] {
	if !nv.skip {
		if condition {
			nv.rules = append(nv.rules, ok)
		} else if otherwise != nil {
			nv.rules = append(nv.rules, otherwise)
		}
	}
	return nv
}

func (nv NumberValidator[T]) With(fns ...func(n T) error) NumberValidator[T] {
	if !nv.skip {
		slices.Grow(nv.rules, len(fns))
		for _, fn := range fns {
			nv.rules = append(nv.rules, NumberRuleFunc[T](fn))
		}
	}
	return nv
}

func (nv NumberValidator[T]) By(rules ...NumberRule[T]) NumberValidator[T] {
	if !nv.skip {
		nv.rules = append(nv.rules, rules...)
	}
	return nv
}

func (nv NumberValidator[T]) Valid() error {
	for _, rule := range nv.rules {
		if err := rule.Validate(nv.data.value); err != nil {
			return NewValueError(nv.data.name, err)
		}
	}
	return nil
}

func (nv NumberValidator[T]) Validate(v T) error {
	for _, rule := range nv.rules {
		if err := rule.Validate(v); err != nil {
			return err
		}
	}
	return nil
}
