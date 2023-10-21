package service

import "github.com/kingofzihua/telegraph/internal/apicore/domain/page"

type PageService struct {
	repo page.Repository
}

func NewPageService(repo page.Repository, logger *log.Logger) *PageService {
	return &PageService{}
}

func (s PageService) CreatePage() {

}
