package vaddy

import (
	"cmp"
	"fmt"
)

func OrderedGt[T cmp.Ordered](minValue T) ValidateValue[T] {
	return func(value T) error {
		if value <= minValue {
			return fmt.Errorf("value too low (%v <= %v)", value, minValue)
		}

		return nil
	}
}

func OrderedGte[T cmp.Ordered](minValue T) ValidateValue[T] {
	return func(value T) error {
		if value < minValue {
			return fmt.Errorf("value too low (%v < %v)", value, minValue)
		}

		return nil
	}
}

func OrderedLt[T cmp.Ordered](minValue T) ValidateValue[T] {
	return func(value T) error {
		if value >= minValue {
			return fmt.Errorf("value too high (%v >= %v)", value, minValue)
		}

		return nil
	}
}

func OrderedLte[T cmp.Ordered](minValue T) ValidateValue[T] {
	return func(value T) error {
		if value > minValue {
			return fmt.Errorf("value too high (%v > %v)", value, minValue)
		}

		return nil
	}
}
