package validation

import "slices"

type stringValidationData[T ~string] struct {
	value T
	name  string
}

type StringValidation[T ~string] struct {
	data  *stringValidationData[T]
	rules []StringRule[T]
	skip  bool
}

func String[T ~string](s T, name string) StringValidation[T] {
	return StringValidation[T]{
		data: &stringValidationData[T]{
			value: s,
			name:  name,
		},
		rules: make([]StringRule[T], 0),
		skip:  false,
	}
}

func StringV[T ~string]() StringValidation[T] {
	return StringValidation[T]{
		data:  nil,
		rules: make([]StringRule[T], 0),
		skip:  false,
	}
}

func (sv StringValidation[T]) Required(condition bool) StringValidation[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, Required[T](condition))
	}
	return sv
}

func (sv StringValidation[T]) Skip(condition bool) StringValidation[T] {
	if !sv.skip && condition {
		sv.skip = true
	}
	return sv
}

func (sv StringValidation[T]) Length(min, max int) StringValidation[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, LengthString[T](min, max))
	}
	return sv
}

func (sv StringValidation[T]) In(elements ...T) StringValidation[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, In(elements...))
	}
	return sv
}

func (sv StringValidation[T]) NotIn(elements ...T) StringValidation[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, NotIn(elements...))
	}
	return sv
}

func (sv StringValidation[T]) Equal(v T) StringValidation[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, Equal(v))
	}
	return sv
}

func (sv StringValidation[T]) Less(v T) StringValidation[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, Less(v))
	}
	return sv
}

func (sv StringValidation[T]) LessEqual(v T) StringValidation[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, LessEqual(v))
	}
	return sv
}

func (sv StringValidation[T]) Greater(v T) StringValidation[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, Greater(v))
	}
	return sv
}

func (sv StringValidation[T]) GreaterEqual(v T) StringValidation[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, GreaterEqual(v))
	}
	return sv
}

func (sv StringValidation[T]) Between(a, b T) StringValidation[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, Between(a, b))
	}
	return sv
}

func (sv StringValidation[T]) BetweenEqual(a, b T) StringValidation[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, BetweenEqual(a, b))
	}
	return sv
}

func (sv StringValidation[T]) When(condition bool, ok StringRule[T], otherwise StringRule[T]) StringValidation[T] {
	if !sv.skip {
		if condition {
			sv.rules = append(sv.rules, ok)
		} else if otherwise != nil {
			sv.rules = append(sv.rules, otherwise)
		}
	}
	return sv
}

func (sv StringValidation[T]) With(fns ...func(s T) error) StringValidation[T] {
	if !sv.skip {
		slices.Grow(sv.rules, len(fns))
		for _, fn := range fns {
			sv.rules = append(sv.rules, StringRuleFunc[T](fn))
		}
	}
	return sv
}

func (sv StringValidation[T]) By(rules ...StringRule[T]) StringValidation[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, rules...)
	}
	return sv
}

func (sv StringValidation[T]) Valid() error {
	for _, rule := range sv.rules {
		if err := rule.Validate(sv.data.value); err != nil {
			return NewValueError(sv.data.name, err)
		}
	}
	return nil
}

func (sv StringValidation[T]) Validate(v T) error {
	for _, rule := range sv.rules {
		if err := rule.Validate(v); err != nil {
			return err
		}
	}
	return nil
}
