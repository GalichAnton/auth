// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: auth.proto

package auth_v1

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

// Validate checks the field values on LoginRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginRequestMultiError, or
// nil if none found.
func (m *LoginRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateEmail(m.GetEmail()); err != nil {
		err = LoginRequestValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetPassword()); l < 6 || l > 24 {
		err := LoginRequestValidationError{
			field:  "Password",
			reason: "value length must be between 6 and 24 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return LoginRequestMultiError(errors)
	}

	return nil
}

func (m *LoginRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *LoginRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// LoginRequestMultiError is an error wrapping multiple validation errors
// returned by LoginRequest.ValidateAll() if the designated constraints aren't met.
type LoginRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginRequestMultiError) AllErrors() []error { return m }

// LoginRequestValidationError is the validation error returned by
// LoginRequest.Validate if the designated constraints aren't met.
type LoginRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginRequestValidationError) ErrorName() string { return "LoginRequestValidationError" }

// Error satisfies the builtin error interface
func (e LoginRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginRequestValidationError{}

// Validate checks the field values on LoginResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginResponseMultiError, or
// nil if none found.
func (m *LoginResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RefreshToken

	if len(errors) > 0 {
		return LoginResponseMultiError(errors)
	}

	return nil
}

// LoginResponseMultiError is an error wrapping multiple validation errors
// returned by LoginResponse.ValidateAll() if the designated constraints
// aren't met.
type LoginResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginResponseMultiError) AllErrors() []error { return m }

// LoginResponseValidationError is the validation error returned by
// LoginResponse.Validate if the designated constraints aren't met.
type LoginResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginResponseValidationError) ErrorName() string { return "LoginResponseValidationError" }

// Error satisfies the builtin error interface
func (e LoginResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginResponseValidationError{}

// Validate checks the field values on GetRefreshTokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetRefreshTokenRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetRefreshTokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetRefreshTokenRequestMultiError, or nil if none found.
func (m *GetRefreshTokenRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetRefreshTokenRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RefreshToken

	if len(errors) > 0 {
		return GetRefreshTokenRequestMultiError(errors)
	}

	return nil
}

// GetRefreshTokenRequestMultiError is an error wrapping multiple validation
// errors returned by GetRefreshTokenRequest.ValidateAll() if the designated
// constraints aren't met.
type GetRefreshTokenRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetRefreshTokenRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetRefreshTokenRequestMultiError) AllErrors() []error { return m }

// GetRefreshTokenRequestValidationError is the validation error returned by
// GetRefreshTokenRequest.Validate if the designated constraints aren't met.
type GetRefreshTokenRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRefreshTokenRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRefreshTokenRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRefreshTokenRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRefreshTokenRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRefreshTokenRequestValidationError) ErrorName() string {
	return "GetRefreshTokenRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetRefreshTokenRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRefreshTokenRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRefreshTokenRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRefreshTokenRequestValidationError{}

// Validate checks the field values on GetRefreshTokenResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetRefreshTokenResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetRefreshTokenResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetRefreshTokenResponseMultiError, or nil if none found.
func (m *GetRefreshTokenResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetRefreshTokenResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RefreshToken

	if len(errors) > 0 {
		return GetRefreshTokenResponseMultiError(errors)
	}

	return nil
}

// GetRefreshTokenResponseMultiError is an error wrapping multiple validation
// errors returned by GetRefreshTokenResponse.ValidateAll() if the designated
// constraints aren't met.
type GetRefreshTokenResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetRefreshTokenResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetRefreshTokenResponseMultiError) AllErrors() []error { return m }

// GetRefreshTokenResponseValidationError is the validation error returned by
// GetRefreshTokenResponse.Validate if the designated constraints aren't met.
type GetRefreshTokenResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRefreshTokenResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRefreshTokenResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRefreshTokenResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRefreshTokenResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRefreshTokenResponseValidationError) ErrorName() string {
	return "GetRefreshTokenResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetRefreshTokenResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRefreshTokenResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRefreshTokenResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRefreshTokenResponseValidationError{}

// Validate checks the field values on GetAccessTokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetAccessTokenRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAccessTokenRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAccessTokenRequestMultiError, or nil if none found.
func (m *GetAccessTokenRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAccessTokenRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RefreshToken

	if len(errors) > 0 {
		return GetAccessTokenRequestMultiError(errors)
	}

	return nil
}

// GetAccessTokenRequestMultiError is an error wrapping multiple validation
// errors returned by GetAccessTokenRequest.ValidateAll() if the designated
// constraints aren't met.
type GetAccessTokenRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAccessTokenRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAccessTokenRequestMultiError) AllErrors() []error { return m }

// GetAccessTokenRequestValidationError is the validation error returned by
// GetAccessTokenRequest.Validate if the designated constraints aren't met.
type GetAccessTokenRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAccessTokenRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAccessTokenRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAccessTokenRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAccessTokenRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAccessTokenRequestValidationError) ErrorName() string {
	return "GetAccessTokenRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetAccessTokenRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAccessTokenRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAccessTokenRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAccessTokenRequestValidationError{}

// Validate checks the field values on GetAccessTokenResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetAccessTokenResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAccessTokenResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAccessTokenResponseMultiError, or nil if none found.
func (m *GetAccessTokenResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAccessTokenResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccessToken

	if len(errors) > 0 {
		return GetAccessTokenResponseMultiError(errors)
	}

	return nil
}

// GetAccessTokenResponseMultiError is an error wrapping multiple validation
// errors returned by GetAccessTokenResponse.ValidateAll() if the designated
// constraints aren't met.
type GetAccessTokenResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAccessTokenResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAccessTokenResponseMultiError) AllErrors() []error { return m }

// GetAccessTokenResponseValidationError is the validation error returned by
// GetAccessTokenResponse.Validate if the designated constraints aren't met.
type GetAccessTokenResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAccessTokenResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAccessTokenResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAccessTokenResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAccessTokenResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAccessTokenResponseValidationError) ErrorName() string {
	return "GetAccessTokenResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetAccessTokenResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAccessTokenResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAccessTokenResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAccessTokenResponseValidationError{}