package adapters

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/go-ostrich/pkg/log"
	"github.com/go-ostrich/pkg/util"
	"gorm.io/gorm"

	"github.com/kingofzihua/telegraph/internal/apicore/data"
	"github.com/kingofzihua/telegraph/internal/apicore/data/model"
	"github.com/kingofzihua/telegraph/internal/apicore/domain/page"
)

var _ page.Repository = (*MySQLPageRepository)(nil)

type MySQLPageRepository struct {
	data *data.Data
	log  *log.Logger
}

func NewMySQLPageRepository(data *data.Data, logger *log.Logger) *MySQLPageRepository {
	return &MySQLPageRepository{
		data: data,
		log:  log.With(logger, "module", "repo"),
	}
}

func (r *MySQLPageRepository) CreatePage(ctx context.Context, page *page.Page) error {
	m, err := marshalPage(page)
	if err != nil {
		return err
	}
	return r.data.DB(ctx).Create(&m).Error
}

func (r *MySQLPageRepository) UpdatePage(ctx context.Context, pk int32, updateFn func(ctx context.Context, page *page.Page) (*page.Page, error)) error {
	// 查询数据
	p, err := r.QueryByID(ctx, pk)
	if err != nil {
		return err
	}
	if p == nil {
		return fmt.Errorf("MySQLPageRepository.UpdatePage error: id(%d) is nil", pk)
	}
	// 更新数据
	updatedPage, err := updateFn(ctx, p)
	if err != nil {
		return err
	}
	return r.data.DB(ctx).Save(&updatedPage).Error
}

func (r *MySQLPageRepository) QueryByID(ctx context.Context, id int32) (*page.Page, error) {
	var p model.Page

	err := r.data.DB(ctx).Where("id = ?", id).First(&p).Error
	if err != nil {
		// empty
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.Wrapf(err, "MySQLPageRepository.QueryByID(%d)", id)
	}

	return unmarshalPage(p)
}

func (r *MySQLPageRepository) QueryByIDs(ctx context.Context, ids []int32) ([]*page.Page, error) {
	var pages []model.Page

	err := r.data.DB(ctx).Where("id IN ?", ids).Find(&pages).Error
	if err != nil {
		return nil, err
	}

	return unmarshalPages(pages)
}

func (r *MySQLPageRepository) DeleteByID(ctx context.Context, pk int32) error {
	err := r.data.DB(ctx).Delete(&model.Page{}, pk).Error
	// empty
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return err
}

func marshalPage(page *page.Page) (model.Page, error) {
	p, err := util.Copy[model.Page](page)
	if err != nil {
		return model.Page{}, fmt.Errorf("adapters.marshalPage error:%+v", err)
	}
	return p, nil
}

func unmarshalPage(model model.Page) (*page.Page, error) {
	p, err := util.Copy[page.Page](model)
	if err != nil {
		return nil, fmt.Errorf("adapters.unmarshalPage error:%+v", err)
	}
	return &p, nil
}

func unmarshalPages(models []model.Page) ([]*page.Page, error) {
	pages := make([]*page.Page, 0, len(models))
	for _, m := range models {
		p, err := unmarshalPage(m)
		if err != nil {
			return nil, err
		}
		pages = append(pages, p)
	}
	return pages, nil
}
