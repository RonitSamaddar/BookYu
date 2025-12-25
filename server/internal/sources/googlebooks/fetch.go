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
	query = fmt.Sprintf("intitle:%s", query)
	return query
}

func GetBook(query, language, apiKey string, limit int) (types.Book, error) {
	url := formURL(query, language, apiKey, limit)
	log.Printf("url: %s", url)
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

	filteredItems := filterVolumes(response.Items, language)
	if len(filteredItems) == 0 {
		return types.Book{}, fmt.Errorf("no volumes found for query %s, language %s", query, language)
	}

	book, err := processVolumeInfo(filteredItems[0].VolumeInfo)
	if err != nil {
		return types.Book{}, err
	}
	return book, nil
}

func formURL(query, language, apiKey string, limit int) string {
	url := consts.GoogleBooksBaseURL
	url += "?" + consts.GoogleBooksApiKeyField + "=" + apiKey
	url += "&" + consts.GoogleBooksQueryField + "=" + query
	url += "&" + consts.GoogleBooksLimitField + "=" + fmt.Sprintf("%d", limit)
	url += "&" + consts.GoogleBooksLangField + "=" + language
	url += "&" + consts.GoogleBooksTypeField + "=" + consts.GoogleBooksTypeBooks
	return url
}
