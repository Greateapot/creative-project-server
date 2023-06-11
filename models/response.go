package models

type Response struct {
	Err  any `json:"err,omitempty"`
	Data any `json:"data,omitempty"`
}

// Все ок, но данные ты не получишь
func GetResponseOK() (r *Response) {
	r = &Response{nil, nil}
	return
}

// Все ок, вот тебе данные
func GetResponseData(data any) (r *Response) {
	r = GetResponseOK()
	r.Data = data
	return
}

// Затычка для странной логики
func GetResponseErr() (r *Response) {
	r = GetResponseOK()
	r.Err = "Err...?"
	return
}

// Нет доступа (lip-lock)
func GetResponseErrAccessDenied() (r *Response) {
	r = GetResponseOK()
	r.Err = "Access denied"
	return
}

// Нет нужного значения в теле запроса
func GetResponseErrNoFieldInQuery() (r *Response) {
	r = GetResponseOK()
	r.Err = "No required field in request query"
	return
}

// Объект уже существует
func GetResponseErrItemAlreadyExists() (r *Response) {
	r = GetResponseOK()
	r.Err = "Item already exists"
	return
}

// Объект не существует
func GetResponseErrItemNotExists() (r *Response) {
	r = GetResponseOK()
	r.Err = "Item not exists"
	return
}

// Не удалось отправить файл
func GetResponseErrSendFile() (r *Response) {
	r = GetResponseOK()
	r.Err = "Cant't send file"
	return
}

// WIP
func GetResponseErrWIP() (r *Response) {
	r = GetResponseOK()
	r.Err = "WIP"
	return
}
