package domain

type Category struct {
	Model
	Name  string `json:"name"`
	Books []Book `json:"-"`
}
