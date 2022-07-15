package protocol

import (
	"codeup.aliyun.com/625e2dd5594c6cca64844075/go-restful-api/app"
	"codeup.aliyun.com/625e2dd5594c6cca64844075/go-restful-api/conf"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/middleware/cors"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"net/http"
	"time"
)

// NewHTTPService 构建函数
func NewHTTPService() *HTTPService {
	r := gin.Default()

	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1M
		Addr:              conf.GetConfig().App.HttpAddr(),
		Handler:           cors.AllowAll().Handler(r),
	}
	return &HTTPService{
		server: server,
		l:      zap.L().Named("HTTP Service"),
		r:      r,
	}
}

// HTTPService http服务
type HTTPService struct {
	server *http.Server
	l      logger.Logger
	r      gin.IRouter
}

// Start 启动服务
func (s *HTTPService) Start() error {
	// 装置子服务路由
	app.InitGin(s.r)

	apps := app.LoadedGinApps()
	s.l.Info("loaded gin apps:%v", apps)

	// 启动 HTTP服务
	s.l.Infof("HTTP服务启动成功, 监听地址: %s", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			s.l.Info("service is stopped")
		}
		return fmt.Errorf("start service error, %s", err.Error())
	}
	return nil
}

// Stop 停止server
func (s *HTTPService) Stop() error {
	s.l.Info("start graceful shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// 优雅关闭HTTP服务
	if err := s.server.Shutdown(ctx); err != nil {
		s.l.Errorf("graceful shutdown timeout, force exit")
	}
	return nil
}
