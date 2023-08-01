package repository

import (
	"context"
	"devops-platform/internal/common/database"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func (r *Repository) Inject(getBean func(string) interface{}) {

	db, ok := getBean(database.BeanDB).(*gorm.DB)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", database.BeanDB)
		return
	}
	r.db = db
}

func (r *Repository) DB(ctx context.Context) *gorm.DB {
	db, ok := ctx.Value(database.BeanDB).(*gorm.DB)
	if !ok {
		return r.db
	}
	return db
}
