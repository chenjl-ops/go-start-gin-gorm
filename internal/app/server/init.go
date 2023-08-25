package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-starter-gin-gorm/internal/app/test"
	"go-starter-gin-gorm/internal/pkg/mysql_gorm"
	"go-starter-gin-gorm/internal/pkg/nacos"
	"go-starter-gin-gorm/internal/pkg/task"
	"time"

	//swaggerFiles "github.com/swaggo/files"
	//ginSwagger "github.com/swaggo/gin-swagger"
	"go-starter-gin-gorm/internal/app/middleware/logger"
	//"go-starter-gin-gorm/internal/app/test"
	"go-starter-gin-gorm/internal/pkg/apollo"
)

/*
TODO
1、Event跨实例间通讯
*/

// 初始化 apollo config
func initApolloConfig() {
	var err error
	err = apollo.ReadRemoteConfig()
	if nil != err {
		log.Fatal(err)
	}
}

// 初始化 nacos config
func initNacosConfig() {
	var err error
	err = nacos.ReadRemoteConfig()
	//fmt.Println("nacos error: ", err)
	if nil != err {
		log.Fatal(err)
	}
}

// initMysql 初始化mysql
func initMysql() {
	fmt.Sprintln("init start init mysql: ========")
	dbCfg, _ := mysql_gorm.NewDB()
	fmt.Sprintln("init get mysql cfg: ", dbCfg)
	err := dbCfg.InitMysql()
	if nil != err {
		log.Fatal(err)
	}
}

// initRedis 初始化redis
//func initRedis() {
//	rdssentinels.NewRedis(nil)
//}

// initLog 初始化log配置
func (s *server) initLog() *gin.Engine {
	logs := logger.LogMiddleware()
	s.App.Use(logs)
	return s.App
}

// initTask 初始化定时任务
func initTask() {
	task.InitTask()
}

// initSnowFlake 初始化雪花算法
//
//	func initSnowFlake() {
//		snowflake.InitSnowWorker(1, 1)
//	}

// InitRouter 加载gin 路由配置
func (s *server) InitRouter(crossDomain bool) *gin.Engine {
	test.Url(s.App)

	if true == crossDomain {
		s.App.Use(cors.New(cors.Config{
			AllowOrigins:  []string{"*"},
			AllowMethods:  []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
			AllowHeaders:  []string{"Origin"},
			ExposeHeaders: []string{"Content-Length"},
			//AllowCredentials: true,
			//AllowOriginFunc: func(origin string) bool {
			//	return origin == "https://github.com"
			//},
			MaxAge: 12 * time.Hour,
		}))
		return s.App
	}
	return s.App
}

// InitSwagger init swagger
//func (s *server) InitSwagger() *gin.Engine {
//	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
//	s.App.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
//	return s.App
//}
