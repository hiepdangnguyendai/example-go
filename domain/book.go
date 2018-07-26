package domain

type Book struct {
	Model
	Name        string   `json:"name"`
	Author      string   `json:"author"`
	Description string   `json:"description"`
	Category_Id UUID     `json:"category_id"`
	Category    Category `json:"-"`
}
