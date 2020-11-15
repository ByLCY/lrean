package main

import (
	"log"
	"net/http"

	"microservice/v1/endpoint"
	"microservice/v1/service"
	"microservice/v1/transport"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	svc := service.NewStringService()

	uppercaseHanlder := httptransport.NewServer(
		endpoint.MakeUppercaseEndpoint(svc),
		transport.DecodeUppercaseRequest,
		transport.EncodeResponse,
	)

	countHanlder := httptransport.NewServer(
		endpoint.MakeCountEndpoint(svc),
		transport.DecodeCountRequest,
		transport.EncodeResponse,
	)

	http.Handle("/uppercase", uppercaseHanlder)
	http.Handle("/count", countHanlder)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
