package presenter

type AppHTTPError struct {
	Msg  string
	Code int
	Err  error
}

func (e *AppHTTPError) Error() string { return e.Msg + " : " + e.Err.Error() }

type MongoDBError struct {
	Msg  string
	Code int
	Err  error
}

func (e *MongoDBError) Error() string { return e.Msg + " : " + e.Err.Error() }

const (
	ERR_NOTFOUND = 001
	ERR_SYNTAX   = 002
)
