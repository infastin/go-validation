package validation

import "slices"

type ptrValidationData[T any] struct {
	value *T
	name  string
}

type PtrValidation[T any] struct {
	data  *ptrValidationData[T]
	rules []PtrRule[T]
	skip  bool
}

func Ptr[T any](p *T, name string) PtrValidation[T] {
	return PtrValidation[T]{
		data: &ptrValidationData[T]{
			value: p,
			name:  name,
		},
		rules: make([]PtrRule[T], 0),
		skip:  false,
	}
}

func PtrV[T any]() PtrValidation[T] {
	return PtrValidation[T]{
		data:  nil,
		rules: make([]PtrRule[T], 0),
		skip:  false,
	}
}

func (pv PtrValidation[T]) NotNil(condition bool) PtrValidation[T] {
	if !pv.skip {
		pv.rules = append(pv.rules, NotNilPtr[T](condition))
	}
	return pv
}

func (pv PtrValidation[T]) Nil(condition bool) PtrValidation[T] {
	if !pv.skip {
		pv.rules = append(pv.rules, NilPtr[T](condition))
	}
	return pv
}

func (pv PtrValidation[T]) Skip(condition bool) PtrValidation[T] {
	if !pv.skip {
		pv.skip = true
	}
	return pv
}

func (pv PtrValidation[T]) When(condition bool, ok PtrRule[T], otherwise PtrRule[T]) PtrValidation[T] {
	if !pv.skip {
		if condition {
			pv.rules = append(pv.rules, ok)
		} else {
			pv.rules = append(pv.rules, otherwise)
		}
	}
	return pv
}

func (pv PtrValidation[T]) With(fns ...func(p *T) error) PtrValidation[T] {
	if !pv.skip {
		slices.Grow(pv.rules, len(fns))
		for _, fn := range fns {
			pv.rules = append(pv.rules, PtrRuleFunc[T](fn))
		}
	}
	return pv
}

func (pv PtrValidation[T]) By(rules ...PtrRule[T]) PtrValidation[T] {
	if !pv.skip {
		pv.rules = append(pv.rules, rules...)
	}
	return pv
}

func (pv PtrValidation[T]) Dive(rules ...Rule[T]) PtrValidation[T] {
	if !pv.skip {
		pv.rules = append(pv.rules, PtrRuleFunc[T](func(p *T) error {
			for _, rule := range rules {
				if err := rule.Validate(*p); err != nil {
					return err
				}
			}
			return nil
		}))
	}
	return pv
}

func (pv PtrValidation[T]) Valid() error {
	for _, rule := range pv.rules {
		if err := rule.Validate(pv.data.value); err != nil {
			return NewValueError(pv.data.name, err)
		}
	}
	return nil
}

func (pv PtrValidation[T]) Validate(v *T) error {
	for _, rule := range pv.rules {
		if err := rule.Validate(v); err != nil {
			return err
		}
	}
	return nil
}
