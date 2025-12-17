package googlebooks

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/RonitSamaddar/BookYu/internal/consts"
	"github.com/RonitSamaddar/BookYu/internal/types"
)

func ProcessNameQuery(name string) string {
	trimmedName := strings.TrimSpace(strings.ToLower(name))
	nameTokens := strings.Split(trimmedName, " ")
	query := strings.Join(nameTokens, "+")
	return query
}

func GetBook(query, apiKey string, limit int) (types.Book, error) {
	url := fmt.Sprintf("%s?key=%s&%s=%s&%s=%d", consts.GoogleBooksBaseURL, apiKey, consts.GoogleBooksQueryField, query, consts.GoogleBooksLimitField, limit)
	resp, err := http.Get(url)
	if err != nil {
		return types.Book{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.Book{}, err
	}
	log.Printf("bodyBytes: %s", string(bodyBytes))

	var response GoogleBooksGetVolumeResponse
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return types.Book{}, err
	}

	log.Printf("response: %+v", response)

	book, err := processVolumeInfo(response.Items[0].VolumeInfo)
	if err != nil {
		return types.Book{}, err
	}
	return book, nil
}
