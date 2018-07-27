package category

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/hiepdangnguyendai/example-go/domain"
	categoryEndpoint "github.com/hiepdangnguyendai/example-go/endpoints/category"
)

func FindRequest(_ context.Context, r *http.Request) (interface{}, error) {
	categoryID, err := domain.UUIDFromString(chi.URLParam(r, "category_id"))
	if err != nil {
		return nil, err
	}
	return categoryEndpoint.FindRequest{CategoryID: categoryID}, nil
}
func FindAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return categoryEndpoint.FindAllRequest{}, nil
}
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req categoryEndpoint.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
func UpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	categoryID, err := domain.UUIDFromString(chi.URLParam(r, "category_id"))
	if err != nil {
		return nil, err
	}
	var req categoryEndpoint.UpdateRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	req.Category.ID = categoryID
	return req, nil
}
func DeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	categoryID, err := domain.UUIDFromString(chi.URLParam(r, "category_id"))
	if err != nil {
		return nil, err
	}
	return categoryEndpoint.DeleteRequest{CategoryID: categoryID}, nil
}
