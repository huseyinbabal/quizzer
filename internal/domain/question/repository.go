package question

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func New(db *gorm.DB, logger *zap.Logger) (*Repository, error) {
	err := db.AutoMigrate(&Entity{})
	if err != nil {
		return nil, errors.Wrap(err, "while migrating Question entity.")
	}
	return &Repository{db: db, logger: logger}, nil
}

func (r *Repository) List(ctx context.Context) []*Entity {
	dbWithContext := r.db.WithContext(ctx)
	var questions []*Entity
	dbWithContext.Model(&Entity{}).Find(&questions)
	return questions
}
