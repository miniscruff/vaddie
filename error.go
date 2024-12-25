package vaddy

import (
	"errors"
	"strconv"
	"strings"
)

const delimiter = ": "

func toPtr[T any](value T) *T {
	return &value
}

// ValidationError is what we return on invalid validations
// Output of the error is a single string of:
// `[Message] Key [@Index] [(Help)]`
type ValidationError struct {
	Key     string
	Message string
	Help    string
	Index   *int
}

func (v *ValidationError) Error() string {
	sb := &strings.Builder{}

	if v.Key != "" {
		sb.WriteString(v.Key)
		sb.WriteString(delimiter)
	}

	if v.Index != nil {
		sb.WriteString("[")
		sb.WriteString(strconv.Itoa(*v.Index))
		sb.WriteString("]")
	}

	sb.WriteString(v.Message)

	if v.Help != "" {
		sb.WriteString(" ( ")
		sb.WriteString(v.Help)
		sb.WriteString(" )")
	}

	return sb.String()
}

func expandErrorKey(err error, key string) error {
	ve, isValidationError := err.(*ValidationError)
	if !isValidationError {
		return &ValidationError{
			Message: err.Error(),
			Key:     key,
		}
	}

	ve.Key = key
	return ve
}

func expandErrorKeyIndex(err error, key string, index int) error {
	ve, isValidationError := err.(*ValidationError)
	if !isValidationError {
		return &ValidationError{
			Message: err.Error(),
			Index:   &index,
			Key:     key,
		}
	}

	ve.Index = &index
	ve.Key = key
	return ve
}

// TODO: accept a Key?
func Join(errs ...error) error {
	return errors.Join(errs...)
}

// TODO:
// ValdiationErrors type
// Join should build ^
// Error() for ^ should support TextMarshal, JSONMarshal
