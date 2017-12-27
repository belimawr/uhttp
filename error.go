package uhttp

type httpError struct {
	err              error
	statusCode       int
	userFriendlyBody []byte
	headers          map[string]string
}

// NewHTTPError returns an implementation of HTTPError
func NewHTTPError(cause error,
	status int,
	body []byte,
	headers map[string]string) HTTPError {

	return httpError{
		err:              cause,
		statusCode:       status,
		userFriendlyBody: body,
		headers:          headers,
	}
}

func (err httpError) Error() string {
	return err.err.Error()
}

func (err httpError) StatusCode() int {
	return err.statusCode
}

func (err httpError) UserFriendlyError() []byte {
	return err.userFriendlyBody
}

func (err httpError) Cause() error {
	return err.err
}

func (err httpError) Headers() map[string]string {
	return err.headers
}
