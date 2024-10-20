package handlers

import (
	"context"
	"fmt"

	"github.com/chaos-io/chaos/logs"

	"github.com/mojo-lang/core/go/pkg/mojo/core"

	"github.com/liankui/tnr/go/pkg/cat"
	"github.com/liankui/tnr/service-go/pkg/model"

	// this service api
	pb "github.com/liankui/tnr/go/pkg/cat/v1"
)

var (
	_ = cat.Cat{}
	_ = core.Null{}
)

func init() {
	model.Init()
}

type catsServer struct {
	pb.UnimplementedCatsServer
}

// NewService returns a naive, stateless implementation of Interface.
func NewService() pb.CatsServer {
	return catsServer{}
}

// CreateCat implements Interface.
func (s catsServer) CreateCat(ctx context.Context, in *pb.CreateCatRequest) (*cat.Cat, error) {
	logs.Infow("CreateCat", "request", in)

	if in.Cat.Name == "" {
		logs.Warnw("need to set the name")
		return nil, core.NewErrorFrom(400, fmt.Sprintf("need to set the name"))
	}

	if _, err := model.GetCatModel().Create(ctx, in.GetCat()); err != nil {
		logs.Warnw("failed to create cat in db", "error", err)
		return nil, err
	}

	resp := &cat.Cat{
		Id: in.Cat.Id,
		// Name:
		// NickName:
		// Sex:
		// Location:
		// Area:
		// Status:
		// State:
		// CreateTime:
		// UpdateTime:
		// DeleteTime:
	}
	return resp, nil
}

// UpdateCat implements Interface.
func (s catsServer) UpdateCat(ctx context.Context, in *pb.UpdateCatRequest) (*cat.Cat, error) {
	logs.Infow("UpdateCat", "request", in)

	resp := &cat.Cat{
		// Id:
		// Name:
		// NickName:
		// Sex:
		// Location:
		// Area:
		// Status:
		// State:
		// CreateTime:
		// UpdateTime:
		// DeleteTime:
	}
	return resp, nil
}

// GetCat implements Interface.
func (s catsServer) GetCat(ctx context.Context, in *pb.GetCatRequest) (*cat.Cat, error) {
	logs.Infow("GetCat", "request", in)

	resp := &cat.Cat{
		// Id:
		// Name:
		// NickName:
		// Sex:
		// Location:
		// Area:
		// Status:
		// State:
		// CreateTime:
		// UpdateTime:
		// DeleteTime:
	}
	return resp, nil
}

// ListCats implements Interface.
func (s catsServer) ListCats(ctx context.Context, in *pb.ListCatsRequest) (*pb.ListCatsResponse, error) {
	logs.Infow("ListCats", "request", in)

	resp := &pb.ListCatsResponse{
		// Cats:
		// TotalCount:
		// NextPageToken:
	}
	return resp, nil
}

// BatchGetCat implements Interface.
func (s catsServer) BatchGetCat(ctx context.Context, in *pb.BatchGetCatRequest) (*pb.BatchGetCatResponse, error) {
	logs.Infow("BatchGetCat", "request", in)

	resp := &pb.BatchGetCatResponse{
		// Cats:
	}
	return resp, nil
}

// DeleteCat implements Interface.
func (s catsServer) DeleteCat(ctx context.Context, in *pb.DeleteCatRequest) (*core.Null, error) {
	logs.Infow("DeleteCat", "request", in)

	resp := &core.Null{}
	return resp, nil
}
