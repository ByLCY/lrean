package endpoint

import (
	"context"

	"microservice/v1/service"

	"github.com/go-kit/kit/endpoint"
)

// MakeUppercaseEndpoint 生成Uppercase的处理函数
func MakeUppercaseEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(service.UppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return service.UppercaseResponse{V: "", Err: err.Error()}, nil
		}
		return service.UppercaseResponse{V: v, Err: ""}, nil
	}
}

// MakeCountEndpoint 生成Count的处理函数
func MakeCountEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(service.CountRequest)
		v := svc.Count(req.S)
		return service.CountResponse{V: v}, nil
	}
}
