package vaddy

// Validator is an interface that structs can implement enabling
// the use of wrappers.
type Validator interface {
	Validate() error
}

// ValidateValue can be used to validate a field value.
type ValidateValue[T any] func(value T) error

// AllOf validates that our value meets all of the validaton rules.
// In addition to the validation rules, if T implements [Validator] it will also
// be called.
func AllOf[T any](value T, key string, validateValues ...ValidateValue[T]) error {
	errs := make([]error, 0)

	if v, isValidator := (any(value)).(Validator); isValidator {
		if err := v.Validate(); err != nil {
			errs = append(errs, expandErrorKey(err, key))
		}
	}

	for _, validation := range validateValues {
		err := validation(value)
		if err != nil {
			errs = append(errs, expandErrorKey(err, key))
		}
	}

	return Join(errs...)
}

// OneOf validates that our value meets at least one of the validaton rules.
// In addition to the validation rules, if T implements [Validator] it will also
// be called. However, meeting this validation does not count as the one of.
func OneOf[T any](value T, key string, validateValues ...ValidateValue[T]) error {
	errs := make([]error, len(validateValues))

	if v, isValidator := (any(value)).(Validator); isValidator {
		if err := v.Validate(); err != nil {
			errs = append(errs, expandErrorKey(err, key))
		}
	}

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

		return Join(errs...)
	}
}
