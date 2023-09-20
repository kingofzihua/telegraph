package service

import (
	"context"

	"github.com/go-ostrich/pkg/log"

	"github.com/kingofzihua/telegraph/internal/apicore/biz"
	"github.com/kingofzihua/telegraph/internal/apicore/data"
	"github.com/kingofzihua/telegraph/internal/apicore/data/page"

	v1 "github.com/kingofzihua/telegraph/proto/server/v1"
)

var _ v1.PageServiceServer = (*PageService)(nil)

type PageService struct {
	v1.UnimplementedPageServiceServer
	pUc *biz.PageUsecase
	log *log.Logger
}

func NewPageService(logger *log.Logger) *PageService {
	d := data.NewData(data.GetDefaultDBClient(), logger)
	repo := page.NewPageStore(d, logger)
	uc := biz.NewPageUsecase(repo, d, logger)
	return &PageService{
		UnimplementedPageServiceServer: v1.UnimplementedPageServiceServer{},
		pUc:                            uc,
		log:                            logger.With("module", "srv"),
	}
}

func (p *PageService) CreatePage(ctx context.Context, req *v1.CreatePageRequest) (*v1.CreatePageResponse, error) {
	panic("implement me")
}

func (p *PageService) UpdatePage(ctx context.Context, req *v1.UpdatePageRequest) (*v1.UpdatePageResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PageService) GetPage(ctx context.Context, req *v1.GetPageRequest) (*v1.GetPageResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PageService) GetPageList(ctx context.Context, req *v1.GetPageListRequest) (*v1.GetPageListResponse, error) {
	//TODO implement me
	panic("implement me")
}
