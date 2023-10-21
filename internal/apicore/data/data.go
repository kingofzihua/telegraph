package data

import (
	"context"

	"github.com/kingofzihua/telegraph/internal/common/transaction"

	"github.com/go-ostrich/pkg/log"

	"gorm.io/gorm"
)

// Data .
type Data struct {
	db  *gorm.DB
	log *log.Logger
}

// InTx Transaction Database Transaction manager
func (d *Data) InTx(ctx context.Context, f func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return f(ctx)
	})
}

type contextTxKey struct{}

func NewData(db *gorm.DB, log *log.Logger) *Data {
	return &Data{
		db:  db,
		log: log.With("module", "data"),
	}
}

// DB return *gorm.DB warp tx
func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db.WithContext(ctx)
}

// NewTransaction return mysql local Transaction manager
func NewTransaction(data *Data) transaction.Transaction {
	return data
}
