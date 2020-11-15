package service

// UppercaseRequest Uppercase的请求结构体
type UppercaseRequest struct {
	S string `json:"s"`
}

// UppercaseResponse Uppercase的返回结构体
type UppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

// CountRequest Count的请求结构体
type CountRequest struct {
	S string `json:"s"`
}

// CountResponse Count的返回结构体
type CountResponse struct {
	V int `json:"v"`
}
