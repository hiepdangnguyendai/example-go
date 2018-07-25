package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/hiepdangnguyendai/example-go/service"

	"github.com/hiepdangnguyendai/example-go/endpoints/book"
	"github.com/hiepdangnguyendai/example-go/endpoints/category"
	"github.com/hiepdangnguyendai/example-go/endpoints/user"
)

// Endpoints .
type Endpoints struct {
	// User Endpoint
	FindUser    endpoint.Endpoint
	FindAllUser endpoint.Endpoint
	CreateUser  endpoint.Endpoint
	UpdateUser  endpoint.Endpoint
	DeleteUser  endpoint.Endpoint
	// Category Endpoint
	FindCategory      endpoint.Endpoint
	FindAllCategories endpoint.Endpoint
	CreateCategory    endpoint.Endpoint
	UpdateCategory    endpoint.Endpoint
	DeleteCategory    endpoint.Endpoint
	// Book Endpoint
	FindBook     endpoint.Endpoint
	FindAllBooks endpoint.Endpoint
	CreateBook   endpoint.Endpoint
	UpdateBook   endpoint.Endpoint
	DeleteBook   endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct
func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{
		// User Endpoint
		FindUser:    user.MakeFindEndPoint(s),
		FindAllUser: user.MakeFindAllEndpoint(s),
		CreateUser:  user.MakeCreateEndpoint(s),
		UpdateUser:  user.MakeUpdateEndpoint(s),
		DeleteUser:  user.MakeDeleteEndpoint(s),
		// Category Endpoint
		FindCategory:      category.MakeFindEndPoint(s),
		FindAllCategories: category.MakeFindAllEndpoint(s),
		CreateCategory:    category.MakeCreateEndpoint(s),
		UpdateCategory:    category.MakeUpdateEndpoint(s),
		DeleteCategory:    category.MakeDeleteEnpoint(s),
		// Book Enpoint
		FindBook:     book.MakeFindEndPoint(s),
		FindAllBooks: book.MakeFindAllEndpoint(s),
		CreateBook:   book.MakeCreateEndpoint(s),
		UpdateBook:   book.MakeUpdateEndpoint(s),
		DeleteBook:   book.MakeDeleteEnpoint(s),
	}
}
