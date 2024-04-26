package validation

import "slices"

type sliceValidationData[T any] struct {
	value []T
	name  string
}

type SliceValidation[T any] struct {
	data  *sliceValidationData[T]
	rules []SliceRule[T]
	skip  bool
	wrap  bool
}

func Slice[T any](s []T, name string) SliceValidation[T] {
	return SliceValidation[T]{
		data: &sliceValidationData[T]{
			value: s,
			name:  name,
		},
		rules: make([]SliceRule[T], 0),
		skip:  false,
	}
}

func SliceV[T any]() SliceValidation[T] {
	return SliceValidation[T]{
		data:  nil,
		rules: make([]SliceRule[T], 0),
		skip:  false,
	}
}

func (sv SliceValidation[T]) Required(condition bool) SliceValidation[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, RequiredSlice[T](condition))
	}
	return sv
}

func (sv SliceValidation[T]) NilOrNotEmpty(condition bool) SliceValidation[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, NilOrNotEmptySlice[T](condition))
	}
	return sv
}

func (sv SliceValidation[T]) NotNil(condition bool) SliceValidation[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, NotNilSlice[T](condition))
	}
	return sv
}

func (sv SliceValidation[T]) Nil(condition bool) SliceValidation[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, NilSlice[T](condition))
	}
	return sv
}

func (sv SliceValidation[T]) Skip(condition bool) SliceValidation[T] {
	if !sv.skip && condition {
		sv.skip = true
	}
	return sv
}

func (sv SliceValidation[T]) Wrap() SliceValidation[T] {
	sv.wrap = true
	return sv
}

func (sv SliceValidation[T]) Length(min, max int) SliceValidation[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, LengthSlice[T](min, max))
	}
	return sv
}

func (sv SliceValidation[T]) When(condition bool, ok SliceRule[T], otherwise SliceRule[T]) SliceValidation[T] {
	if !sv.skip {
		if condition {
			sv.rules = append(sv.rules, ok)
		} else {
			sv.rules = append(sv.rules, otherwise)
		}
	}
	return sv
}

func (sv SliceValidation[T]) With(fns ...func(s []T) error) SliceValidation[T] {
	if !sv.skip {
		slices.Grow(sv.rules, len(fns))
		for _, fn := range fns {
			sv.rules = append(sv.rules, SliceRuleFunc[T](fn))
		}
	}
	return sv
}

func (sv SliceValidation[T]) By(rules ...SliceRule[T]) SliceValidation[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, rules...)
	}
	return sv
}

func (sv SliceValidation[T]) Dive(rules ...Rule[T]) SliceValidation[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, SliceRuleFunc[T](func(s []T) error {
			for i := range s {
				for _, rule := range rules {
					if err := rule.Validate(s[i]); err != nil {
						if sv.wrap {
							err = NewWrapError(err)
						}
						return NewIndexError(i, err)
					}
				}
			}
			return nil
		}))
	}
	return sv
}

func (sv SliceValidation[T]) DivePtr(rules ...Rule[*T]) SliceValidation[T] {
	if !sv.skip {
		sv.rules = append(sv.rules, SliceRuleFunc[T](func(s []T) error {
			for i := range s {
				for _, rule := range rules {
					if err := rule.Validate(&s[i]); err != nil {
						if sv.wrap {
							err = NewWrapError(err)
						}
						return NewIndexError(i, err)
					}
				}
			}
			return nil
		}))
	}
	return sv
}

func (sv SliceValidation[T]) Valid() error {
	for _, rule := range sv.rules {
		if err := rule.Validate(sv.data.value); err != nil {
			return NewValueError(sv.data.name, err)
		}
	}
	return nil
}

func (sv SliceValidation[T]) Validate(v []T) error {
	for _, rule := range sv.rules {
		if err := rule.Validate(v); err != nil {
			return err
		}
	}
	return nil
}
