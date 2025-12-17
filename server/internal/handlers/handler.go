package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func New(source Source) *Handler {
	return &Handler{
		source: source,
	}
}

func (h *Handler) RegisterRoutes() {
	http.HandleFunc("/search", h.bookSearchHandler)
}

func (h *Handler) bookSearchHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %+v", r.URL.String())

	name := r.URL.Query().Get("name")
	log.Printf("Searching for book %s", name)
	book, err := h.source.GetBook(name)
	if err != nil {
		log.Printf("Error getting book: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "Book: %+v", book)
}
