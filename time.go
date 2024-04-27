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

func (tv TimeValidator) Required(condition bool) TimeValidator {
	if !tv.skip {
		tv.rules = append(tv.rules, RequiredTime(condition))
	}
	return tv
}

func (tv TimeValidator) Skip(condition bool) TimeValidator {
	if !tv.skip && condition {
		tv.skip = true
	}
	return tv
}

func (tv TimeValidator) In(elements ...time.Time) TimeValidator {
	if !tv.skip {
		tv.rules = append(tv.rules, InTime(elements...))
	}
	return tv
}

func (tv TimeValidator) NotIn(elements ...time.Time) TimeValidator {
	if !tv.skip {
		tv.rules = append(tv.rules, NotInTime(elements...))
	}
	return tv
}

func (tv TimeValidator) Equal(v time.Time) TimeValidator {
	if !tv.skip {
		tv.rules = append(tv.rules, EqualTime(v))
	}
	return tv
}

func (tv TimeValidator) Less(v time.Time) TimeValidator {
	if !tv.skip {
		tv.rules = append(tv.rules, LessTime(v))
	}
	return tv
}

func (tv TimeValidator) LessEqual(v time.Time) TimeValidator {
	if !tv.skip {
		tv.rules = append(tv.rules, LessEqualTime(v))
	}
	return tv
}

func (tv TimeValidator) Greater(v time.Time) TimeValidator {
	if !tv.skip {
		tv.rules = append(tv.rules, GreaterTime(v))
	}
	return tv
}

func (tv TimeValidator) GreaterEqual(v time.Time) TimeValidator {
	if !tv.skip {
		tv.rules = append(tv.rules, GreaterEqualTime(v))
	}
	return tv
}

func (tv TimeValidator) Between(a, b time.Time) TimeValidator {
	if !tv.skip {
		tv.rules = append(tv.rules, BetweenTime(a, b))
	}
	return tv
}

func (tv TimeValidator) BetweenEqual(a, b time.Time) TimeValidator {
	if !tv.skip {
		tv.rules = append(tv.rules, BetweenEqualTime(a, b))
	}
	return tv
}

func (tv TimeValidator) When(condition bool, ok, otherwise TimeRule) TimeValidator {
	if !tv.skip {
		if condition {
			tv.rules = append(tv.rules, ok)
		} else if otherwise != nil {
			tv.rules = append(tv.rules, otherwise)
		}
	}
	return tv
}

func (tv TimeValidator) With(fns ...func(v time.Time) error) TimeValidator {
	if !tv.skip {
		slices.Grow(tv.rules, len(fns))
		for _, fn := range fns {
			tv.rules = append(tv.rules, TimeRuleFunc(fn))
		}
	}
	return tv
}

func (tv TimeValidator) By(rules ...TimeRule) TimeValidator {
	if !tv.skip {
		tv.rules = append(tv.rules, rules...)
	}
	return tv
}

func (tv TimeValidator) Valid() error {
	for _, rule := range tv.rules {
		if err := rule.Validate(tv.data.value); err != nil {
			return NewValueError(tv.data.name, err)
		}
	}
	return nil
}

func (tv TimeValidator) Validate(v time.Time) error {
	for _, rule := range tv.rules {
		if err := rule.Validate(v); err != nil {
			return err
		}
	}
	return nil
}
