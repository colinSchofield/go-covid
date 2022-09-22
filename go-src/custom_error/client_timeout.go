package custom_error

// This will occur when the (unreliable!) RestAPI service is down or unavailable
type ClientTimeout struct {
	Wrapped error
}

func (ct ClientTimeout) Error() string {
	return ct.Wrapped.Error()
}
