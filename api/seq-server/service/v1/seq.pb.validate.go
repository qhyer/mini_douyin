// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: seq-server/service/v1/seq.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on GetIDRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetIDRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetIDRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetIDRequestMultiError, or
// nil if none found.
func (m *GetIDRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetIDRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for BusinessId

	if len(errors) > 0 {
		return GetIDRequestMultiError(errors)
	}

	return nil
}

// GetIDRequestMultiError is an error wrapping multiple validation errors
// returned by GetIDRequest.ValidateAll() if the designated constraints aren't met.
type GetIDRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetIDRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetIDRequestMultiError) AllErrors() []error { return m }

// GetIDRequestValidationError is the validation error returned by
// GetIDRequest.Validate if the designated constraints aren't met.
type GetIDRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetIDRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetIDRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetIDRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetIDRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetIDRequestValidationError) ErrorName() string { return "GetIDRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetIDRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetIDRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetIDRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetIDRequestValidationError{}

// Validate checks the field values on GetIDResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetIDResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetIDResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetIDResponseMultiError, or
// nil if none found.
func (m *GetIDResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetIDResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ID

	// no validation rules for IsOk

	if len(errors) > 0 {
		return GetIDResponseMultiError(errors)
	}

	return nil
}

// GetIDResponseMultiError is an error wrapping multiple validation errors
// returned by GetIDResponse.ValidateAll() if the designated constraints
// aren't met.
type GetIDResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetIDResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetIDResponseMultiError) AllErrors() []error { return m }

// GetIDResponseValidationError is the validation error returned by
// GetIDResponse.Validate if the designated constraints aren't met.
type GetIDResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetIDResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetIDResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetIDResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetIDResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetIDResponseValidationError) ErrorName() string { return "GetIDResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetIDResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetIDResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetIDResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetIDResponseValidationError{}

// Validate checks the field values on UpdateMaxSeqRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateMaxSeqRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateMaxSeqRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateMaxSeqRequestMultiError, or nil if none found.
func (m *UpdateMaxSeqRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateMaxSeqRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for BusinessId

	// no validation rules for Seq

	// no validation rules for Step

	if len(errors) > 0 {
		return UpdateMaxSeqRequestMultiError(errors)
	}

	return nil
}

// UpdateMaxSeqRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateMaxSeqRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateMaxSeqRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateMaxSeqRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateMaxSeqRequestMultiError) AllErrors() []error { return m }

// UpdateMaxSeqRequestValidationError is the validation error returned by
// UpdateMaxSeqRequest.Validate if the designated constraints aren't met.
type UpdateMaxSeqRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateMaxSeqRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateMaxSeqRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateMaxSeqRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateMaxSeqRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateMaxSeqRequestValidationError) ErrorName() string {
	return "UpdateMaxSeqRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateMaxSeqRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateMaxSeqRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateMaxSeqRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateMaxSeqRequestValidationError{}

// Validate checks the field values on UpdateMaxSeqResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateMaxSeqResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateMaxSeqResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateMaxSeqResponseMultiError, or nil if none found.
func (m *UpdateMaxSeqResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateMaxSeqResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for IsOk

	if len(errors) > 0 {
		return UpdateMaxSeqResponseMultiError(errors)
	}

	return nil
}

// UpdateMaxSeqResponseMultiError is an error wrapping multiple validation
// errors returned by UpdateMaxSeqResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdateMaxSeqResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateMaxSeqResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateMaxSeqResponseMultiError) AllErrors() []error { return m }

// UpdateMaxSeqResponseValidationError is the validation error returned by
// UpdateMaxSeqResponse.Validate if the designated constraints aren't met.
type UpdateMaxSeqResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateMaxSeqResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateMaxSeqResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateMaxSeqResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateMaxSeqResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateMaxSeqResponseValidationError) ErrorName() string {
	return "UpdateMaxSeqResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateMaxSeqResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateMaxSeqResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateMaxSeqResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateMaxSeqResponseValidationError{}

// Validate checks the field values on Business with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Business) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Business with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BusinessMultiError, or nil
// if none found.
func (m *Business) ValidateAll() error {
	return m.validate(true)
}

func (m *Business) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for LastTime

	// no validation rules for Step

	// no validation rules for CurSeq

	// no validation rules for MaxSeq

	// no validation rules for UpdateTime

	// no validation rules for CreateTime

	if len(errors) > 0 {
		return BusinessMultiError(errors)
	}

	return nil
}

// BusinessMultiError is an error wrapping multiple validation errors returned
// by Business.ValidateAll() if the designated constraints aren't met.
type BusinessMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BusinessMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BusinessMultiError) AllErrors() []error { return m }

// BusinessValidationError is the validation error returned by
// Business.Validate if the designated constraints aren't met.
type BusinessValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BusinessValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BusinessValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BusinessValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BusinessValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BusinessValidationError) ErrorName() string { return "BusinessValidationError" }

// Error satisfies the builtin error interface
func (e BusinessValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBusiness.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BusinessValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BusinessValidationError{}
