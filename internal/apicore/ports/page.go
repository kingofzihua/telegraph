package ports

import (
	"context"

	"github.com/go-ostrich/pkg/log"

	v1 "github.com/kingofzihua/telegraph/proto/server/v1"
)

var _ v1.PageServiceServer = (*PageService)(nil)

type PageService struct {
	v1.UnimplementedPageServiceServer
}

func NewPageService(logger *log.Logger) *PageService {
	return &PageService{
		UnimplementedPageServiceServer: v1.UnimplementedPageServiceServer{},
	}
}

func (s *PageService) CreatePage(ctx context.Context, req *v1.CreatePageRequest) (*v1.CreatePageResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PageService) UpdatePage(ctx context.Context, req *v1.UpdatePageRequest) (*v1.UpdatePageResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PageService) GetPage(ctx context.Context, req *v1.GetPageRequest) (*v1.GetPageResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PageService) GetPageList(ctx context.Context, req *v1.GetPageListRequest) (*v1.GetPageListResponse, error) {
	//TODO implement me
	panic("implement me")
}
