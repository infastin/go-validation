package validation

import (
	"slices"
	"time"
)

type timeValidatorData struct {
	value time.Time
	name  string
}

type TimeValidator struct {
	data  *timeValidatorData
	rules []TimeRule
	skip  bool
}

func Time(v time.Time, name string) TimeValidator {
	return TimeValidator{
		data: &timeValidatorData{
			value: v,
			name:  name,
		},
		rules: make([]TimeRule, 0),
		skip:  false,
	}
}

func TimeV() TimeValidator {
	return TimeValidator{
		data:  nil,
		rules: make([]TimeRule, 0),
		skip:  false,
	}
}

func (sv TimeValidator) Required(condition bool) TimeValidator {
	if !sv.skip {
		sv.rules = append(sv.rules, RequiredTime(condition))
	}
	return sv
}

func (sv TimeValidator) Skip(condition bool) TimeValidator {
	if !sv.skip && condition {
		sv.skip = true
	}
	return sv
}

func (sv TimeValidator) In(elements ...time.Time) TimeValidator {
	if !sv.skip {
		sv.rules = append(sv.rules, InTime(elements...))
	}
	return sv
}

func (sv TimeValidator) NotIn(elements ...time.Time) TimeValidator {
	if !sv.skip {
		sv.rules = append(sv.rules, NotInTime(elements...))
	}
	return sv
}

func (sv TimeValidator) Equal(v time.Time) TimeValidator {
	if !sv.skip {
		sv.rules = append(sv.rules, EqualTime(v))
	}
	return sv
}

func (sv TimeValidator) Less(v time.Time) TimeValidator {
	if !sv.skip {
		sv.rules = append(sv.rules, LessTime(v))
	}
	return sv
}

func (sv TimeValidator) LessEqual(v time.Time) TimeValidator {
	if !sv.skip {
		sv.rules = append(sv.rules, LessEqualTime(v))
	}
	return sv
}

func (sv TimeValidator) Greater(v time.Time) TimeValidator {
	if !sv.skip {
		sv.rules = append(sv.rules, GreaterTime(v))
	}
	return sv
}

func (sv TimeValidator) GreaterEqual(v time.Time) TimeValidator {
	if !sv.skip {
		sv.rules = append(sv.rules, GreaterEqualTime(v))
	}
	return sv
}

func (sv TimeValidator) Between(a, b time.Time) TimeValidator {
	if !sv.skip {
		sv.rules = append(sv.rules, BetweenTime(a, b))
	}
	return sv
}

func (sv TimeValidator) BetweenEqual(a, b time.Time) TimeValidator {
	if !sv.skip {
		sv.rules = append(sv.rules, BetweenEqualTime(a, b))
	}
	return sv
}

func (sv TimeValidator) When(condition bool, ok, otherwise TimeRule) TimeValidator {
	if !sv.skip {
		if condition {
			sv.rules = append(sv.rules, ok)
		} else if otherwise != nil {
			sv.rules = append(sv.rules, otherwise)
		}
	}
	return sv
}

func (sv TimeValidator) With(fns ...func(v time.Time) error) TimeValidator {
	if !sv.skip {
		slices.Grow(sv.rules, len(fns))
		for _, fn := range fns {
			sv.rules = append(sv.rules, TimeRuleFunc(fn))
		}
	}
	return sv
}

func (sv TimeValidator) By(rules ...TimeRule) TimeValidator {
	if !sv.skip {
		sv.rules = append(sv.rules, rules...)
	}
	return sv
}

func (sv TimeValidator) Valid() error {
	for _, rule := range sv.rules {
		if err := rule.Validate(sv.data.value); err != nil {
			return NewValueError(sv.data.name, err)
		}
	}
	return nil
}

func (sv TimeValidator) Validate(v time.Time) error {
	for _, rule := range sv.rules {
		if err := rule.Validate(v); err != nil {
			return err
		}
	}
	return nil
}
