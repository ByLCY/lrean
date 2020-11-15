package transport

import (
	"context"
	"encoding/json"
	"microservice/v1/service"
	"net/http"
)

// DecodeUppercaseRequest 对请求进行JSON编码
func DecodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request service.UppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// DecodeCountRequest 对请求进行JSON解码
func DecodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request service.CountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// EncodeResponse 对响应编码
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
