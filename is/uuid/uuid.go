package isuuid

import (
	"github.com/infastin/go-validation"
	"github.com/infastin/go-validation/constraints"
)

var (
	ErrUUIDv1 = validation.NewRuleError("is_uuid_v1", "must be a valid UUID v1")
	ErrUUIDv3 = validation.NewRuleError("is_uuid_v3", "must be a valid UUID v3")
	ErrUUIDv4 = validation.NewRuleError("is_uuid_v4", "must be a valid UUID v4")
	ErrUUIDv5 = validation.NewRuleError("is_uuid_v5", "must be a valid UUID v5")
)

func V1[T constraints.UUID](v T) error {
	if v[6]>>4 != 1 {
		return ErrUUIDv1
	}
	return nil
}

func V3[T constraints.UUID](v T) error {
	if v[6]>>4 != 3 {
		return ErrUUIDv3
	}
	return nil
}

func V4[T constraints.UUID](v T) error {
	if v[6]>>4 != 4 {
		return ErrUUIDv4
	}
	return nil
}

func V5[T constraints.UUID](v T) error {
	if v[6]>>4 != 5 {
		return ErrUUIDv5
	}
	return nil
}
