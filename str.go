package vaddie

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// StrEmpty validates our value is empty.
func StrEmpty() ValidateValue[string] {
	return func(value string) error {
		length := len(value)
		if length != 0 {
			return &ValidationError{
				Message: "not empty",
				Help:    fmt.Sprintf("%d > 0", length),
			}
		}

		return nil
	}
}

// StrSpace validates our value is entirely made up of unicode space characters.
func StrSpace() ValidateValue[string] {
	return func(value string) error {
		for _, v := range value {
			if !unicode.IsSpace(v) {
				return &ValidationError{
					Message: "not entirely whitespace",
					Help:    fmt.Sprintf("%q has printable character %q", value, v),
				}
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

// StrUnicodeMin validates our value is at least a minimum length of unicode characters.
// This properly compares strings that include things such as emojis and CJK symbols.
func StrUnicodeMin(minLength int) ValidateValue[string] {
	return func(value string) error {
		length := utf8.RuneCountInString(value)
		if length < minLength {
			return &ValidationError{
				Message: "unicode length too short",
				Help:    fmt.Sprintf("%d < %d", length, minLength),
			}
		}

		return nil
	}
}

// StrUnicodeMax validates our value is no more then a maximum length of unicode characters.
// This properly compares strings that include things such as emojis and CJK symbols.
func StrUnicodeMax(maxLength int) ValidateValue[string] {
	return func(value string) error {
		length := utf8.RuneCountInString(value)
		if length > maxLength {
			return &ValidationError{
				Message: "unicode length too long",
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
