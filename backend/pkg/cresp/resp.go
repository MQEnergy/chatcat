package cresp

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(data interface{}) *Response {
	return &Response{
		Code: 0,
		Msg:  "success",
		Data: data,
	}
}

func Fail(message string) *Response {
	return &Response{
		Code: -1,
		Msg:  message,
		Data: nil,
	}
}
