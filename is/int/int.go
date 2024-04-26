package isint

import (
	"github.com/infastin/go-validation"
	"github.com/infastin/go-validation/constraints"
)

var (
	ErrPort = validation.NewRuleError("is_port", "must be a valid port number")
)

func Port[T constraints.Int](i T) error {
	if i <= 0 || int32(i) >= 65535 {
		return ErrPort
	}
	return nil
}
