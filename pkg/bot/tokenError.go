package bot

type TokenError struct {
	Err error
}

func (e *TokenError) Error() string {
	return "Token is not set " + e.Err.Error()
}
