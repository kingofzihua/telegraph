package biz

import (
	"context"

	"github.com/go-ostrich/pkg/log"
)

type PageRepo interface {
	Create(context.Context, *Page) error
	Update(context.Context, *Page) error
	QueryByID(ctx context.Context, id uint32) (*Page, error)
	QueryByIDs(ctx context.Context, ids []uint32) ([]*Page, error)
	DeleteByID(ctx context.Context, id uint32) error
}

type PageUsecase struct {
	pRepo PageRepo
	tm    Transaction
	log   *log.Logger
}

func NewPageUsecase(pRepo PageRepo, tm Transaction, logger *log.Logger) *PageUsecase {
	return &PageUsecase{
		pRepo: pRepo,
		tm:    tm,
		log:   logger.With("module", "biz"),
	}
}
