package controller

const (
	MultiPartNextPartEoF   = "multipart: NextPart: EOF"
	RequestContentTypeIsnt = "request Content-Type isn't multipart/form-data"
	HttpNoSuchFile         = "http: no such file"
	ImageKey               = "icons"
)

type SuccessResponse struct {
	Result string `json:"result"`
}
type ErrorResponse struct {
	Error string `json:"error"`
}

func errorResponse(err error) ErrorResponse {
	return ErrorResponse{Error: err.Error()}
}
