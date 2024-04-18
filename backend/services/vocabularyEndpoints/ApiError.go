package vocabularyEndpoints

type APIError struct {
	Msg         string
	Status      int
	originalErr error
}

func (o APIError) Error() string {
	return o.Msg
}
