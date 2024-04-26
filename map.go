package validation

import "slices"

type mapValidationData[T any] struct {
	value map[string]T
	name  string
}

type MapValidation[T any] struct {
	data  *mapValidationData[T]
	rules []MapRule[T]
	skip  bool
}

func Map[T any](m map[string]T, name string) MapValidation[T] {
	return MapValidation[T]{
		data: &mapValidationData[T]{
			value: m,
			name:  name,
		},
		rules: make([]MapRule[T], 0),
		skip:  false,
	}
}

func MapV[T any]() MapValidation[T] {
	return MapValidation[T]{
		data:  nil,
		rules: make([]MapRule[T], 0),
		skip:  false,
	}
}

func (mv MapValidation[T]) Required(condition bool) MapValidation[T] {
	if !mv.skip {
		mv.rules = append(mv.rules, RequiredMap[T](condition))
	}
	return mv
}

func (mv MapValidation[T]) NilOrNotEmpty(condition bool) MapValidation[T] {
	if !mv.skip {
		mv.rules = append(mv.rules, NilOrNotEmptyMap[T](condition))
	}
	return mv
}

func (mv MapValidation[T]) Empty(condition bool) MapValidation[T] {
	if !mv.skip {
		mv.rules = append(mv.rules, EmptyMap[T](condition))
	}
	return mv
}

func (mv MapValidation[T]) NotNil(condition bool) MapValidation[T] {
	if !mv.skip {
		mv.rules = append(mv.rules, NotNilMap[T](condition))
	}
	return mv
}

func (mv MapValidation[T]) Nil(condition bool) MapValidation[T] {
	if !mv.skip {
		mv.rules = append(mv.rules, NilMap[T](condition))
	}
	return mv
}

func (mv MapValidation[T]) Skip(condition bool) MapValidation[T] {
	if !mv.skip && condition {
		mv.skip = true
	}
	return mv
}

func (mv MapValidation[T]) Length(min, max int) MapValidation[T] {
	if !mv.skip {
		mv.rules = append(mv.rules, LengthMap[T](min, max))
	}
	return mv
}

func (mv MapValidation[T]) When(condition bool, ok MapRule[T], otherwise MapRule[T]) MapValidation[T] {
	if !mv.skip {
		if condition {
			mv.rules = append(mv.rules, ok)
		} else {
			mv.rules = append(mv.rules, otherwise)
		}
	}
	return mv
}

func (mv MapValidation[T]) With(fns ...func(s map[string]T) error) MapValidation[T] {
	if !mv.skip {
		slices.Grow(mv.rules, len(fns))
		for _, fn := range fns {
			mv.rules = append(mv.rules, MapRuleFunc[T](fn))
		}
	}
	return mv
}

func (mv MapValidation[T]) By(rules ...MapRule[T]) MapValidation[T] {
	if !mv.skip {
		mv.rules = append(mv.rules, rules...)
	}
	return mv
}

func (mv MapValidation[T]) Valid() error {
	for _, rule := range mv.rules {
		if err := rule.Validate(mv.data.value); err != nil {
			return NewValueError(mv.data.name, err)
		}
	}
	return nil
}

func (mv MapValidation[T]) Validate(m map[string]T) error {
	for _, rule := range mv.rules {
		if err := rule.Validate(m); err != nil {
			return err
		}
	}
	return nil
}
