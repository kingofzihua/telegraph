package page

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-ostrich/pkg/log"
	"github.com/go-ostrich/pkg/util"
	"gorm.io/gorm"

	"github.com/kingofzihua/telegraph/internal/apicore/biz"
	"github.com/kingofzihua/telegraph/internal/apicore/data"
	"github.com/kingofzihua/telegraph/internal/apicore/data/model"
)

var _ biz.PageRepo = (*pageStore)(nil)

type pageStore struct {
	data *data.Data
	log  *log.Logger
}

func NewPageStore(data *data.Data, logger *log.Logger) biz.PageRepo {
	return &pageStore{
		data: data,
		log:  log.With(logger, "module", "store"),
	}
}

func (s *pageStore) Create(ctx context.Context, page *biz.Page) error {
	p, err := toModelPage(page)
	if err != nil {
		return err
	}

	return s.data.DB(ctx).Create(&p).Error
}

func (s *pageStore) Update(ctx context.Context, page *biz.Page) error {
	category, err := util.Copy[model.Page](page)

	if err != nil {
		return fmt.Errorf("pageStore.Update error:%+v", err)
	}

	return s.data.DB(ctx).Save(&category).Error
}

func (s *pageStore) QueryByID(ctx context.Context, id uint32) (*biz.Page, error) {
	var page model.Page

	err := s.data.DB(ctx).Where("id = ?", id).First(&page).Error
	if err != nil {
		// empty
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return toBizPage(page)
}

func (s *pageStore) QueryByIDs(ctx context.Context, ids []uint32) ([]*biz.Page, error) {
	var pages []model.Page

	err := s.data.DB(ctx).Where("id IN ?", ids).Find(&pages).Error
	if err != nil {
		return nil, err
	}

	return toBizPages(pages)
}

func (s *pageStore) DeleteByID(ctx context.Context, id uint32) error {
	err := s.data.DB(ctx).Delete(&model.Page{}, id).Error
	// empty
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return err
}

func toModelPage(page *biz.Page) (model.Page, error) {
	p, err := util.Copy[model.Page](page)
	if err != nil {
		return model.Page{}, fmt.Errorf("toModelPage error:%+v", err)
	}
	return p, nil
}

func toBizPage(model model.Page) (*biz.Page, error) {
	p, err := util.Copy[biz.Page](model)
	if err != nil {
		return nil, fmt.Errorf("toBizPage error:%+v", err)
	}
	return &p, nil
}

func toBizPages(models []model.Page) ([]*biz.Page, error) {
	pages := make([]*biz.Page, 0, len(models))
	for _, m := range models {
		page, err := toBizPage(m)
		if err != nil {
			return nil, err
		}
		pages = append(pages, page)
	}
	return pages, nil
}
