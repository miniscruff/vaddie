package vaddie

// If only runs validations when cond is true.
//
// Instead of defining a custom error func that first does an if check, you
// can make a simpler condition check.
func If(cond bool, errs ...error) error {
	if !cond {
		return nil
	}

	return Join(errs...)
}
