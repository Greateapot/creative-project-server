package models

/*
ErrCodes:

0 - no err

1 - access denied

2 - err parse req body

3 - no path in req body

4 - no title in req body

5 - no type in req body

6 - path not found

7 - title not found

8 - type not found

9 - title already exists

10 - path not exists

11 - WIP
*/
type Response struct {
	Ok      bool `json:"ok"`
	ErrCode byte `json:"err,omitempty"`
	Data    any  `json:"data,omitempty"`
}

func CreateOkResponse() *Response {
	r := &Response{}
	r.Ok = true
	return r
}

func CreateErrResponse(code byte) *Response {
	r := &Response{}
	r.Ok = false
	r.ErrCode = code
	return r
}

func CreateDataResponse(data any) *Response {
	r := &Response{}
	r.Ok = true
	r.Data = data
	return r
}
