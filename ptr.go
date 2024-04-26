package validation

import "slices"

type ptrValidatorData[T any] struct {
	value *T
	name  string
}

type PtrValidator[T any] struct {
	data  *ptrValidatorData[T]
	rules []PtrRule[T]
	skip  bool
}

func Ptr[T any](p *T, name string) PtrValidator[T] {
	return PtrValidator[T]{
		data: &ptrValidatorData[T]{
			value: p,
			name:  name,
		},
		rules: make([]PtrRule[T], 0),
		skip:  false,
	}
}

func PtrV[T any]() PtrValidator[T] {
	return PtrValidator[T]{
		data:  nil,
		rules: make([]PtrRule[T], 0),
		skip:  false,
	}
}

func (pv PtrValidator[T]) NotNil(condition bool) PtrValidator[T] {
	if !pv.skip {
		pv.rules = append(pv.rules, NotNilPtr[T](condition))
	}
	return pv
}

func (pv PtrValidator[T]) Nil(condition bool) PtrValidator[T] {
	if !pv.skip {
		pv.rules = append(pv.rules, NilPtr[T](condition))
	}
	return pv
}

func (pv PtrValidator[T]) Skip(condition bool) PtrValidator[T] {
	if !pv.skip {
		pv.skip = true
	}
	return pv
}

func (pv PtrValidator[T]) When(condition bool, ok, otherwise PtrRule[T]) PtrValidator[T] {
	if !pv.skip {
		if condition {
			pv.rules = append(pv.rules, ok)
		} else {
			pv.rules = append(pv.rules, otherwise)
		}
	}
	return pv
}

func (pv PtrValidator[T]) With(fns ...func(p *T) error) PtrValidator[T] {
	if !pv.skip {
		slices.Grow(pv.rules, len(fns))
		for _, fn := range fns {
			pv.rules = append(pv.rules, PtrRuleFunc[T](fn))
		}
	}
	return pv
}

func (pv PtrValidator[T]) By(rules ...PtrRule[T]) PtrValidator[T] {
	if !pv.skip {
		pv.rules = append(pv.rules, rules...)
	}
	return pv
}

func (pv PtrValidator[T]) ValueBy(rules ...AnyRule[T]) PtrValidator[T] {
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

func (pv PtrValidator[T]) ValueWith(fns ...func(p T) error) PtrValidator[T] {
	if !pv.skip {
		pv.rules = append(pv.rules, PtrRuleFunc[T](func(p *T) error {
			for _, fn := range fns {
				if err := fn(*p); err != nil {
					return err
				}
			}
			return nil
		}))
	}
	return pv
}

func (pv PtrValidator[T]) Valid() error {
	for _, rule := range pv.rules {
		if err := rule.Validate(pv.data.value); err != nil {
			return NewValueError(pv.data.name, err)
		}
	}
	return nil
}

func (pv PtrValidator[T]) Validate(v *T) error {
	for _, rule := range pv.rules {
		if err := rule.Validate(v); err != nil {
			return err
		}
	}
	return nil
}
