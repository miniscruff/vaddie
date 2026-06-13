package vaddie

import "fmt"

type WholeNumeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type FractionalNumeric interface {
	~float32 | ~float64
}

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
	~float32 | ~float64
}

func NumericMultipleOf[T WholeNumeric](mult T) ValidateValue[T] {
	return func(value T) error {
		if mult <= 0 {
			return &ValidationError{
				Message: "attempted to divided by zero",
				Help:    fmt.Sprintf("%d <= 0", mult),
			}
		}

		if value % mult != 0 {
			return &ValidationError{
				Message: "values are not multiples",
				Help:    fmt.Sprintf("%d %% %d != 0", value, mult),
			}
		}

		return nil
	}
}
