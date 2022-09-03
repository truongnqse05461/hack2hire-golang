package dtos

type Response struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"`
}

type Meta struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
