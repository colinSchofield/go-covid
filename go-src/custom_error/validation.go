package custom_error

/*
	custom_error -- this package contains the set of custom errors, that are typically used in communication between the
	service and controller layers.
*/

// This error wraps any JSON validation issues
type Validation struct {
	Wrapped error
}

func (v Validation) Error() string {
	return v.Wrapped.Error()
}
