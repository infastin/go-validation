package validation

import (
	"fmt"
	"time"

	"github.com/infastin/go-validation/constraints"
)

type compareRule[T constraints.Ordered] struct {
	comp       func(x T) bool
	buildError func() error
}

func Equal[T constraints.Ordered](v T) compareRule[T] {
	return compareRule[T]{
		comp: func(x T) bool {
			return x == v
		},
		buildError: func() func() error {
			var err error
			return func() error {
				if err == nil {
					err = buildEqualError(v)
				}
				return err
			}
		}(),
	}
}

func Less[T constraints.Ordered](v T) compareRule[T] {
	return compareRule[T]{
		comp: func(x T) bool {
			return x < v
		},
		buildError: func() func() error {
			var err error
			return func() error {
				if err == nil {
					err = buildLessError(v)
				}
				return err
			}
		}(),
	}
}

func LessEqual[T constraints.Ordered](v T) compareRule[T] {
	return compareRule[T]{
		comp: func(x T) bool {
			return x <= v
		},
		buildError: func() func() error {
			var err error
			return func() error {
				if err == nil {
					err = buildLessEqualError(v)
				}
				return err
			}
		}(),
	}
}

func Greater[T constraints.Ordered](v T) compareRule[T] {
	return compareRule[T]{
		comp: func(x T) bool {
			return x > v
		},
		buildError: func() func() error {
			var err error
			return func() error {
				if err == nil {
					err = buildGreaterError(v)
				}
				return err
			}
		}(),
	}
}

func GreaterEqual[T constraints.Ordered](v T) compareRule[T] {
	return compareRule[T]{
		comp: func(x T) bool {
			return x > v
		},
		buildError: func() func() error {
			var err error
			return func() error {
				if err == nil {
					err = buildGreaterEqualError(v)
				}
				return err
			}
		}(),
	}
}

func Between[T constraints.Ordered](a, b T) compareRule[T] {
	return compareRule[T]{
		comp: func(x T) bool {
			return x > a && x < b
		},
		buildError: func() func() error {
			var err error
			return func() error {
				if err == nil {
					err = buildBetweenError(a, b)
				}
				return err
			}
		}(),
	}
}

func BetweenEqual[T constraints.Ordered](a, b T) compareRule[T] {
	return compareRule[T]{
		comp: func(x T) bool {
			return x >= a && x <= b
		},
		buildError: func() func() error {
			var err error
			return func() error {
				if err == nil {
					err = buildBetweenEqualError(a, b)
				}
				return err
			}
		}(),
	}
}

func (r compareRule[T]) Validate(v T) error {
	if !r.comp(v) {
		return r.buildError()
	}
	return nil
}

type compareTimeRule struct {
	comp       func(x time.Time) bool
	buildError func() error
}

func EqualTime(v time.Time) compareTimeRule {
	return compareTimeRule{
		comp: func(x time.Time) bool {
			return x.Equal(v)
		},
		buildError: func() func() error {
			var err error
			return func() error {
				if err == nil {
					err = buildEqualError(v)
				}
				return err
			}
		}(),
	}
}

func LessTime(v time.Time) compareTimeRule {
	return compareTimeRule{
		comp: func(x time.Time) bool {
			return x.Compare(v) < 0
		},
		buildError: func() func() error {
			var err error
			return func() error {
				if err == nil {
					err = buildLessError(v)
				}
				return err
			}
		}(),
	}
}

func LessEqualTime(v time.Time) compareTimeRule {
	return compareTimeRule{
		comp: func(x time.Time) bool {
			return x.Compare(v) <= 0
		},
		buildError: func() func() error {
			var err error
			return func() error {
				if err == nil {
					err = buildLessEqualError(v)
				}
				return err
			}
		}(),
	}
}

func GreaterTime(v time.Time) compareTimeRule {
	return compareTimeRule{
		comp: func(x time.Time) bool {
			return x.Compare(v) > 0
		},
		buildError: func() func() error {
			var err error
			return func() error {
				if err == nil {
					err = buildGreaterError(v)
				}
				return err
			}
		}(),
	}
}

func GreaterEqualTime(v time.Time) compareTimeRule {
	return compareTimeRule{
		comp: func(x time.Time) bool {
			return x.Compare(v) >= 0
		},
		buildError: func() func() error {
			var err error
			return func() error {
				if err == nil {
					err = buildGreaterEqualError(v)
				}
				return err
			}
		}(),
	}
}

func BetweenTime(a, b time.Time) compareTimeRule {
	return compareTimeRule{
		comp: func(x time.Time) bool {
			return x.After(a) && x.Before(b)
		},
		buildError: func() func() error {
			var err error
			return func() error {
				if err == nil {
					err = buildBetweenError(a, b)
				}
				return err
			}
		}(),
	}
}

func BetweenEqualTime(a, b time.Time) compareTimeRule {
	return compareTimeRule{
		comp: func(x time.Time) bool {
			return x.Compare(a) >= 0 && x.Compare(b) <= 0
		},
		buildError: func() func() error {
			var err error
			return func() error {
				if err == nil {
					err = buildBetweenEqualError(a, b)
				}
				return err
			}
		}(),
	}
}

func (r compareTimeRule) Validate(v time.Time) error {
	if !r.comp(v) {
		return r.buildError()
	}
	return nil
}

func buildEqualError(v any) error {
	return NewRuleError("equal", fmt.Sprintf("must be equal to %v", v))
}

func buildLessError(v any) error {
	return NewRuleError("less", fmt.Sprintf("must be less than %v", v))
}

func buildLessEqualError(v any) error {
	return NewRuleError("less_equal", fmt.Sprintf("must be no greater than %v", v))
}

func buildGreaterError(v any) error {
	return NewRuleError("greater", fmt.Sprintf("must be greater than %v", v))
}

func buildGreaterEqualError(v any) error {
	return NewRuleError("greater_equal", fmt.Sprintf("must be no less than %v", v))
}

func buildBetweenError(a, b any) error {
	return NewRuleError("between", fmt.Sprintf("must exclusively be between %v and %v", a, b))
}

func buildBetweenEqualError(a, b any) error {
	return NewRuleError("between_equal", fmt.Sprintf("must inclusively be between %v and %v", a, b))
}