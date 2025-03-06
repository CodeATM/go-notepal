package requests

type CreateNoteRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required,min=6"`
}
