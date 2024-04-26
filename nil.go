package validation

type nilPtrRule[T any] struct {
	condition bool
}

var ErrNil = NewRuleError("nil", "must be blank")

func NilPtr[T any](condition bool) nilPtrRule[T] {
	return nilPtrRule[T]{
		condition: condition,
	}
}

func (r nilPtrRule[T]) Validate(p *T) error {
	if r.condition && p != nil {
		return ErrNil
	}
	return nil
}

type nilSliceRule[T any] struct {
	condition bool
}

func NilSlice[T any](condition bool) nilSliceRule[T] {
	return nilSliceRule[T]{
		condition: condition,
	}
}

func (r nilSliceRule[T]) Validate(s []T) error {
	if r.condition && s != nil {
		return ErrNil
	}
	return nil
}

type nilMapRule[T any] struct {
	condition bool
}

func NilMap[T any](condition bool) nilMapRule[T] {
	return nilMapRule[T]{
		condition: condition,
	}
}

func (r nilMapRule[T]) Validate(m map[string]T) error {
	if r.condition && m != nil {
		return ErrNil
	}
	return nil
}
