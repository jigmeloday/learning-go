package products

import (
	"log"
	"net/http"

	"github.com/sikozonpc/ecom/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProductHandler(w http.ResponseWriter, r *http.Request) {
	err := h.service.ListProducts(r.Context())

	if err != nil {
		log.Printf("error %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	products := []string{"Hello", "go"}
	json.Write(w, http.StatusOK, products)
}
