package custom_error

// This error wraps any JSON validation issues
type Validation struct {
	Wrapped error
}

func (v Validation) Error() string {
	return v.Wrapped.Error()
}
