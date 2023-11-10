package model

type Diagnostic struct {
	ID    string `json:"id" validate:"required"`
	Title string `json:"title" validate:"required"`
}
