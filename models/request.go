package models

type Request struct {
	Path  string `json:"path,omitempty"`
	Title string `json:"title,omitempty"`
	Type  int    `json:"type,omitempty"`
}

func NewRequest() (r *Request) {
	r = &Request{"", "", -1}
	return
}
