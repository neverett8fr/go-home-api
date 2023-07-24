package entities

type ReqIn struct {
	// temporary, have dedicatd struct in future
	Method    string `json:"method"`
	Route     string `json:"route"`
	Condition string `json:"condition"`
}

type Response struct {
	Errors []error     `json:"errors"`
	Data   interface{} `json:"data"`
}

func NewResponse(data interface{}, err ...error) *Response {

	return &Response{
		Errors: err,
		Data:   data,
	}
}
