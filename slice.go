package validation

import "slices"

type sliceValidatorData[T any] struct {
	value []T
	name  string
}

type SliceValidator[T any] struct {
	data  *sliceValidatorData[T]
	rules []SliceRule[T]
	skip  bool
}

func Slice[T any](s []T, name string) SliceValidator[T] {
	return SliceValidator[T]{
		data: &sliceValidatorData[T]{
			value: s,
			name:  name,
		},
		rules: make([]SliceRule[T], 0),
		skip:  false,
	}
}

func SliceV[T any]() SliceValidator[T] {
	return SliceValidator[T]{
		data:  nil,
		rules: make([]SliceRule[T], 0),
		skip:  false,
	}
}

func (sv SliceValidator[T]) Required(condition bool) SliceValidator[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, RequiredSlice[T](condition))
	}
	return sv
}

func (sv SliceValidator[T]) NilOrNotEmpty(condition bool) SliceValidator[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, NilOrNotEmptySlice[T](condition))
	}
	return sv
}

func (sv SliceValidator[T]) Empty(condition bool) SliceValidator[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, EmptySlice[T](condition))
	}
	return sv
}

func (sv SliceValidator[T]) NotNil(condition bool) SliceValidator[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, NotNilSlice[T](condition))
	}
	return sv
}

func (sv SliceValidator[T]) Nil(condition bool) SliceValidator[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, NilSlice[T](condition))
	}
	return sv
}

func (sv SliceValidator[T]) Skip(condition bool) SliceValidator[T] {
	if !sv.skip && condition {
		sv.skip = true
	}
	return sv
}

func (sv SliceValidator[T]) Length(min, max int) SliceValidator[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, LengthSlice[T](min, max))
	}
	return sv
}

func (sv SliceValidator[T]) When(condition bool, ok, otherwise SliceRule[T]) SliceValidator[T] {
	if !sv.skip {
		if condition {
			sv.rules = append(sv.rules, ok)
		} else {
			sv.rules = append(sv.rules, otherwise)
		}
	}
	return sv
}

func (sv SliceValidator[T]) With(fns ...func(s []T) error) SliceValidator[T] {
	if !sv.skip {
		slices.Grow(sv.rules, len(fns))
		for _, fn := range fns {
			sv.rules = append(sv.rules, SliceRuleFunc[T](fn))
		}
	}
	return sv
}

func (sv SliceValidator[T]) By(rules ...SliceRule[T]) SliceValidator[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, rules...)
	}
	return sv
}

func (sv SliceValidator[T]) ValuesBy(rules ...AnyRule[T]) SliceValidator[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, SliceRuleFunc[T](func(s []T) error {
			for i := range s {
				for _, rule := range rules {
					if err := rule.Validate(s[i]); err != nil {
						return NewIndexError(i, err)
					}
				}
			}
			return nil
		}))
	}
	return sv
}

func (sv SliceValidator[T]) ValuesWith(fns ...func(v T) error) SliceValidator[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, SliceRuleFunc[T](func(s []T) error {
			for i := range s {
				for _, fn := range fns {
					if err := fn(s[i]); err != nil {
						return NewIndexError(i, err)
					}
				}
			}
			return nil
		}))
	}
	return sv
}

func (sv SliceValidator[T]) ValuesPtrBy(rules ...AnyRule[*T]) SliceValidator[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, SliceRuleFunc[T](func(s []T) error {
			for i := range s {
				for _, rule := range rules {
					if err := rule.Validate(&s[i]); err != nil {
						return NewIndexError(i, err)
					}
				}
			}
			return nil
		}))
	}
	return sv
}

func (sv SliceValidator[T]) ValuesPtrWith(fns ...func(v *T) error) SliceValidator[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, SliceRuleFunc[T](func(s []T) error {
			for i := range s {
				for _, fn := range fns {
					if err := fn(&s[i]); err != nil {
						return NewIndexError(i, err)
					}
				}
			}
			return nil
		}))
	}
	return sv
}

func (sv SliceValidator[T]) Valid() error {
	for _, rule := range sv.rules {
		if err := rule.Validate(sv.data.value); err != nil {
			return NewValueError(sv.data.name, err)
		}
	}
	return nil
}

func (sv SliceValidator[T]) Validate(v []T) error {
	for _, rule := range sv.rules {
		if err := rule.Validate(v); err != nil {
			return err
		}
	}
	return nil
}
