package schemes

type SchemeResponses struct {
	StatusCode int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type SchemeResponsePaginated struct {
	StatusCode int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Count      int64       `json:"count"`
	Page       int         `json:"page"`
	TotalPage  int         `json:"totalPage"`
}
