package model

import (
	"context"
	"sync"

	"github.com/chaos-io/chaos/db"
	"github.com/chaos-io/chaos/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"

	"github.com/segmentio/ksuid"

	"github.com/liankui/tnr/go/pkg/cat"
)

var catModel *CatModel
var catModelOnce sync.Once

type CatModel struct {
	DB *db.DB
}

func GetCatModel() *CatModel {
	catModelOnce.Do(func() {
		catModel = NewCatModel()
	})

	return catModel
}

func NewCatModel() *CatModel {
	m := &CatModel{DB: initDB()}
	if !m.DB.Config.DisableAutoMigrate || !d.Migrator().HasTable(&cat.Cat{}) {
		if err := d.AutoMigrate(&cat.Cat{}); err != nil {
			logs.Error("Init CatModel model err: ", err)
			panic(err)
		}
	}

	return m
}

func (m *CatModel) Create(ctx context.Context, cat *cat.Cat) (int64, error) {
	if cat.Id == "" {
		cat.Id = ksuid.New().String()
	}

	now := core.Now()
	if cat.CreateTime == nil {
		cat.CreateTime = now
		cat.UpdateTime = now
	}

	result := m.DB.WithContext(ctx).Create(cat)
	return result.RowsAffected, result.Error
}

func (m *CatModel) Get(ctx context.Context, id string) (*cat.Cat, error) {
	cat := &cat.Cat{}
	return cat, m.DB.WithContext(ctx).First(cat, "id = ?", id).Error
}

func (m *CatModel) Delete(ctx context.Context, id string) (int64, error) {
	result := m.DB.WithContext(ctx).Where("id = ?", id).Delete(&cat.Cat{})
	return result.RowsAffected, result.Error
}

func (m *CatModel) Update(ctx context.Context, cat *cat.Cat) (int64, error) {
	result := m.DB.WithContext(ctx).Updates(cat)
	return result.RowsAffected, result.Error
}

func (m *CatModel) List(ctx context.Context, filter string, condition ...string) ([]*cat.Cat, error) {
	var cat []*cat.Cat

	tx := m.DB.WithContext(ctx)
	// todo add condition

	return cat, tx.Find(&cat).Error
}

func (m *CatModel) BatchCreate(ctx context.Context, cat ...*cat.Cat) (int64, error) {
	result := m.DB.WithContext(ctx).CreateInBatches(cat, len(cat))
	return result.RowsAffected, result.Error
}

func (m *CatModel) BatchGet(ctx context.Context, ids ...string) ([]*cat.Cat, error) {
	var cat []*cat.Cat
	return cat, m.DB.WithContext(ctx).Find(&cat, "id = ?", ids).Error
}

func (m *CatModel) BatchDelete(ctx context.Context, ids ...string) (int64, error) {
	result := m.DB.WithContext(ctx).Where("id = ?", ids).Delete(&cat.Cat{})
	return result.RowsAffected, result.Error
}

func (m *CatModel) BatchUpdate(ctx context.Context, cat []*cat.Cat) (int64, error) {
	result := m.DB.WithContext(ctx).Updates(cat)
	return result.RowsAffected, result.Error
}
