package main

import (
	_ "devops-platform/internal/common/init"
	_ "devops-platform/internal/deploy-system/init"
	"devops-platform/pkg/beans"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

func init() {
	beans.RegisterStopWaiter(func() {
		logrus.Info("服务启动成功...")
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		s := <-quit
		logrus.WithField("signal", s.String()).Info("接收到停止信号")
	})
}
