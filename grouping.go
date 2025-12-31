package vaddie

// Validator is an interface that structs can implement enabling
// the use of wrappers.
type Validator interface {
	Validate() error
}

// ValidateValue can be used to validate a field value.
type ValidateValue[T any] func(value T) error

// AllOf validates that our value meets all of the validaton rules.
// If T implements the [Validator] interface, it is validated first.
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
// If T implements the [Validator] interface, it is validated first but does not
// immediately return nil if it passes.
// Instead we will still go through the validate funcs until one other check passes.
func OneOf[T any](value T, key string, validateValues ...ValidateValue[T]) error {
	errs := make([]error, 0, len(validateValues))

	if v, isValidator := (any(value)).(Validator); isValidator {
		if err := v.Validate(); err != nil {
			errs = append(errs, expandErrorKey(err, key))
		}
	}

	for _, validation := range validateValues {
		err := validation(value)
		if err == nil {
			return nil
		}

		errs = append(errs, expandErrorKey(err, key))
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

// Optional will validate a value meets our rules if and only if it is not nil.
// A nil value will always meet validation.
// If T implements the [Validator] interface, and is not nil, that is run first.
func Optional[T any](value *T, key string, validateValues ...ValidateValue[T]) error {
	if value == nil {
		return nil
	}

	errs := make([]error, 0)

	if v, isValidator := (any(value)).(Validator); isValidator {
		if err := v.Validate(); err != nil {
			errs = append(errs, expandErrorKey(err, key))
		}
	}

	for _, validation := range validateValues {
		err := validation(*value)
		if err != nil {
			errs = append(errs, expandErrorKey(err, key))
		}
	}

	return Join(errs...)
}

// Required will validate a value meets our rules if and only if it is not nil.
// A nil value will error immediately.
// If T implements the [Validator] interface, and is not nil, that is run first.
func Required[T any](value *T, key string, validateValues ...ValidateValue[T]) error {
	if value == nil {
		return &ValidationError{
			Key:     key,
			Message: "is nil",
		}
	}

	errs := make([]error, 0)

	if v, isValidator := (any(value)).(Validator); isValidator {
		if err := v.Validate(); err != nil {
			errs = append(errs, expandErrorKey(err, key))
		}
	}

	for _, validation := range validateValues {
		err := validation(*value)
		if err != nil {
			errs = append(errs, expandErrorKey(err, key))
		}
	}

	return Join(errs...)
}
