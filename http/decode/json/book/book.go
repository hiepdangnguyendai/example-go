package book

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/hiepdangnguyendai/example-go/domain"
	bookEndpoint "github.com/hiepdangnguyendai/example-go/endpoints/book"
)

func FindRequest(_ context.Context, r *http.Request) (interface{}, error) {
	bookID, err := domain.UUIDFromString(chi.URLParam(r, "book_id"))
	if err != nil {
		return nil, err
	}
	return bookEndpoint.FindRequest{BookID: bookID}, nil
}
func FindAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return bookEndpoint.FindAllRequest{}, nil
}
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req bookEndpoint.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
func UpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	bookID, err := domain.UUIDFromString(chi.URLParam(r, "book_id"))
	if err != nil {
		return nil, err
	}
	var req bookEndpoint.UpdateRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	req.Book.ID = bookID
	return req, nil
}
func DeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	bookID, err := domain.UUIDFromString(chi.URLParam(r, "book_id"))
	if err != nil {
		return nil, err
	}
	return bookEndpoint.DeleteRequest{BookID: bookID}, nil
}
