package ports

// OutputBoundary はユースケースの出力インターフェース
type OutputBoundary[T any] interface {
	Unwrap() (T, error)
	Error() error
	Response() T
}

// output はユースケースの出力
type output[T any] struct {
	err      error
	response T
}

// NewOutput はユースケースの出力を生成する
func NewOutput[T any](err error, response T) OutputBoundary[T] {
	return &output[T]{err: err, response: response}
}

// Unwrap はユースケースの出力を取得する
func (o *output[T]) Unwrap() (T, error) {
	return o.response, o.err
}

// Error はユースケースのエラーを取得する
func (o *output[T]) Error() error {
	return o.err
}

// Response はユースケースのレスポンスを取得する
func (o *output[T]) Response() T {
	return o.response
}
