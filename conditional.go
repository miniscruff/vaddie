package vaddie

// IfExp is an experimental take on if branches.
//
// Instead of defining a custom error func that first does an if check, you
// can make a simpler condition check.
//
// Note: this is still experimental and may be dropped or changed suddenly.
func IfExp(cond bool, errs ...error) error {
	if !cond {
		return nil
	}

	return Join(errs...)
}
