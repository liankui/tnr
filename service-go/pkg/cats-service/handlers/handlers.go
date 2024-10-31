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
	}
	return resp, nil
}

// UpdateCat implements Interface.
func (s catsServer) UpdateCat(ctx context.Context, in *pb.UpdateCatRequest) (*cat.Cat, error) {
	logs.Infow("UpdateCat", "request", in)

	if in.Cat.Id == "" {
		logs.Warnw("need to set the id")
		return nil, core.NewErrorFrom(400, "need to set the id")
	}

	if in.Cat.UpdateTime == nil {
		in.Cat.UpdateTime = core.Now()
	}

	if _, err := model.GetCatModel().Update(ctx, in.Cat); err != nil {
		logs.Warnw("failed to update cat model", "error", err)
		return nil, core.NewErrorFrom(500, fmt.Sprintf("update cat model error: %v", err))
	}

	resp := &cat.Cat{
		Id: in.Cat.Id,
	}
	return resp, nil
}

// GetCat implements Interface.
func (s catsServer) GetCat(ctx context.Context, in *pb.GetCatRequest) (*cat.Cat, error) {
	logs.Infow("GetCat", "request", in)

	if in.Id == "" {
		logs.Warnw("need to set the id")
		return nil, core.NewErrorFrom(400, "need to set the id")
	}

	c, err := model.GetCatModel().Get(ctx, in.Id)
	if err != nil {
		logs.Warnw("failed to get the cat", "id", in.Id, "error", err)
		return nil, core.NewErrorFrom(500, fmt.Sprintf("not found the cat, error: %v", err))
	}

	resp := &cat.Cat{
		Id:         c.Id,
		Name:       c.Name,
		NickName:   c.NickName,
		Sex:        c.Sex,
		Location:   c.Location,
		Area:       c.Area,
		Status:     c.Status,
		State:      c.State,
		CreateTime: c.CreateTime,
		UpdateTime: c.UpdateTime,
		// DeleteTime:
	}
	return resp, nil
}

// ListCats implements Interface.
func (s catsServer) ListCats(ctx context.Context, in *pb.ListCatsRequest) (*pb.ListCatsResponse, error) {
	logs.Infow("ListCats", "request", in)

	cats, err := model.GetCatModel().List(ctx, "")
	if err != nil {
		logs.Warnw("failed to get the cats", "error", err)
		return nil, core.NewErrorFrom(500, fmt.Sprintf("not found the cats, error: %v", err))
	}

	resp := &pb.ListCatsResponse{
		Cats: cats,
		// TotalCount:
		// NextPageToken:
	}
	return resp, nil
}

// BatchGetCat implements Interface.
func (s catsServer) BatchGetCat(ctx context.Context, in *pb.BatchGetCatRequest) (*pb.BatchGetCatResponse, error) {
	logs.Infow("BatchGetCat", "request", in)

	if len(in.Ids) == 0 {
		logs.Warnw("need to set the ids")
		return nil, core.NewErrorFrom(400, "need to set the ids")
	}

	cats, err := model.GetCatModel().BatchGet(ctx, in.Ids...)
	if err != nil {
		logs.Warnw("failed to get the cats from ids", "ids", in.Ids, "error", err)
		return nil, core.NewErrorFrom(500, fmt.Sprintf("not found the cats, error: %v", err))
	}

	resp := &pb.BatchGetCatResponse{
		Cats: cats,
	}
	return resp, nil
}

// DeleteCat implements Interface.
func (s catsServer) DeleteCat(ctx context.Context, in *pb.DeleteCatRequest) (*core.Null, error) {
	logs.Infow("DeleteCat", "request", in)

	if in.Id == "" {
		logs.Warnw("need to set the id")
		return nil, core.NewErrorFrom(400, "need to set the id")
	}

	if _, err := model.GetCatModel().Delete(ctx, in.Id); err != nil {
		logs.Warnw("failed to delete the cat", "id", in.Id, "error", err)
		return nil, core.NewErrorFrom(500, fmt.Sprintf("not delete the cat, error: %v", err))
	}

	resp := &core.Null{}
	return resp, nil
}
