package custom_error

// This error occurs whenever the uid for a user in DynamoDB (or the iso for the country) cannot be found
type NotFound struct {
	Wrapped error
}

func (ct NotFound) Error() string {
	return ct.Wrapped.Error()
}
