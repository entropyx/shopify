package shopify

const (
	ErrorUnauthorized = "[API] Invalid API key or access token (unrecognized login or wrong password)"
)

func IsErrorUnathorized(err error) bool {
	if err.Error() == ErrorUnauthorized {
		return true
	}
	return false
}
