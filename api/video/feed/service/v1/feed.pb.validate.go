// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: video/feed/service/v1/feed.proto

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

// Validate checks the field values on FeedRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FeedRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FeedRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FeedRequestMultiError, or
// nil if none found.
func (m *FeedRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *FeedRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.LatestTime != nil {
		// no validation rules for LatestTime
	}

	if m.UserId != nil {
		// no validation rules for UserId
	}

	if len(errors) > 0 {
		return FeedRequestMultiError(errors)
	}

	return nil
}

// FeedRequestMultiError is an error wrapping multiple validation errors
// returned by FeedRequest.ValidateAll() if the designated constraints aren't met.
type FeedRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FeedRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FeedRequestMultiError) AllErrors() []error { return m }

// FeedRequestValidationError is the validation error returned by
// FeedRequest.Validate if the designated constraints aren't met.
type FeedRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FeedRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FeedRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FeedRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FeedRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FeedRequestValidationError) ErrorName() string { return "FeedRequestValidationError" }

// Error satisfies the builtin error interface
func (e FeedRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFeedRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FeedRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FeedRequestValidationError{}

// Validate checks the field values on FeedResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FeedResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FeedResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FeedResponseMultiError, or
// nil if none found.
func (m *FeedResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *FeedResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for StatusCode

	for idx, item := range m.GetVideoList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, FeedResponseValidationError{
						field:  fmt.Sprintf("VideoList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, FeedResponseValidationError{
						field:  fmt.Sprintf("VideoList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FeedResponseValidationError{
					field:  fmt.Sprintf("VideoList[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.StatusMsg != nil {
		// no validation rules for StatusMsg
	}

	if m.NextTime != nil {
		// no validation rules for NextTime
	}

	if len(errors) > 0 {
		return FeedResponseMultiError(errors)
	}

	return nil
}

// FeedResponseMultiError is an error wrapping multiple validation errors
// returned by FeedResponse.ValidateAll() if the designated constraints aren't met.
type FeedResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FeedResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FeedResponseMultiError) AllErrors() []error { return m }

// FeedResponseValidationError is the validation error returned by
// FeedResponse.Validate if the designated constraints aren't met.
type FeedResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FeedResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FeedResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FeedResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FeedResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FeedResponseValidationError) ErrorName() string { return "FeedResponseValidationError" }

// Error satisfies the builtin error interface
func (e FeedResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFeedResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FeedResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FeedResponseValidationError{}

// Validate checks the field values on GetPublishedVideoByUserIdRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *GetPublishedVideoByUserIdRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetPublishedVideoByUserIdRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// GetPublishedVideoByUserIdRequestMultiError, or nil if none found.
func (m *GetPublishedVideoByUserIdRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetPublishedVideoByUserIdRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for ToUserId

	if len(errors) > 0 {
		return GetPublishedVideoByUserIdRequestMultiError(errors)
	}

	return nil
}

// GetPublishedVideoByUserIdRequestMultiError is an error wrapping multiple
// validation errors returned by
// GetPublishedVideoByUserIdRequest.ValidateAll() if the designated
// constraints aren't met.
type GetPublishedVideoByUserIdRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetPublishedVideoByUserIdRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetPublishedVideoByUserIdRequestMultiError) AllErrors() []error { return m }

// GetPublishedVideoByUserIdRequestValidationError is the validation error
// returned by GetPublishedVideoByUserIdRequest.Validate if the designated
// constraints aren't met.
type GetPublishedVideoByUserIdRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPublishedVideoByUserIdRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPublishedVideoByUserIdRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPublishedVideoByUserIdRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPublishedVideoByUserIdRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPublishedVideoByUserIdRequestValidationError) ErrorName() string {
	return "GetPublishedVideoByUserIdRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetPublishedVideoByUserIdRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPublishedVideoByUserIdRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPublishedVideoByUserIdRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPublishedVideoByUserIdRequestValidationError{}

// Validate checks the field values on GetPublishedVideoByUserIdResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *GetPublishedVideoByUserIdResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetPublishedVideoByUserIdResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// GetPublishedVideoByUserIdResponseMultiError, or nil if none found.
func (m *GetPublishedVideoByUserIdResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetPublishedVideoByUserIdResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for StatusCode

	for idx, item := range m.GetVideoList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetPublishedVideoByUserIdResponseValidationError{
						field:  fmt.Sprintf("VideoList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetPublishedVideoByUserIdResponseValidationError{
						field:  fmt.Sprintf("VideoList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetPublishedVideoByUserIdResponseValidationError{
					field:  fmt.Sprintf("VideoList[%v]", idx),
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
		return GetPublishedVideoByUserIdResponseMultiError(errors)
	}

	return nil
}

// GetPublishedVideoByUserIdResponseMultiError is an error wrapping multiple
// validation errors returned by
// GetPublishedVideoByUserIdResponse.ValidateAll() if the designated
// constraints aren't met.
type GetPublishedVideoByUserIdResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetPublishedVideoByUserIdResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetPublishedVideoByUserIdResponseMultiError) AllErrors() []error { return m }

// GetPublishedVideoByUserIdResponseValidationError is the validation error
// returned by GetPublishedVideoByUserIdResponse.Validate if the designated
// constraints aren't met.
type GetPublishedVideoByUserIdResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPublishedVideoByUserIdResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPublishedVideoByUserIdResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPublishedVideoByUserIdResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPublishedVideoByUserIdResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPublishedVideoByUserIdResponseValidationError) ErrorName() string {
	return "GetPublishedVideoByUserIdResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetPublishedVideoByUserIdResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPublishedVideoByUserIdResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPublishedVideoByUserIdResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPublishedVideoByUserIdResponseValidationError{}

// Validate checks the field values on GetUserFavoriteVideoListByUserIdRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *GetUserFavoriteVideoListByUserIdRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// GetUserFavoriteVideoListByUserIdRequest with the rules defined in the proto
// definition for this message. If any rules are violated, the result is a
// list of violation errors wrapped in
// GetUserFavoriteVideoListByUserIdRequestMultiError, or nil if none found.
func (m *GetUserFavoriteVideoListByUserIdRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserFavoriteVideoListByUserIdRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for ToUserId

	if len(errors) > 0 {
		return GetUserFavoriteVideoListByUserIdRequestMultiError(errors)
	}

	return nil
}

// GetUserFavoriteVideoListByUserIdRequestMultiError is an error wrapping
// multiple validation errors returned by
// GetUserFavoriteVideoListByUserIdRequest.ValidateAll() if the designated
// constraints aren't met.
type GetUserFavoriteVideoListByUserIdRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserFavoriteVideoListByUserIdRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserFavoriteVideoListByUserIdRequestMultiError) AllErrors() []error { return m }

// GetUserFavoriteVideoListByUserIdRequestValidationError is the validation
// error returned by GetUserFavoriteVideoListByUserIdRequest.Validate if the
// designated constraints aren't met.
type GetUserFavoriteVideoListByUserIdRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserFavoriteVideoListByUserIdRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserFavoriteVideoListByUserIdRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserFavoriteVideoListByUserIdRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserFavoriteVideoListByUserIdRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserFavoriteVideoListByUserIdRequestValidationError) ErrorName() string {
	return "GetUserFavoriteVideoListByUserIdRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserFavoriteVideoListByUserIdRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserFavoriteVideoListByUserIdRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserFavoriteVideoListByUserIdRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserFavoriteVideoListByUserIdRequestValidationError{}

// Validate checks the field values on GetUserFavoriteVideoListByUserIdResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *GetUserFavoriteVideoListByUserIdResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// GetUserFavoriteVideoListByUserIdResponse with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// GetUserFavoriteVideoListByUserIdResponseMultiError, or nil if none found.
func (m *GetUserFavoriteVideoListByUserIdResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserFavoriteVideoListByUserIdResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for StatusCode

	for idx, item := range m.GetVideoList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetUserFavoriteVideoListByUserIdResponseValidationError{
						field:  fmt.Sprintf("VideoList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetUserFavoriteVideoListByUserIdResponseValidationError{
						field:  fmt.Sprintf("VideoList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetUserFavoriteVideoListByUserIdResponseValidationError{
					field:  fmt.Sprintf("VideoList[%v]", idx),
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
		return GetUserFavoriteVideoListByUserIdResponseMultiError(errors)
	}

	return nil
}

// GetUserFavoriteVideoListByUserIdResponseMultiError is an error wrapping
// multiple validation errors returned by
// GetUserFavoriteVideoListByUserIdResponse.ValidateAll() if the designated
// constraints aren't met.
type GetUserFavoriteVideoListByUserIdResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserFavoriteVideoListByUserIdResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserFavoriteVideoListByUserIdResponseMultiError) AllErrors() []error { return m }

// GetUserFavoriteVideoListByUserIdResponseValidationError is the validation
// error returned by GetUserFavoriteVideoListByUserIdResponse.Validate if the
// designated constraints aren't met.
type GetUserFavoriteVideoListByUserIdResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserFavoriteVideoListByUserIdResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserFavoriteVideoListByUserIdResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserFavoriteVideoListByUserIdResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserFavoriteVideoListByUserIdResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserFavoriteVideoListByUserIdResponseValidationError) ErrorName() string {
	return "GetUserFavoriteVideoListByUserIdResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserFavoriteVideoListByUserIdResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserFavoriteVideoListByUserIdResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserFavoriteVideoListByUserIdResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserFavoriteVideoListByUserIdResponseValidationError{}

// Validate checks the field values on Video with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Video) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Video with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in VideoMultiError, or nil if none found.
func (m *Video) ValidateAll() error {
	return m.validate(true)
}

func (m *Video) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetAuthor()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, VideoValidationError{
					field:  "Author",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, VideoValidationError{
					field:  "Author",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAuthor()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return VideoValidationError{
				field:  "Author",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for PlayUrl

	// no validation rules for CoverUrl

	// no validation rules for FavoriteCount

	// no validation rules for CommentCount

	// no validation rules for IsFavorite

	// no validation rules for Title

	if len(errors) > 0 {
		return VideoMultiError(errors)
	}

	return nil
}

// VideoMultiError is an error wrapping multiple validation errors returned by
// Video.ValidateAll() if the designated constraints aren't met.
type VideoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VideoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VideoMultiError) AllErrors() []error { return m }

// VideoValidationError is the validation error returned by Video.Validate if
// the designated constraints aren't met.
type VideoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VideoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VideoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VideoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VideoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VideoValidationError) ErrorName() string { return "VideoValidationError" }

// Error satisfies the builtin error interface
func (e VideoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVideo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VideoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VideoValidationError{}

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *User) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on User with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in UserMultiError, or nil if none found.
func (m *User) ValidateAll() error {
	return m.validate(true)
}

func (m *User) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for IsFollow

	if m.FollowCount != nil {
		// no validation rules for FollowCount
	}

	if m.FollowerCount != nil {
		// no validation rules for FollowerCount
	}

	if m.Avatar != nil {
		// no validation rules for Avatar
	}

	if m.BackgroundImage != nil {
		// no validation rules for BackgroundImage
	}

	if m.Signature != nil {
		// no validation rules for Signature
	}

	if m.TotalFavorited != nil {
		// no validation rules for TotalFavorited
	}

	if m.WorkCount != nil {
		// no validation rules for WorkCount
	}

	if m.FavoriteCount != nil {
		// no validation rules for FavoriteCount
	}

	if len(errors) > 0 {
		return UserMultiError(errors)
	}

	return nil
}

// UserMultiError is an error wrapping multiple validation errors returned by
// User.ValidateAll() if the designated constraints aren't met.
type UserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserMultiError) AllErrors() []error { return m }

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}
