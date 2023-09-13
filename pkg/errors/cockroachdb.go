package errors

import (
	cerr "github.com/cockroachdb/errors"
)

// New creates an error with a simple error message.
// A stack trace is retained.
//
// Note: the message string is assumed to not contain
// PII and is included in Sentry reports.
// Use errors.Newf("%s", <unsafestring>) for errors
// strings that may contain PII information.
//
// Detail output:
// - message via `Error()` and formatting using `%v`/`%s`/`%q`.
// - everything when formatting with `%+v`.
// - stack trace and message via `errors.GetSafeDetails()`.
// - stack trace and message in Sentry reports.
func New(msg string) error { return cerr.New(msg) }

// Newf creates an error with a formatted error message.
// A stack trace is retained.
//
// Note: the format string is assumed to not contain
// PII and is included in Sentry reports.
// Use errors.Newf("%s", <unsafestring>) for errors
// strings that may contain PII information.
//
// See the doc of `New()` for more details.
func Newf(format string, args ...interface{}) error { return cerr.Newf(format, args...) }

// Errorf aliases Newf().
func Errorf(format string, args ...interface{}) error { return cerr.Errorf(format, args...) }

// Cause aliases UnwrapAll() for compatibility with github.com/pkg/errors.
func Cause(err error) error { return cerr.UnwrapAll(err) }

// Unwrap aliases UnwrapOnce() for compatibility with xerrors.
func Unwrap(err error) error { return cerr.UnwrapOnce(err) }

// Wrap wraps an error with a message prefix.
// A stack trace is retained.
//
// Note: the prefix string is assumed to not contain
// PII and is included in Sentry reports.
// Use errors.Wrapf(err, "%s", <unsafestring>) for errors
// strings that may contain PII information.
//
// Detail output:
// - original error message + prefix via `Error()` and formatting using `%v`/`%s`/`%q`.
// - everything when formatting with `%+v`.
// - stack trace and message via `errors.GetSafeDetails()`.
// - stack trace and message in Sentry reports.
func Wrap(err error, msg string) error { return cerr.Wrap(err, msg) }

// WithMessage annotates err with a new message.
// If err is nil, WithMessage returns nil.
// The message is considered safe for reporting
// and is included in Sentry reports.
func WithMessage(err error, msg string) error { return cerr.WithMessage(err, msg) }

// WithMessagef annotates err with the format specifier.
// If err is nil, WithMessagef returns nil.
// The message is formatted as per redact.Sprintf,
// to separate safe and unsafe strings for Sentry reporting.
func WithMessagef(err error, format string, args ...interface{}) error {
	return cerr.WithMessagef(err, format, args...)
}

// As finds the first error in err's chain that matches the type to which target
// points, and if so, sets the target to its value and returns true. An error
// matches a type if it is assignable to the target type, or if it has a method
// As(interface{}) bool such that As(target) returns true. As will panic if target
// is not a non-nil pointer to a type which implements error or is of interface type.
//
// The As method should set the target to its value and return true if err
// matches the type to which target points.
//
// Note: this implementation differs from that of xerrors as follows:
// - it also supports recursing through causes with Cause().
// - if it detects an API use error, its panic object is a valid error.
func As(err error, target interface{}) bool { return cerr.As(err, target) }

// Is determines whether one of the causes of the given error or any
// of its causes is equivalent to some reference error.
//
// As in the Go standard library, an error is considered to match a
// reference error if it is equal to that target or if it implements a
// method Is(error) bool such that Is(reference) returns true.
//
// Note: the inverse is not true - making an Is(reference) method
// return false does not imply that errors.Is() also returns
// false. Errors can be equal because their network equality marker is
// the same. To force errors to appear different to Is(), use
// errors.Mark().
//
// Note: if any of the error types has been migrated from a previous
// package location or a different type, ensure that
// RegisterTypeMigration() was called prior to Is().
func Is(err, reference error) bool { return cerr.Is(err, reference) }
