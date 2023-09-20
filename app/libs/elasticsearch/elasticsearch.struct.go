package elasticsearchLib

// CustomError represents a custom Elasticsearch error.
type CustomError struct {
    StatusCode int
    ErrorEs      map[string]interface{}
}