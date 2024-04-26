package validation

import (
	"slices"
	"time"
)

type timeValidationData struct {
	value time.Time
	name  string
}

type TimeValidation struct {
	data  *timeValidationData
	rules []TimeRule
	skip  bool
}

func Time(v time.Time, name string) TimeValidation {
	return TimeValidation{
		data: &timeValidationData{
			value: v,
			name:  name,
		},
		rules: make([]TimeRule, 0),
		skip:  false,
	}
}

func TimeV() TimeValidation {
	return TimeValidation{
		data:  nil,
		rules: make([]TimeRule, 0),
		skip:  false,
	}
}

func (sv TimeValidation) Required(condition bool) TimeValidation {
	if !sv.skip {
		sv.rules = append(sv.rules, RequiredTime(condition))
	}
	return sv
}

func (sv TimeValidation) Skip(condition bool) TimeValidation {
	if !sv.skip && condition {
		sv.skip = true
	}
	return sv
}

func (sv TimeValidation) In(elements ...time.Time) TimeValidation {
	if !sv.skip {
		sv.rules = append(sv.rules, InTime(elements...))
	}
	return sv
}

func (sv TimeValidation) NotIn(elements ...time.Time) TimeValidation {
	if !sv.skip {
		sv.rules = append(sv.rules, NotInTime(elements...))
	}
	return sv
}

func (sv TimeValidation) Equal(v time.Time) TimeValidation {
	if !sv.skip {
		sv.rules = append(sv.rules, EqualTime(v))
	}
	return sv
}

func (sv TimeValidation) Less(v time.Time) TimeValidation {
	if !sv.skip {
		sv.rules = append(sv.rules, LessTime(v))
	}
	return sv
}

func (sv TimeValidation) LessEqual(v time.Time) TimeValidation {
	if !sv.skip {
		sv.rules = append(sv.rules, LessEqualTime(v))
	}
	return sv
}

func (sv TimeValidation) Greater(v time.Time) TimeValidation {
	if !sv.skip {
		sv.rules = append(sv.rules, GreaterTime(v))
	}
	return sv
}

func (sv TimeValidation) GreaterEqual(v time.Time) TimeValidation {
	if !sv.skip {
		sv.rules = append(sv.rules, GreaterEqualTime(v))
	}
	return sv
}

func (sv TimeValidation) Between(a, b time.Time) TimeValidation {
	if !sv.skip {
		sv.rules = append(sv.rules, BetweenTime(a, b))
	}
	return sv
}

func (sv TimeValidation) BetweenEqual(a, b time.Time) TimeValidation {
	if !sv.skip {
		sv.rules = append(sv.rules, BetweenEqualTime(a, b))
	}
	return sv
}

func (sv TimeValidation) When(condition bool, ok TimeRule, otherwise TimeRule) TimeValidation {
	if !sv.skip {
		if condition {
			sv.rules = append(sv.rules, ok)
		} else if otherwise != nil {
			sv.rules = append(sv.rules, otherwise)
		}
	}
	return sv
}

func (sv TimeValidation) With(fns ...func(v time.Time) error) TimeValidation {
	if !sv.skip {
		slices.Grow(sv.rules, len(fns))
		for _, fn := range fns {
			sv.rules = append(sv.rules, TimeRuleFunc(fn))
		}
	}
	return sv
}

func (sv TimeValidation) By(rules ...TimeRule) TimeValidation {
	if !sv.skip {
		sv.rules = append(sv.rules, rules...)
	}
	return sv
}

func (sv TimeValidation) Valid() error {
	for _, rule := range sv.rules {
		if err := rule.Validate(sv.data.value); err != nil {
			return NewValueError(sv.data.name, err)
		}
	}
	return nil
}

func (sv TimeValidation) Validate(v time.Time) error {
	for _, rule := range sv.rules {
		if err := rule.Validate(v); err != nil {
			return err
		}
	}
	return nil
}
