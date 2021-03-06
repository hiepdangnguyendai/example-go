package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/hiepdangnguyendai/example-go/endpoints"

	bookDecode "github.com/hiepdangnguyendai/example-go/http/decode/json/book"
	categoryDecode "github.com/hiepdangnguyendai/example-go/http/decode/json/category"
	lendDecode "github.com/hiepdangnguyendai/example-go/http/decode/json/lend"
	userDecode "github.com/hiepdangnguyendai/example-go/http/decode/json/user"
)

// NewHTTPHandler ...
func NewHTTPHandler(endpoints endpoints.Endpoints,
	logger log.Logger,
	useCORS bool) http.Handler {
	r := chi.NewRouter()

	// if running on local (using `make dev`), include cors middleware
	if useCORS {
		cors := cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			AllowCredentials: true,
		})
		r.Use(cors.Handler)
	}

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Get("/_warm", httptransport.NewServer(
		endpoint.Nop,
		httptransport.NopRequestDecoder,
		httptransport.EncodeJSONResponse,
		options...,
	).ServeHTTP)
	r.Route("/users", func(r chi.Router) {
		r.Get("/", httptransport.NewServer(
			endpoints.FindAllUser,
			userDecode.FindAllRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Get("/{user_id}", httptransport.NewServer(
			endpoints.FindUser,
			userDecode.FindRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Post("/", httptransport.NewServer(
			endpoints.CreateUser,
			userDecode.CreateRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Put("/{user_id}", httptransport.NewServer(
			endpoints.UpdateUser,
			userDecode.UpdateRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Delete("/{user_id}", httptransport.NewServer(
			endpoints.DeleteUser,
			userDecode.DeleteRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
	})
	r.Route("/categories", func(r chi.Router) {
		r.Get("/", httptransport.NewServer(
			endpoints.FindAllCategories,
			categoryDecode.FindAllRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Get("/{category_id}", httptransport.NewServer(
			endpoints.FindCategory,
			categoryDecode.FindRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Post("/", httptransport.NewServer(
			endpoints.CreateCategory,
			categoryDecode.CreateRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Put("/{category_id}", httptransport.NewServer(
			endpoints.UpdateCategory,
			categoryDecode.UpdateRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Delete("/{category_id}", httptransport.NewServer(
			endpoints.DeleteCategory,
			categoryDecode.DeleteRequest,
			encodeResponse,
			options...,
		).ServeHTTP)

	})
	r.Route("/books", func(r chi.Router) {
		r.Get("/", httptransport.NewServer(
			endpoints.FindAllBooks,
			bookDecode.FindAllRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Get("/{book_id}", httptransport.NewServer(
			endpoints.FindBook,
			bookDecode.FindRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Post("/", httptransport.NewServer(
			endpoints.CreateBook,
			bookDecode.CreateRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Put("/{book_id}", httptransport.NewServer(
			endpoints.UpdateBook,
			bookDecode.UpdateRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Delete("/{book_id}", httptransport.NewServer(
			endpoints.DeleteBook,
			bookDecode.DeleteRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
	})
	r.Route("/lends", func(r chi.Router) {
		r.Get("/", httptransport.NewServer(
			endpoints.FindAllLens,
			lendDecode.FindAllRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Get("/{lend_id}", httptransport.NewServer(
			endpoints.FindLend,
			lendDecode.FindRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Post("/", httptransport.NewServer(
			endpoints.CreateLend,
			lendDecode.CreateRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Put("/{lend_id}", httptransport.NewServer(
			endpoints.UpdateLend,
			lendDecode.UpdateRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Delete("/{lend_id}", httptransport.NewServer(
			endpoints.DeleteLend,
			lendDecode.DeleteRequest,
			encodeResponse,
			options...,
		).ServeHTTP)

	})

	return r
}
