package service

import (
	"github.com/hiepdangnguyendai/example-go/service/book"
	"github.com/hiepdangnguyendai/example-go/service/category"
	"github.com/hiepdangnguyendai/example-go/service/lend"
	"github.com/hiepdangnguyendai/example-go/service/user"
)

// Service define list of all services in projects
type Service struct {
	UserService     user.Service
	CategoryService category.Service
	BookService     book.Service
	LendService     lend.Service
}
