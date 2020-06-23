package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	storeHandler := httptransport.NewServer(
		// Endpoint.
		func(_ context.Context, request interface{}) (interface{}, error) {
			product := request.(Product)

			// Create product.
			fmt.Println("create product!")

			return product, nil
		},

		// Decoder.
		func(_ context.Context, r *http.Request) (interface{}, error) {
			var product Product
			if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
				return nil, err
			}
			return product, nil
		},

		// Encoder.
		encodeResponse,
	)

	r := chi.NewRouter()
	r.Post("/products", storeHandler.ServeHTTP)
	http.ListenAndServe(":8080", r)
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
