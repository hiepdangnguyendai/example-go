package category

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	"github.com/hiepdangnguyendai/example-go/domain"
	"github.com/hiepdangnguyendai/example-go/service"
)

type CreateData struct {
	Name string `json:"name"`
}
type CreateRequest struct {
	Category CreateData `json:"category"`
}
type CreateResponse struct {
	Category domain.Category `json:"category"`
}

func (CreateResponse) StatusCode() int {
	return http.StatusCreated
}
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req      = request.(CreateRequest)
			category = &domain.Category{
				Name: req.Category.Name,
			}
		)
		err := s.CategoryService.Create(ctx, category)
		if err != nil {
			return nil, err
		}
		return CreateResponse{Category: *category}, nil
	}
}

type FindRequest struct {
	CategoryID domain.UUID
}
type FindResponse struct {
	Category *domain.Category `json:"category"`
}

func MakeFindEndPoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var categoryFind domain.Category
		req := request.(FindRequest)
		categoryFind.ID = req.CategoryID
		category, err := s.CategoryService.Find(ctx, &categoryFind)
		if err != nil {
			return nil, err
		}
		return FindResponse{Category: category}, nil
	}
}

type FindAllRequest struct{}
type FindAllResponse struct {
	Categories []domain.Category `json:"categories"`
}

func MakeFindAllEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(FindAllRequest)
		categories, err := s.CategoryService.FindAll(ctx)
		if err != nil {
			return nil, err
		}
		return FindAllResponse{Categories: categories}, nil
	}
}

type UpdateData struct {
	ID   domain.UUID `json:"-"`
	Name string      `json"name"`
}
type UpdateRequest struct {
	Category UpdateData `json:"category"`
}
type UpdateResponse struct {
	Category domain.Category `json:"category"`
}

func MakeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req      = request.(UpdateRequest)
			category = domain.Category{
				Model: domain.Model{ID: req.Category.ID},
				Name:  req.Category.Name,
			}
		)
		res, err := s.CategoryService.Update(ctx, &category)
		if err != nil {
			return nil, err
		}
		return UpdateResponse{Category: *res}, nil
	}
}

type DeleteRequest struct {
	CategoryID domain.UUID
}
type DeleteResponse struct {
	Status string `json:"status"`
}

func MakeDeleteEnpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			categoryFind = domain.Category{}
			req          = request.(DeleteRequest)
		)
		categoryFind.ID = req.CategoryID
		err := s.CategoryService.Delete(ctx, &categoryFind)
		if err != nil {
			return nil, err
		}
		return DeleteResponse{"success"}, nil
	}
}
