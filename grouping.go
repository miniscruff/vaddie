package vaddy

// ValidateValue can be used to validate a field value.
type ValidateValue[T any] func(value T) error

// Validator is an interface that structs can implement enabling
// the use of wrappers
type Validator interface {
	Validate() error
}

// Potential name
func Is[T any](value T, key string, validateValues ...ValidateValue[T]) error {
	errs := make([]error, 0)
	for _, validation := range validateValues {
		err := validation(value)
		if err != nil {
			errs = append(errs, expandErrorKey(err, key))
		}

	}

	return Join(errs...)
}

func OneOf[T any](value T, validateValues ...ValidateValue[T]) error {
	errs := make([]error, len(validateValues))
	for i, validation := range validateValues {
		err := validation(value)
		if err == nil {
			return nil
		}

		errs[i] = err

	}

	return Join(errs...)
}

func And[T any](validateValues ...ValidateValue[T]) ValidateValue[T] {
	return func(value T) error {
		errs := make([]error, 0)
		for _, validation := range validateValues {
			err := validation(value)
			if err != nil {
				errs = append(errs, err)
			}

		}

		return Join(errs...)
	}
}

func Or[T any](validateValues ...ValidateValue[T]) ValidateValue[T] {
	return func(value T) error {
		errs := make([]error, len(validateValues))
		for i, validation := range validateValues {
			err := validation(value)
			if err == nil {
				return nil
			}

			errs[i] = err
		}

		return Join(errs...)
	}
}
