package output

type ErrorOutputPort interface {
	BadRequest(err error)
	InternalServerError(err error)
	NotFound(err error)
	Unauthorized(err error)
}
