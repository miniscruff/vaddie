package vaddie

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// StrNotEmpty validates that a given string is not empty.
func StrNotEmpty() ValidateValue[string] {
	return func(value string) error {
		if value == "" {
			return &ValidationError{
				Message: "is empty",
			}
		}

		return nil
	}
}

// StrMin validates our value is at least a minimum length.
func StrMin(minLength int) ValidateValue[string] {
	return func(value string) error {
		length := len(value)
		if length < minLength {
			return &ValidationError{
				Message: "too short",
				Help:    fmt.Sprintf("%d < %d", length, minLength),
			}
		}

		return nil
	}
}

// StrMax validates our value is no more then a maximum length.
func StrMax(maxLength int) ValidateValue[string] {
	return func(value string) error {
		length := len(value)
		if length > maxLength {
			return &ValidationError{
				Message: "too long",
				Help:    fmt.Sprintf("%d > %d", length, maxLength),
			}
		}

		return nil
	}
}

// StrLetters validates every rune is a letter.
func StrLetters() ValidateValue[string] {
	return func(value string) error {
		for i, v := range value {
			if !unicode.IsLetter(v) {
				return &ValidationError{
					Message: "non-letter rune",
					Help:    fmt.Sprintf("%q at index %d", v, i),
				}
			}
		}

		return nil
	}
}

// StrAscii validates every rune is an ascii value.
func StrAscii() ValidateValue[string] {
	return func(value string) error {
		for i, v := range value {
			if v > unicode.MaxASCII {
				return &ValidationError{
					Message: "non-ascii rune",
					Help:    fmt.Sprintf("%q at index %d", v, i),
				}
			}
		}

		return nil
	}
}

// StrHasPrefix validates our string has the provided prefix.
func StrHasPrefix(prefix string) ValidateValue[string] {
	return func(value string) error {
		if !strings.HasPrefix(value, prefix) {
			return &ValidationError{
				Message: "does not have prefix",
				Help:    fmt.Sprintf("%q does not have expected prefix %q", value, prefix),
			}
		}

		return nil
	}
}

// StrNotHasPrefix validates our string does not have the provided prefix.
func StrNotHasPrefix(prefix string) ValidateValue[string] {
	return func(value string) error {
		if strings.HasPrefix(value, prefix) {
			return &ValidationError{
				Message: "does have prefix",
				Help:    fmt.Sprintf("%q does have unexpected prefix %q", value, prefix),
			}
		}

		return nil
	}
}

// StrHasSuffix validates our string has the provided suffix.
func StrHasSuffix(suffix string) ValidateValue[string] {
	return func(value string) error {
		if !strings.HasSuffix(value, suffix) {
			return &ValidationError{
				Message: "does not have suffix",
				Help:    fmt.Sprintf("%q does not have expected suffix %q", value, suffix),
			}
		}

		return nil
	}
}

// StrNotHasSuffix validates our string does not have the provided suffix.
func StrNotHasSuffix(suffix string) ValidateValue[string] {
	return func(value string) error {
		if strings.HasSuffix(value, suffix) {
			return &ValidationError{
				Message: "does have suffix",
				Help:    fmt.Sprintf("%q does have unexpected suffix %q", value, suffix),
			}
		}

		return nil
	}
}

// StrContains validates our string contains the provided substring.
func StrContains(substr string) ValidateValue[string] {
	return func(value string) error {
		if !strings.Contains(value, substr) {
			return &ValidationError{
				Message: "does not have substr",
				Help:    fmt.Sprintf("%q does not have expected substr %q", value, substr),
			}
		}

		return nil
	}
}

// StrNotContains validates our string does not contain the provided substring.
func StrNotContains(substr string) ValidateValue[string] {
	return func(value string) error {
		if strings.Contains(value, substr) {
			return &ValidationError{
				Message: "does have substr",
				Help:    fmt.Sprintf("%q does have unexpected substr %q", value, substr),
			}
		}

		return nil
	}
}

// StrContainsAny validates whether any Unicode code points in chars are within value.
func StrContainsAny(chars string) ValidateValue[string] {
	return func(value string) error {
		if !strings.ContainsAny(value, chars) {
			return &ValidationError{
				Message: "does not have chars",
				Help:    fmt.Sprintf("%q does not have any of the chars %q", value, chars),
			}
		}

		return nil
	}
}

// StrNotContainsAny validates whether all Unicode code points in chars are not within value.
func StrNotContainsAny(chars string) ValidateValue[string] {
	return func(value string) error {
		if strings.ContainsAny(value, chars) {
			return &ValidationError{
				Message: "does have chars",
				Help:    fmt.Sprintf("%q does have unexpected chars %q", value, chars),
			}
		}

		return nil
	}
}

// StrMatch validates whether the value is matched by the provided regex.
// The regex is compiled once, if the regex is invalid every validation will fail but
// not panic.
// Use [StrRegexp] if you want to handle compiling the regexp yourself.
func StrMatch(reg string) ValidateValue[string] {
	rg, err := regexp.Compile(reg)

	return func(value string) error {
		if err != nil {
			return &ValidationError{
				Message: "regex did not compile",
				Help:    fmt.Sprintf("%q is invalid", reg),
			}
		}

		if !rg.MatchString(value) {
			return &ValidationError{
				Message: "does not match regex",
				Help:    fmt.Sprintf("%q does not match", value),
			}
		}

		return nil
	}
}

// StrRegexp validates whether the value is matched by the provided regex.
// Use [StrMatch] if you want vaddie to manage compiling the regexp.
func StrRegexp(rg *regexp.Regexp) ValidateValue[string] {
	return func(value string) error {
		if !rg.MatchString(value) {
			return &ValidationError{
				Message: "does not match regex",
				Help:    fmt.Sprintf("%q does not match", value),
			}
		}

		return nil
	}
}
