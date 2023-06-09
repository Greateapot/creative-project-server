package models

type Response struct {
	Ok      bool `json:"ok"`
	ErrCode byte `json:"err,omitempty"`
	Data    any  `json:"data,omitempty"`
}

// Все ок, вот тебе данные
func CreateDataResponse(data any) *Response {
	r := &Response{}
	r.Ok = true
	r.Data = data
	return r
}

// Все ок, но данные ты не получишь
func GetResponseOK() *Response {
	return &Response{true, 0, nil}
}

// Затычка для странной логики
func GetResponseErr() *Response {
	return &Response{false, 0, nil}
}

// Нет доступа (lip-lock)
func GetResponseErrAccessDenied() *Response {
	return &Response{false, 1, nil}
}

// Нет нужного значения в теле запроса
func GetResponseErrNoValueInBody() *Response {
	return &Response{false, 2, nil}
}

// Объект уже существует
func GetResponseErrItemAlreadyExists() *Response {
	return &Response{false, 3, nil}
}

// Объект не существует
func GetResponseErrItemNotExists() *Response {
	return &Response{false, 4, nil}
}

// Не удалось отправить файл
func GetResponseErrSendFile() *Response {
	return &Response{false, 5, nil}
}

// WIP
func GetResponseErrWIP() *Response {
	return &Response{false, 99, nil}
}
