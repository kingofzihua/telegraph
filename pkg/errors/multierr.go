package errors

import (
	"go.uber.org/multierr"
)

func Combine(errors ...error) error {
	return multierr.Combine(errors...)
}

func Append(left error, right error) error {
	return multierr.Append(left, right)
}

func AppendInto(into *error, err error) (errored bool) {
	return multierr.AppendInto(into, err)
}
