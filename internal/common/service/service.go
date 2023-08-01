package service

import (
	"context"
	"devops-platform/internal/common/database"
	"devops-platform/pkg/common"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func (s *Service) Inject(getBean func(string) interface{}) {
	ok := true
	s.db, ok = getBean(database.BeanDB).(*gorm.DB)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", database.BeanDB)
		return
	}
}

//BeginTransaction
func (s *Service) BeginTransaction(ctx context.Context, point string) (resultCtx context.Context, err error) {

	resultCtx = ctx
	tx, ok := ctx.Value(database.BeanDB).(*gorm.DB)

	if ok {
		tx = tx.Begin()
	} else {
		tx = s.db.Begin()
	}
	err = tx.Error
	if err != nil {
		err = common.WarpServiceError(50001, "数据库开启事务异常", err)
		logrus.WithError(err).WithField("point", point).Error("开启事务异常")
		return
	}
	resultCtx = context.WithValue(ctx, database.BeanDB, tx)

	return
}

//FinishTransaction
func (s *Service) FinishTransaction(ctx context.Context, err error, point string) (errs error) {
	db, ok := ctx.Value(database.BeanDB).(*gorm.DB)
	if !ok {
		return
	}
	if err != nil {
		errs = s.rollbackTransaction(db, err, point)
	} else {
		errs = s.commitTransaction(db, err, point)
	}
	return
}

func (s *Service) rollbackTransaction(db *gorm.DB, err error, point string) (errs error) {
	errs = db.Rollback().Error
	if errs != nil {
		errs = common.WarpServiceError(50002, "数据库回滚异常"+errs.Error(), err)
		logrus.WithError(errs).WithField("point", point).Error("结束事务异常")
	} else {
		errs = err
	}
	return
}
func (s *Service) commitTransaction(db *gorm.DB, err error, point string) (errs error) {
	errs = db.Commit().Error
	if errs != nil {
		errs = common.WarpServiceError(50003, "数据库提交异常"+errs.Error(), err)
		logrus.WithError(errs).WithField("point", point).Error("结束事务异常")
	} else {
		errs = err
	}
	return
}
