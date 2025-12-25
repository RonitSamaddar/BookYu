package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/RonitSamaddar/BookYu/internal/bookManager"
	"github.com/RonitSamaddar/BookYu/internal/consts"
)

func New(source Source) *Handler {
	return &Handler{
		source: source,
	}
}

func (h *Handler) RegisterRoutes() {
	http.HandleFunc(consts.SearchHandlerEndpoint, h.bookSearchHandler)
}

func (h *Handler) bookSearchHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %+v", r.URL.String())

	name := r.URL.Query().Get("name")
	language := r.URL.Query().Get("lang")
	if language == "" {
		language = "en"
	}
	log.Printf("Searching for book %s", name)
	book, err := h.source.GetBook(name, language)
	if err != nil {
		log.Printf("Error getting book: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	log.Printf("Book: %+v", book)

	bookResponse := bookManager.GetBookResponse(book)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	err = enc.Encode(bookResponse)
	if err != nil {
		log.Printf("Error marshaling book: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
