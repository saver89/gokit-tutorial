package transports

import (
	"context"
	"encoding/json"
	"gokit-tutorial/dto"
	"gokit-tutorial/service"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func MakeUppercaseEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.UppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return dto.UppercaseResponse{V: v, Err: err.Error()}, nil
		}
		return dto.UppercaseResponse{V: v, Err: ""}, nil
	}
}

func MakeCountEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(dto.CountRequest)
		v := svc.Count(req.S)
		return dto.CountResponse{V: v}, nil
	}
}

func DecodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request dto.UppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request dto.CountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
