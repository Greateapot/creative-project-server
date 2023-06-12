package models

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Затычка для странной логики
func Err() (err *Error) {
	err = &Error{0, "Err...?"}
	return
}

// Нет доступа (lip-lock)
func ErrAccessDenied() (err *Error) {
	err = &Error{1, "Access denied"}
	return
}

// Нет нужного значения в теле запроса
func ErrNoRequiredFieldInRequestQuery() (err *Error) {
	err = &Error{2, "No required field in request query"}
	return
}

// Объект уже существует
func ErrItemAlreadyExists() (err *Error) {
	err = &Error{3, "Item already exists"}
	return
}

// Объект не существует
func ErrItemNotExists() (err *Error) {
	err = &Error{4, "Item not exists"}
	return
}

// Не удалось отправить файл
func ErrSendFile() (err *Error) {
	err = &Error{5, "Cant't send file"}
	return
}

// WIP
func ErrWIP() (err *Error) {
	err = &Error{6, "WIP"}
	return
}
