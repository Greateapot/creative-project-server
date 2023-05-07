package models

/*
0x0{N} (*):

1 - access denied

0xA{N} (ADD):

1 - no path

2 - no type

3 - err type

4 - no title // B1, D1

5 - title already exists


0xB{N} (GET):

1 - no title // A4, D1

2 - not found // B3

3 - not found // B2


0xC{N} (CONFIG):

1 - err scanDelay

2 - no changes


0xD{N} (DEL):

1 - no title // A4, B1
*/
type Err struct {
	Code    byte   `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Ok   bool `json:"ok"`
	Err  *Err  `json:"err,omitempty"`
	Data any  `json:"data,omitempty"`
}

func CreateOkResponse() *Response {
	r := &Response{}
	r.Ok = true
	return r
}

func CreateErrResponse(code byte, message string) *Response {
	r := &Response{}
	r.Ok = false
	r.Err = &Err{code, message}
	return r
}

func CreateDataResponse(data any) *Response {
	r := &Response{}
	r.Ok = true
	r.Data = data
	return r
}
