package service

import (
	"context"
	"devops-platform/internal/common/config"
	"devops-platform/internal/common/web/internal/domain"
	"devops-platform/pkg/beans"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

type httpServerConfig interface {
	GetServerAddress() string
	GetEnv() string
}

type HttpServerLifecycle struct {
	Server *http.Server
}

func (s *HttpServerLifecycle) StartOrder() int {
	return 0
}
func (s *HttpServerLifecycle) StopOrder() int {
	return 0
}
func (s *HttpServerLifecycle) Start() {

	if os.Getenv(domain.ModeKey) == domain.ModeUnitTesting {
		logrus.Info("单元测试模式，不启动HTTP服务器...")
		return
	}

	logrus.WithField("server_address", s.Server.Addr).Info("即将启动HTTP服务器...")
	go s.start()

}

func (s *HttpServerLifecycle) Stop() {

	logrus.Info("即将关闭HTTP服务器")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Server.Shutdown(ctx); err != nil {
		logrus.WithField("server_address", s.Server.Addr).WithError(err).Error("HTTP服务器关闭异常")
	}
}

func (s *HttpServerLifecycle) start() {
	// 服务连接
	if err := s.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.WithField("server_address", s.Server.Addr).WithError(err).Panic("HTTP服务器启动失败")
	}

}

func (s *HttpServerLifecycle) PreInject(getBean func(string) interface{}) {
	cfg, ok := getBean(config.BeanApp).(httpServerConfig)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", config.BeanApp)
		return
	}

	if cfg.GetEnv() == config.EnvProd {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(s.log, gin.Recovery(), cors.New(corsConfig()))
	s.Server = &http.Server{
		Addr:    cfg.GetServerAddress(),
		Handler: router,
	}

	beans.Register(domain.BeanGinEngine, router)

}

func corsConfig() cors.Config {
	return cors.Config{
		//准许跨域请求网站,多个使用,分开,限制使用*
		AllowOrigins: []string{"*"},
		//准许使用的请求方式
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		//准许使用的请求表头
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
		//显示的请求表头
		ExposeHeaders: []string{"Content-Type"},
		//凭证共享,确定共享
		AllowCredentials: true,
		//容许跨域的原点网站,可以直接return true就万事大吉了
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		//超时时间设定
		MaxAge: 12 * time.Hour,
	}
}

func (s *HttpServerLifecycle) log(c *gin.Context) {

	// 开始时间
	startTime := time.Now()

	// 处理请求
	c.Next()

	// 状态码
	statusCode := c.Writer.Status()

	logger := logrus.WithField("latencyTime", time.Now().Sub(startTime)).
		WithField("HttpStatus", statusCode).
		WithField("clientIP", c.ClientIP())

	/*
	 * 打印日志中加err
	 */
	if errInterface, exists := c.Get("err"); exists {
		if err, ok := errInterface.(error); ok {
			logger = logger.WithError(err)
		}
	}

	switch {
	case statusCode >= http.StatusOK && statusCode < http.StatusMultipleChoices:
		logger.Debug(c.Request.Method, " ", c.Request.RequestURI)
	case statusCode >= http.StatusMultipleChoices && statusCode < http.StatusBadRequest:
		logger.Info(c.Request.Method, " ", c.Request.RequestURI)
	case statusCode >= http.StatusBadRequest && statusCode < http.StatusInternalServerError:
		logger.Warn(c.Request.Method, " ", c.Request.RequestURI)
	default:
		logger.Error(c.Request.Method, " ", c.Request.RequestURI)
	}
}
