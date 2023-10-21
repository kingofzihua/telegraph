package page

import "context"

type Repository interface {
	CreatePage(ctx context.Context, page *Page) error
	UpdatePage(ctx context.Context, pk int32, updateFn func(context.Context, *Page) (*Page, error)) error
	QueryByID(ctx context.Context, id int32) (*Page, error)
	QueryByIDs(ctx context.Context, ids []int32) ([]*Page, error)
	DeleteByID(ctx context.Context, id int32) error
}
