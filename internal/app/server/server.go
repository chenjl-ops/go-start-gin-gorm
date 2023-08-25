package server

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/go-playground/validator"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type server struct {
	//Config   *apollo.Specification
	App *gin.Engine
	//Validate *validator.Validate
}

func NewServer() (*server, error) {
	//gin.ForceConsoleColor()
	return &server{
		//Config: globalConfig,
		App: gin.New(),
	}, nil
}

func StartServer() error {
	//initApolloConfig()
	initNacosConfig()
	initMysql()
	initTask()
	//initRedis()
	//initSnowFlake()

	server, err1 := NewServer()
	if err1 != nil {
		return err1
	}
	// 初始化日志
	server.initLog()
	// 初始化swagger
	//server.InitSwagger()
	// 初始化路由
	server.InitRouter(false)

	//启动服务
	//err := server.Run()

	//优雅启动
	err := server.GracefulRun()
	if err != nil {
		return err
	}
	return nil
}

// Run 启动服务
func (s *server) Run() error {
	return s.App.Run(":8080")
}

// GracefulRun 优雅启动服务
func (s *server) GracefulRun() error {

	srv := &http.Server{
		Addr:    ":8080",
		Handler: s.App.Handler(),
	}
	log.Info("Start Service Successful")

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
		return err
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
	return nil
}
