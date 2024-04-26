package validation

type Validator interface {
	Valid() error
}

func All(validators ...Validator) error {
	var errs Errors
	for _, v := range validators {
		if err := v.Valid(); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) != 0 {
		return errs
	}
	return nil
}
