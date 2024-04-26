package validation

import "slices"

type stringValidatorData[T ~string] struct {
	value T
	name  string
}

type StringValidator[T ~string] struct {
	data  *stringValidatorData[T]
	rules []StringRule[T]
	skip  bool
}

func String[T ~string](s T, name string) StringValidator[T] {
	return StringValidator[T]{
		data: &stringValidatorData[T]{
			value: s,
			name:  name,
		},
		rules: make([]StringRule[T], 0),
		skip:  false,
	}
}

func StringV[T ~string]() StringValidator[T] {
	return StringValidator[T]{
		data:  nil,
		rules: make([]StringRule[T], 0),
		skip:  false,
	}
}

func (sv StringValidator[T]) Required(condition bool) StringValidator[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, Required[T](condition))
	}
	return sv
}

func (sv StringValidator[T]) Skip(condition bool) StringValidator[T] {
	if !sv.skip && condition {
		sv.skip = true
	}
	return sv
}

func (sv StringValidator[T]) Length(min, max int) StringValidator[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, LengthString[T](min, max))
	}
	return sv
}

func (sv StringValidator[T]) In(elements ...T) StringValidator[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, In(elements...))
	}
	return sv
}

func (sv StringValidator[T]) NotIn(elements ...T) StringValidator[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, NotIn(elements...))
	}
	return sv
}

func (sv StringValidator[T]) Equal(v T) StringValidator[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, Equal(v))
	}
	return sv
}

func (sv StringValidator[T]) Less(v T) StringValidator[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, Less(v))
	}
	return sv
}

func (sv StringValidator[T]) LessEqual(v T) StringValidator[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, LessEqual(v))
	}
	return sv
}

func (sv StringValidator[T]) Greater(v T) StringValidator[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, Greater(v))
	}
	return sv
}

func (sv StringValidator[T]) GreaterEqual(v T) StringValidator[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, GreaterEqual(v))
	}
	return sv
}

func (sv StringValidator[T]) Between(a, b T) StringValidator[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, Between(a, b))
	}
	return sv
}

func (sv StringValidator[T]) BetweenEqual(a, b T) StringValidator[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, BetweenEqual(a, b))
	}
	return sv
}

func (sv StringValidator[T]) When(condition bool, ok StringRule[T], otherwise StringRule[T]) StringValidator[T] {
	if !sv.skip {
		if condition {
			sv.rules = append(sv.rules, ok)
		} else if otherwise != nil {
			sv.rules = append(sv.rules, otherwise)
		}
	}
	return sv
}

func (sv StringValidator[T]) With(fns ...func(s T) error) StringValidator[T] {
	if !sv.skip {
		slices.Grow(sv.rules, len(fns))
		for _, fn := range fns {
			sv.rules = append(sv.rules, StringRuleFunc[T](fn))
		}
	}
	return sv
}

func (sv StringValidator[T]) By(rules ...StringRule[T]) StringValidator[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, rules...)
	}
	return sv
}

func (sv StringValidator[T]) Valid() error {
	for _, rule := range sv.rules {
		if err := rule.Validate(sv.data.value); err != nil {
			return NewValueError(sv.data.name, err)
		}
	}
	return nil
}

func (sv StringValidator[T]) Validate(v T) error {
	for _, rule := range sv.rules {
		if err := rule.Validate(v); err != nil {
			return err
		}
	}
	return nil
}
