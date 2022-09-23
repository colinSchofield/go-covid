package custom_error

/*
	custom_error -- this package contains the set of custom errors, that are typically used in communication between the
	service and controller layers.
*/

// This will occur when the (unreliable!) RestAPI service is down or unavailable
type ClientTimeout struct {
	Wrapped error
}

func (ct ClientTimeout) Error() string {
	return ct.Wrapped.Error()
}
