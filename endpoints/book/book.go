package book

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	"github.com/hiepdangnguyendai/example-go/domain"
	"github.com/hiepdangnguyendai/example-go/service"
)

type CreateData struct {
	Name        string      `json:"name"`
	Author      string      `json:"author"`
	Description string      `json:"Description"`
	Category_Id domain.UUID `json:"category_id"`
}
type CreateRequest struct {
	Book CreateData `json:"book"`
}
type CreateResponse struct {
	Book domain.Book `json:"book"`
}

func (CreateResponse) StatusCode() int {
	return http.StatusCreated
}
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req  = request.(CreateRequest)
			book = &domain.Book{
				Name:        req.Book.Name,
				Author:      req.Book.Author,
				Description: req.Book.Description,
				Category_Id: req.Book.Category_Id,
			}
		)
		err := s.BookService.Create(ctx, book)
		if err != nil {
			return nil, err
		}
		return CreateResponse{Book: *book}, nil
	}
}

type FindRequest struct {
	BookID domain.UUID
}
type FindResponse struct {
	Book *domain.Book `json:"book"`
}

func MakeFindEndPoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var bookFind domain.Book
		req := request.(FindRequest)
		bookFind.ID = req.BookID
		book, err := s.BookService.Find(ctx, &bookFind)
		if err != nil {
			return nil, err
		}
		return FindResponse{Book: book}, nil
	}
}

type FindAllRequest struct{}
type FindAllResponse struct {
	Books []domain.Book `json:"books"`
}

func MakeFindAllEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(FindAllRequest)
		books, err := s.BookService.FindAll(ctx)
		if err != nil {
			return nil, err
		}
		return FindAllResponse{Books: books}, nil
	}
}

type UpdateData struct {
	ID          domain.UUID `json:"-"`
	Name        string      `json:"name"`
	Author      string      `json:"author"`
	Description string      `json:"description"`
	Category_Id domain.UUID `json:"category_id"`
}
type UpdateRequest struct {
	Book UpdateData `json:"book"`
}
type UpdateResponse struct {
	Book domain.Book `json:"book"`
}

func MakeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req  = request.(UpdateRequest)
			book = domain.Book{
				Model:       domain.Model{ID: req.Book.ID},
				Name:        req.Book.Name,
				Author:      req.Book.Author,
				Description: req.Book.Description,
				Category_Id: req.Book.Category_Id,
			}
		)
		res, err := s.BookService.Update(ctx, &book)
		if err != nil {
			return nil, err
		}
		return UpdateResponse{Book: *res}, nil
	}
}

type DeleteRequest struct {
	BookID domain.UUID
}
type DeleteResponse struct {
	Status string `json:"status"`
}

func MakeDeleteEnpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			bookFind = domain.Book{}
			req      = request.(DeleteRequest)
		)
		bookFind.ID = req.BookID
		err := s.BookService.Delete(ctx, &bookFind)
		if err != nil {
			return nil, err
		}
		return DeleteResponse{"success"}, nil
	}
}
