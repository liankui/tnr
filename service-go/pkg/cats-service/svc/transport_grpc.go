// Code generated by chaosmojo. DO NOT EDIT.
// Rerunning chaosmojo will overwrite this file.
// Version: 1.0
// Version Date:

package svc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"

	"github.com/liankui/tnr/go/pkg/cat"
	"github.com/mojo-lang/core/go/pkg/mojo/core"

	// this service api
	pb "github.com/liankui/tnr/go/pkg/cat/v1"
)

var (
	_ = cat.Cat{}
	_ = core.Null{}
)

// MakeGRPCServer makes a set of endpoints available as a gRPC CatsServer.
func MakeGRPCServer(endpoints Endpoints, tracer stdopentracing.Tracer, logger log.Logger) pb.CatsServer {
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerBefore(metadataToContext),
		grpctransport.ServerErrorLogger(logger),
	}

	addTracerOption := func(methodName string) []grpctransport.ServerOption {
		if tracer != nil {
			return append(serverOptions, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, methodName, logger)))
		}
		return serverOptions
	}

	return &grpcServer{
		// Cats

		createCat: grpctransport.NewServer(
			endpoints.CreateCatEndpoint,
			DecodeGRPCCreateCatRequest,
			EncodeGRPCCreateCatResponse,
			addTracerOption("create_cat")...,
		//append(serverOptions, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "create_cat", logger)))...,
		),
		updateCat: grpctransport.NewServer(
			endpoints.UpdateCatEndpoint,
			DecodeGRPCUpdateCatRequest,
			EncodeGRPCUpdateCatResponse,
			addTracerOption("update_cat")...,
		//append(serverOptions, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "update_cat", logger)))...,
		),
		getCat: grpctransport.NewServer(
			endpoints.GetCatEndpoint,
			DecodeGRPCGetCatRequest,
			EncodeGRPCGetCatResponse,
			addTracerOption("get_cat")...,
		//append(serverOptions, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "get_cat", logger)))...,
		),
		listCats: grpctransport.NewServer(
			endpoints.ListCatsEndpoint,
			DecodeGRPCListCatsRequest,
			EncodeGRPCListCatsResponse,
			addTracerOption("list_cats")...,
		//append(serverOptions, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "list_cats", logger)))...,
		),
		batchGetCat: grpctransport.NewServer(
			endpoints.BatchGetCatEndpoint,
			DecodeGRPCBatchGetCatRequest,
			EncodeGRPCBatchGetCatResponse,
			addTracerOption("batch_get_cat")...,
		//append(serverOptions, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "batch_get_cat", logger)))...,
		),
		deleteCat: grpctransport.NewServer(
			endpoints.DeleteCatEndpoint,
			DecodeGRPCDeleteCatRequest,
			EncodeGRPCDeleteCatResponse,
			addTracerOption("delete_cat")...,
		//append(serverOptions, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "delete_cat", logger)))...,
		),
	}
}

// grpcServer implements the CatsServer interface
type grpcServer struct {
	pb.UnimplementedCatsServer

	createCat   grpctransport.Handler
	updateCat   grpctransport.Handler
	getCat      grpctransport.Handler
	listCats    grpctransport.Handler
	batchGetCat grpctransport.Handler
	deleteCat   grpctransport.Handler
}

// Methods for grpcServer to implement CatsServer interface

func (s *grpcServer) CreateCat(ctx context.Context, req *pb.CreateCatRequest) (*cat.Cat, error) {
	_, rep, err := s.createCat.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*cat.Cat), nil
}

func (s *grpcServer) UpdateCat(ctx context.Context, req *pb.UpdateCatRequest) (*cat.Cat, error) {
	_, rep, err := s.updateCat.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*cat.Cat), nil
}

func (s *grpcServer) GetCat(ctx context.Context, req *pb.GetCatRequest) (*cat.Cat, error) {
	_, rep, err := s.getCat.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*cat.Cat), nil
}

func (s *grpcServer) ListCats(ctx context.Context, req *pb.ListCatsRequest) (*pb.ListCatsResponse, error) {
	_, rep, err := s.listCats.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ListCatsResponse), nil
}

func (s *grpcServer) BatchGetCat(ctx context.Context, req *pb.BatchGetCatRequest) (*pb.BatchGetCatResponse, error) {
	_, rep, err := s.batchGetCat.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.BatchGetCatResponse), nil
}

func (s *grpcServer) DeleteCat(ctx context.Context, req *pb.DeleteCatRequest) (*core.Null, error) {
	_, rep, err := s.deleteCat.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*core.Null), nil
}

// Server Decode

// DecodeGRPCCreateCatRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC CreateCat request to a user-domain CreateCat request. Primarily useful in a server.
func DecodeGRPCCreateCatRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateCatRequest)
	return req, nil
}

// DecodeGRPCUpdateCatRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC UpdateCat request to a user-domain UpdateCat request. Primarily useful in a server.
func DecodeGRPCUpdateCatRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.UpdateCatRequest)
	return req, nil
}

// DecodeGRPCGetCatRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC GetCat request to a user-domain GetCat request. Primarily useful in a server.
func DecodeGRPCGetCatRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetCatRequest)
	return req, nil
}

// DecodeGRPCListCatsRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC ListCats request to a user-domain ListCats request. Primarily useful in a server.
func DecodeGRPCListCatsRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.ListCatsRequest)
	return req, nil
}

// DecodeGRPCBatchGetCatRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC BatchGetCat request to a user-domain BatchGetCat request. Primarily useful in a server.
func DecodeGRPCBatchGetCatRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.BatchGetCatRequest)
	return req, nil
}

// DecodeGRPCDeleteCatRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC DeleteCat request to a user-domain DeleteCat request. Primarily useful in a server.
func DecodeGRPCDeleteCatRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.DeleteCatRequest)
	return req, nil
}

// Server Encode

// EncodeGRPCCreateCatResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain CreateCat response to a gRPC CreateCat reply. Primarily useful in a server.
func EncodeGRPCCreateCatResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*cat.Cat)
	return resp, nil
}

// EncodeGRPCUpdateCatResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain UpdateCat response to a gRPC UpdateCat reply. Primarily useful in a server.
func EncodeGRPCUpdateCatResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*cat.Cat)
	return resp, nil
}

// EncodeGRPCGetCatResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain GetCat response to a gRPC GetCat reply. Primarily useful in a server.
func EncodeGRPCGetCatResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*cat.Cat)
	return resp, nil
}

// EncodeGRPCListCatsResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain ListCats response to a gRPC ListCats reply. Primarily useful in a server.
func EncodeGRPCListCatsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ListCatsResponse)
	return resp, nil
}

// EncodeGRPCBatchGetCatResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain BatchGetCat response to a gRPC BatchGetCat reply. Primarily useful in a server.
func EncodeGRPCBatchGetCatResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.BatchGetCatResponse)
	return resp, nil
}

// EncodeGRPCDeleteCatResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain DeleteCat response to a gRPC DeleteCat reply. Primarily useful in a server.
func EncodeGRPCDeleteCatResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*core.Null)
	return resp, nil
}

// Helpers

func metadataToContext(ctx context.Context, md metadata.MD) context.Context {
	for k, v := range md {
		if v != nil {
			// The key is added both in metadata format (k) which is all lower
			// and the http.CanonicalHeaderKey of the key so that it can be
			// accessed in either format
			ctx = context.WithValue(ctx, k, v[0])
			ctx = context.WithValue(ctx, http.CanonicalHeaderKey(k), v[0])
		}
	}

	return ctx
}