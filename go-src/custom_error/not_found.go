package custom_error

/*
	custom_error -- this package contains the set of custom errors, that are typically used in communication between the
	service and controller layers.
*/

// This error occurs whenever the uid for a user in DynamoDB (or the iso for the country) cannot be found
type NotFound struct {
	Wrapped error
}

func (ct NotFound) Error() string {
	return ct.Wrapped.Error()
}
