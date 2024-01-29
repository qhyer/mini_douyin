// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user/passport/service/v1/passport.proto

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

// Validate checks the field values on DouyinUserRegisterRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DouyinUserRegisterRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DouyinUserRegisterRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DouyinUserRegisterRequestMultiError, or nil if none found.
func (m *DouyinUserRegisterRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DouyinUserRegisterRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Username

	// no validation rules for Password

	if len(errors) > 0 {
		return DouyinUserRegisterRequestMultiError(errors)
	}

	return nil
}

// DouyinUserRegisterRequestMultiError is an error wrapping multiple validation
// errors returned by DouyinUserRegisterRequest.ValidateAll() if the
// designated constraints aren't met.
type DouyinUserRegisterRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DouyinUserRegisterRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DouyinUserRegisterRequestMultiError) AllErrors() []error { return m }

// DouyinUserRegisterRequestValidationError is the validation error returned by
// DouyinUserRegisterRequest.Validate if the designated constraints aren't met.
type DouyinUserRegisterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DouyinUserRegisterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DouyinUserRegisterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DouyinUserRegisterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DouyinUserRegisterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DouyinUserRegisterRequestValidationError) ErrorName() string {
	return "DouyinUserRegisterRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DouyinUserRegisterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDouyinUserRegisterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DouyinUserRegisterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DouyinUserRegisterRequestValidationError{}

// Validate checks the field values on DouyinUserRegisterResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DouyinUserRegisterResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DouyinUserRegisterResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DouyinUserRegisterResponseMultiError, or nil if none found.
func (m *DouyinUserRegisterResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DouyinUserRegisterResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for StatusCode

	// no validation rules for StatusMsg

	// no validation rules for UserId

	// no validation rules for Token

	if len(errors) > 0 {
		return DouyinUserRegisterResponseMultiError(errors)
	}

	return nil
}

// DouyinUserRegisterResponseMultiError is an error wrapping multiple
// validation errors returned by DouyinUserRegisterResponse.ValidateAll() if
// the designated constraints aren't met.
type DouyinUserRegisterResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DouyinUserRegisterResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DouyinUserRegisterResponseMultiError) AllErrors() []error { return m }

// DouyinUserRegisterResponseValidationError is the validation error returned
// by DouyinUserRegisterResponse.Validate if the designated constraints aren't met.
type DouyinUserRegisterResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DouyinUserRegisterResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DouyinUserRegisterResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DouyinUserRegisterResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DouyinUserRegisterResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DouyinUserRegisterResponseValidationError) ErrorName() string {
	return "DouyinUserRegisterResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DouyinUserRegisterResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDouyinUserRegisterResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DouyinUserRegisterResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DouyinUserRegisterResponseValidationError{}

// Validate checks the field values on DouyinUserLoginRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DouyinUserLoginRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DouyinUserLoginRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DouyinUserLoginRequestMultiError, or nil if none found.
func (m *DouyinUserLoginRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DouyinUserLoginRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Username

	// no validation rules for Password

	if len(errors) > 0 {
		return DouyinUserLoginRequestMultiError(errors)
	}

	return nil
}

// DouyinUserLoginRequestMultiError is an error wrapping multiple validation
// errors returned by DouyinUserLoginRequest.ValidateAll() if the designated
// constraints aren't met.
type DouyinUserLoginRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DouyinUserLoginRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DouyinUserLoginRequestMultiError) AllErrors() []error { return m }

// DouyinUserLoginRequestValidationError is the validation error returned by
// DouyinUserLoginRequest.Validate if the designated constraints aren't met.
type DouyinUserLoginRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DouyinUserLoginRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DouyinUserLoginRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DouyinUserLoginRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DouyinUserLoginRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DouyinUserLoginRequestValidationError) ErrorName() string {
	return "DouyinUserLoginRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DouyinUserLoginRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDouyinUserLoginRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DouyinUserLoginRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DouyinUserLoginRequestValidationError{}

// Validate checks the field values on DouyinUserLoginResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DouyinUserLoginResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DouyinUserLoginResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DouyinUserLoginResponseMultiError, or nil if none found.
func (m *DouyinUserLoginResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DouyinUserLoginResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for StatusCode

	// no validation rules for UserId

	// no validation rules for Token

	if m.StatusMsg != nil {
		// no validation rules for StatusMsg
	}

	if len(errors) > 0 {
		return DouyinUserLoginResponseMultiError(errors)
	}

	return nil
}

// DouyinUserLoginResponseMultiError is an error wrapping multiple validation
// errors returned by DouyinUserLoginResponse.ValidateAll() if the designated
// constraints aren't met.
type DouyinUserLoginResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DouyinUserLoginResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DouyinUserLoginResponseMultiError) AllErrors() []error { return m }

// DouyinUserLoginResponseValidationError is the validation error returned by
// DouyinUserLoginResponse.Validate if the designated constraints aren't met.
type DouyinUserLoginResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DouyinUserLoginResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DouyinUserLoginResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DouyinUserLoginResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DouyinUserLoginResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DouyinUserLoginResponseValidationError) ErrorName() string {
	return "DouyinUserLoginResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DouyinUserLoginResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDouyinUserLoginResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DouyinUserLoginResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DouyinUserLoginResponseValidationError{}

// Validate checks the field values on DouyinGetUserInfoRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DouyinGetUserInfoRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DouyinGetUserInfoRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DouyinGetUserInfoRequestMultiError, or nil if none found.
func (m *DouyinGetUserInfoRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DouyinGetUserInfoRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	if len(errors) > 0 {
		return DouyinGetUserInfoRequestMultiError(errors)
	}

	return nil
}

// DouyinGetUserInfoRequestMultiError is an error wrapping multiple validation
// errors returned by DouyinGetUserInfoRequest.ValidateAll() if the designated
// constraints aren't met.
type DouyinGetUserInfoRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DouyinGetUserInfoRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DouyinGetUserInfoRequestMultiError) AllErrors() []error { return m }

// DouyinGetUserInfoRequestValidationError is the validation error returned by
// DouyinGetUserInfoRequest.Validate if the designated constraints aren't met.
type DouyinGetUserInfoRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DouyinGetUserInfoRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DouyinGetUserInfoRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DouyinGetUserInfoRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DouyinGetUserInfoRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DouyinGetUserInfoRequestValidationError) ErrorName() string {
	return "DouyinGetUserInfoRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DouyinGetUserInfoRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDouyinGetUserInfoRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DouyinGetUserInfoRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DouyinGetUserInfoRequestValidationError{}

// Validate checks the field values on DouyinGetUserInfoResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DouyinGetUserInfoResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DouyinGetUserInfoResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DouyinGetUserInfoResponseMultiError, or nil if none found.
func (m *DouyinGetUserInfoResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DouyinGetUserInfoResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for StatusCode

	if m.StatusMsg != nil {
		// no validation rules for StatusMsg
	}

	if m.Info != nil {

		if all {
			switch v := interface{}(m.GetInfo()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DouyinGetUserInfoResponseValidationError{
						field:  "Info",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DouyinGetUserInfoResponseValidationError{
						field:  "Info",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetInfo()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DouyinGetUserInfoResponseValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return DouyinGetUserInfoResponseMultiError(errors)
	}

	return nil
}

// DouyinGetUserInfoResponseMultiError is an error wrapping multiple validation
// errors returned by DouyinGetUserInfoResponse.ValidateAll() if the
// designated constraints aren't met.
type DouyinGetUserInfoResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DouyinGetUserInfoResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DouyinGetUserInfoResponseMultiError) AllErrors() []error { return m }

// DouyinGetUserInfoResponseValidationError is the validation error returned by
// DouyinGetUserInfoResponse.Validate if the designated constraints aren't met.
type DouyinGetUserInfoResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DouyinGetUserInfoResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DouyinGetUserInfoResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DouyinGetUserInfoResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DouyinGetUserInfoResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DouyinGetUserInfoResponseValidationError) ErrorName() string {
	return "DouyinGetUserInfoResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DouyinGetUserInfoResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDouyinGetUserInfoResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DouyinGetUserInfoResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DouyinGetUserInfoResponseValidationError{}

// Validate checks the field values on DouyinMultipleGetUserInfoRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *DouyinMultipleGetUserInfoRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DouyinMultipleGetUserInfoRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// DouyinMultipleGetUserInfoRequestMultiError, or nil if none found.
func (m *DouyinMultipleGetUserInfoRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DouyinMultipleGetUserInfoRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DouyinMultipleGetUserInfoRequestMultiError(errors)
	}

	return nil
}

// DouyinMultipleGetUserInfoRequestMultiError is an error wrapping multiple
// validation errors returned by
// DouyinMultipleGetUserInfoRequest.ValidateAll() if the designated
// constraints aren't met.
type DouyinMultipleGetUserInfoRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DouyinMultipleGetUserInfoRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DouyinMultipleGetUserInfoRequestMultiError) AllErrors() []error { return m }

// DouyinMultipleGetUserInfoRequestValidationError is the validation error
// returned by DouyinMultipleGetUserInfoRequest.Validate if the designated
// constraints aren't met.
type DouyinMultipleGetUserInfoRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DouyinMultipleGetUserInfoRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DouyinMultipleGetUserInfoRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DouyinMultipleGetUserInfoRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DouyinMultipleGetUserInfoRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DouyinMultipleGetUserInfoRequestValidationError) ErrorName() string {
	return "DouyinMultipleGetUserInfoRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DouyinMultipleGetUserInfoRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDouyinMultipleGetUserInfoRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DouyinMultipleGetUserInfoRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DouyinMultipleGetUserInfoRequestValidationError{}

// Validate checks the field values on DouyinMultipleGetUserInfoResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *DouyinMultipleGetUserInfoResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DouyinMultipleGetUserInfoResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// DouyinMultipleGetUserInfoResponseMultiError, or nil if none found.
func (m *DouyinMultipleGetUserInfoResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DouyinMultipleGetUserInfoResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for StatusCode

	for idx, item := range m.GetInfos() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DouyinMultipleGetUserInfoResponseValidationError{
						field:  fmt.Sprintf("Infos[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DouyinMultipleGetUserInfoResponseValidationError{
						field:  fmt.Sprintf("Infos[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DouyinMultipleGetUserInfoResponseValidationError{
					field:  fmt.Sprintf("Infos[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.StatusMsg != nil {
		// no validation rules for StatusMsg
	}

	if len(errors) > 0 {
		return DouyinMultipleGetUserInfoResponseMultiError(errors)
	}

	return nil
}

// DouyinMultipleGetUserInfoResponseMultiError is an error wrapping multiple
// validation errors returned by
// DouyinMultipleGetUserInfoResponse.ValidateAll() if the designated
// constraints aren't met.
type DouyinMultipleGetUserInfoResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DouyinMultipleGetUserInfoResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DouyinMultipleGetUserInfoResponseMultiError) AllErrors() []error { return m }

// DouyinMultipleGetUserInfoResponseValidationError is the validation error
// returned by DouyinMultipleGetUserInfoResponse.Validate if the designated
// constraints aren't met.
type DouyinMultipleGetUserInfoResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DouyinMultipleGetUserInfoResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DouyinMultipleGetUserInfoResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DouyinMultipleGetUserInfoResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DouyinMultipleGetUserInfoResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DouyinMultipleGetUserInfoResponseValidationError) ErrorName() string {
	return "DouyinMultipleGetUserInfoResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DouyinMultipleGetUserInfoResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDouyinMultipleGetUserInfoResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DouyinMultipleGetUserInfoResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DouyinMultipleGetUserInfoResponseValidationError{}

// Validate checks the field values on UserInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserInfoMultiError, or nil
// if none found.
func (m *UserInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *UserInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for IsExist

	// no validation rules for Id

	// no validation rules for Name

	if m.Avatar != nil {
		// no validation rules for Avatar
	}

	if m.BackgroundImage != nil {
		// no validation rules for BackgroundImage
	}

	if m.Signature != nil {
		// no validation rules for Signature
	}

	if len(errors) > 0 {
		return UserInfoMultiError(errors)
	}

	return nil
}

// UserInfoMultiError is an error wrapping multiple validation errors returned
// by UserInfo.ValidateAll() if the designated constraints aren't met.
type UserInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserInfoMultiError) AllErrors() []error { return m }

// UserInfoValidationError is the validation error returned by
// UserInfo.Validate if the designated constraints aren't met.
type UserInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserInfoValidationError) ErrorName() string { return "UserInfoValidationError" }

// Error satisfies the builtin error interface
func (e UserInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserInfoValidationError{}
