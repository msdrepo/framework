package responses

type SuccessResponse struct {
	Status     string          `json:"status"`
	StatusCode int             `json:"status_code"`
	Message    string          `json:"message"`
	Data       interface{}     `json:"data"`
	Meta       *PaginationMeta `json:"meta,omitempty"`
	Token      interface{}     `json:"token"`
}

type ErrorResponse struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	TotalPages  int   `json:"total_pages"`
	TotalItems  int64 `json:"total_items"`
}
