package ports

type OutputBoundary[T any] interface {
	Unwrap() (T, error)
	Error() error
	Response() T
}

type Output[T any] struct {
	err      error
	response T
}

func NewOutput[T any](err error, response T) OutputBoundary[T] {
	return &Output[T]{err: err, response: response}
}

func (o *Output[T]) Unwrap() (T, error) {
	return o.response, o.err
}

func (o *Output[T]) Error() error {
	return o.err
}

func (o *Output[T]) Response() T {
	return o.response
}
