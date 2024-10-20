// Code generated by chaosmojo. DO NOT EDIT.
// Rerunning chaosmojo will overwrite this file.
// Version: 1.0
// Version Date:

package svc

// This file contains methods to make individual endpoints from services,
// request and response types to serve those endpoints, as well as encoders and
// decoders for those types, for all of our supported transport serialization
// formats.

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"

	"github.com/liankui/tnr/go/pkg/cat"
	"github.com/mojo-lang/core/go/pkg/mojo/core"

	// this service api
	pb "github.com/liankui/tnr/go/pkg/cat/v1"
)

var (
	_ = cat.Cat{}
	_ = core.Null{}
)

// Endpoints collects all of the endpoints that compose an add service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
//
// In a server, it's useful for functions that need to operate on a per-endpoint
// basis. For example, you might pass an Endpoints to a function that produces
// an http.Handler, with each method (endpoint) wired up to a specific path. (It
// is probably a mistake in design to invoke the Service methods on the
// Endpoints struct in a server.)
//
// In a client, it's useful to collect individually constructed endpoints into a
// single type that implements the Service interface. For example, you might
// construct individual endpoints using transport/http.NewClient, combine them into an Endpoints, and return it to the caller as a Service.
type Endpoints struct {
	CreateCatEndpoint   endpoint.Endpoint
	UpdateCatEndpoint   endpoint.Endpoint
	GetCatEndpoint      endpoint.Endpoint
	ListCatsEndpoint    endpoint.Endpoint
	BatchGetCatEndpoint endpoint.Endpoint
	DeleteCatEndpoint   endpoint.Endpoint
}

// Endpoints

func (e Endpoints) CreateCat(ctx context.Context, in *pb.CreateCatRequest) (*cat.Cat, error) {
	response, err := e.CreateCatEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*cat.Cat), nil
}

func (e Endpoints) UpdateCat(ctx context.Context, in *pb.UpdateCatRequest) (*cat.Cat, error) {
	response, err := e.UpdateCatEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*cat.Cat), nil
}

func (e Endpoints) GetCat(ctx context.Context, in *pb.GetCatRequest) (*cat.Cat, error) {
	response, err := e.GetCatEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*cat.Cat), nil
}

func (e Endpoints) ListCats(ctx context.Context, in *pb.ListCatsRequest) (*pb.ListCatsResponse, error) {
	response, err := e.ListCatsEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.ListCatsResponse), nil
}

func (e Endpoints) BatchGetCat(ctx context.Context, in *pb.BatchGetCatRequest) (*pb.BatchGetCatResponse, error) {
	response, err := e.BatchGetCatEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.BatchGetCatResponse), nil
}

func (e Endpoints) DeleteCat(ctx context.Context, in *pb.DeleteCatRequest) (*core.Null, error) {
	response, err := e.DeleteCatEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*core.Null), nil
}

// Make Endpoints

func MakeCreateCatEndpoint(s pb.CatsServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.CreateCatRequest)
		v, err := s.CreateCat(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeUpdateCatEndpoint(s pb.CatsServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.UpdateCatRequest)
		v, err := s.UpdateCat(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeGetCatEndpoint(s pb.CatsServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.GetCatRequest)
		v, err := s.GetCat(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeListCatsEndpoint(s pb.CatsServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.ListCatsRequest)
		v, err := s.ListCats(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeBatchGetCatEndpoint(s pb.CatsServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.BatchGetCatRequest)
		v, err := s.BatchGetCat(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeDeleteCatEndpoint(s pb.CatsServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.DeleteCatRequest)
		v, err := s.DeleteCat(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

// WrapAllExcept wraps each Endpoint field of struct Endpoints with a
// go-kit/kit/endpoint.Middleware.
// Use this for applying a set of middlewares to every endpoint in the service.
// Optionally, endpoints can be passed in by name to be excluded from being wrapped.
// WrapAllExcept(middleware, "Status", "Ping")
func (e *Endpoints) WrapAllExcept(middleware endpoint.Middleware, excluded ...string) {
	included := map[string]struct{}{
		"create_cat":    struct{}{},
		"update_cat":    struct{}{},
		"get_cat":       struct{}{},
		"list_cats":     struct{}{},
		"batch_get_cat": struct{}{},
		"delete_cat":    struct{}{},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc, _ := range included {
		if inc == "create_cat" {
			e.CreateCatEndpoint = middleware(e.CreateCatEndpoint)
		}
		if inc == "update_cat" {
			e.UpdateCatEndpoint = middleware(e.UpdateCatEndpoint)
		}
		if inc == "get_cat" {
			e.GetCatEndpoint = middleware(e.GetCatEndpoint)
		}
		if inc == "list_cats" {
			e.ListCatsEndpoint = middleware(e.ListCatsEndpoint)
		}
		if inc == "batch_get_cat" {
			e.BatchGetCatEndpoint = middleware(e.BatchGetCatEndpoint)
		}
		if inc == "delete_cat" {
			e.DeleteCatEndpoint = middleware(e.DeleteCatEndpoint)
		}
	}
}

// LabeledMiddleware will get passed the endpoint name when passed to
// WrapAllLabeledExcept, this can be used to write a generic metrics
// middleware which can send the endpoint name to the metrics collector.
type LabeledMiddleware func(string, endpoint.Endpoint) endpoint.Endpoint

// WrapAllLabeledExcept wraps each Endpoint field of struct Endpoints with a
// LabeledMiddleware, which will receive the name of the endpoint. See
// LabeldMiddleware. See method WrapAllExept for details on excluded
// functionality.
func (e *Endpoints) WrapAllLabeledExcept(middleware func(string, endpoint.Endpoint) endpoint.Endpoint, excluded ...string) {
	included := map[string]struct{}{
		"create_cat":    struct{}{},
		"update_cat":    struct{}{},
		"get_cat":       struct{}{},
		"list_cats":     struct{}{},
		"batch_get_cat": struct{}{},
		"delete_cat":    struct{}{},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc, _ := range included {
		if inc == "create_cat" {
			e.CreateCatEndpoint = middleware("create_cat", e.CreateCatEndpoint)
		}
		if inc == "update_cat" {
			e.UpdateCatEndpoint = middleware("update_cat", e.UpdateCatEndpoint)
		}
		if inc == "get_cat" {
			e.GetCatEndpoint = middleware("get_cat", e.GetCatEndpoint)
		}
		if inc == "list_cats" {
			e.ListCatsEndpoint = middleware("list_cats", e.ListCatsEndpoint)
		}
		if inc == "batch_get_cat" {
			e.BatchGetCatEndpoint = middleware("batch_get_cat", e.BatchGetCatEndpoint)
		}
		if inc == "delete_cat" {
			e.DeleteCatEndpoint = middleware("delete_cat", e.DeleteCatEndpoint)
		}
	}
}
