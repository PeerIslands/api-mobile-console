package presenter

type AppHTTPError struct {
	Msg  string
	Code int
	Err  error
}

func (e *AppHTTPError) Error() string { return e.Msg + " : " + e.Err.Error() }
