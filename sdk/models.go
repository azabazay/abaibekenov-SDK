package sdk

type Response struct {
	Docs []map[string]interface{} `json:"docs"`
	// Docs   []interface{} `json:"docs"`
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Page   int `json:"page"`
	Pages  int `json:"pages"`
}

type Book struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}
