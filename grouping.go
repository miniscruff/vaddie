package vaddie

// Validator is an interface that structs can implement enabling
// the use of wrappers.
type Validator interface {
	Validate() error
}

// ValidateValue can be used to validate a field value.
type ValidateValue[T any] func(value T) error

// AllOf validates that our value meets all of the validaton rules.
func AllOf[T any](value T, key string, validateValues ...ValidateValue[T]) error {
	errs := make([]error, 0)

	for _, validation := range validateValues {
		err := validation(value)
		if err != nil {
			errs = append(errs, expandErrorKey(err, key))
		}
	}

	return Join(errs...)
}

// OneOf validates that our value meets at least one of the validaton rules.
func OneOf[T any](value T, key string, validateValues ...ValidateValue[T]) error {
	errs := make([]error, len(validateValues))

	for i, validation := range validateValues {
		err := validation(value)
		if err == nil {
			return nil
		}

		errs[i] = expandErrorKey(err, key)
	}

	return Join(errs...)
}

// And combines many validation rules into one.
// All validations must be true for the validation to be successful.
func And[T any](validateValues ...ValidateValue[T]) ValidateValue[T] {
	return func(value T) error {
		errs := make([]error, 0)

		for _, validation := range validateValues {
			err := validation(value)
			if err != nil {
				errs = append(errs, err)
			}
		}

		// TODO: should this return a validation error?
		return Join(errs...)
	}
}

// Or combines many validation rules into one.
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

		// TODO: should this return a validation error?
		return Join(errs...)
	}
}

// Optional will validate a value meets our rules if and only if it is not nil.
// A nil value will always meet validation.
func Optional[T any](value *T, key string, validateValues ...ValidateValue[T]) error {
	if value == nil {
		return nil
	}

	errs := make([]error, 0)

	for _, validation := range validateValues {
		err := validation(*value)
		if err != nil {
			errs = append(errs, expandErrorKey(err, key))
		}
	}

	return Join(errs...)
}
