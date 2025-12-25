package consts

// Consts for configs
const (
	GoogleBooksApiKeyEnvVar = "GOOGLE_BOOKS_API_KEY"
)

// Consts for handlers
const (
	SearchHandlerEndpoint = "/search"
)

// Consts for server
const (
	ServerPort = "8080"
)

// Consts for sources
const (
	SourceNameGoogleBooks  = "GoogleBooks"
	GoogleBooksBaseURL     = "https://www.googleapis.com/books/v1/volumes"
	GoogleBooksQueryField  = "q"
	GoogleBooksLangField   = "langRestrict"
	GoogleBooksTypeField   = "type"
	GoogleBooksApiKeyField = "key"
	GoogleBooksTypeBooks   = "books"
	GoogleBooksLimitField  = "maxResults"
	GoogleBooksIsbn13Type  = "ISBN_13"
	GoogleBooksIsbn10Type  = "ISBN_10"
)
