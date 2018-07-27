package lend

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"

	"github.com/hiepdangnguyendai/example-go/domain"
	"github.com/hiepdangnguyendai/example-go/service"
)

type CreateData struct {
	UserID domain.UUID `json:"user_id"`
	BooKID domain.UUID `json:"book_id"`
	To     time.Time   `json:"to"`
}
type CreateRequest struct {
	Lend CreateData `json:"lend"`
}
type CreateResponse struct {
	Lend domain.Lend `json:"lend"`
}

func (CreateResponse) StatusCode() int {
	return http.StatusCreated
}
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req  = request.(CreateRequest)
			lend = &domain.Lend{
				UserID: req.Lend.UserID,
				BookID: req.Lend.BooKID,
				To:     req.Lend.To,
			}
		)
		err := s.LendService.Create(ctx, lend)
		if err != nil {
			return nil, err
		}
		return CreateResponse{Lend: *lend}, nil
	}
}

type FindRequest struct {
	LendID domain.UUID
}
type FindResponse struct {
	Lend *domain.Lend `json:"lend"`
}

func MakeFindEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var lend_ domain.Lend
		req := request.(FindRequest)
		lend_.ID = req.LendID
		lend, err := s.LendService.Find(ctx, &lend_)
		if err != nil {
			return nil, err
		}
		return FindResponse{Lend: lend}, nil
	}
}

type FindAllRequest struct{}
type FindAllResponse struct {
	Lends []domain.Lend `json:"lends"`
}

func MakeFindAllEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(FindAllRequest)
		lends, err := s.LendService.FindAll(ctx)
		if err != nil {
			return nil, err
		}
		return FindAllResponse{Lends: lends}, nil
	}
}

type UpdateData struct {
	ID     domain.UUID `json:"-"`
	UserID domain.UUID `json:"user_id"`
	BookID domain.UUID `json:"book_id"`
	From   time.Time   `json:"from"`
	To     time.Time   `json:"to"`
}
type UpdateRequest struct {
	Lens UpdateData `json:"lend"`
}
type UpdateResponse struct {
	Lend domain.Lend `json:"lend"`
}

func MakeUpdateEnpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req  = request.(UpdateRequest)
			lend = domain.Lend{
				Model:  domain.Model{ID: req.Lens.ID},
				UserID: req.Lens.UserID,
				BookID: req.Lens.BookID,
				From:   req.Lens.From,
				To:     req.Lens.To,
			}
		)
		res, err := s.LendService.Update(ctx, &lend)
		if err != nil {
			return nil, err
		}
		return UpdateResponse{Lend: *res}, nil
	}
}

type DeleteRequest struct {
	LendID domain.UUID
}
type DeleteResponse struct {
	Status string `json:"status"`
}

func MakeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			lendFind = domain.Lend{}
			req      = request.(DeleteRequest)
		)

		lendFind.ID = req.LendID
		err := s.LendService.Delete(ctx, &lendFind)
		if err != nil {
			return nil, err
		}
		return DeleteResponse{"success"}, nil

	}
}
