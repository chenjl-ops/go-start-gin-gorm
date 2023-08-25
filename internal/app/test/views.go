package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-starter-gin-gorm/internal/pkg/mysql_gorm"
	//"go-starter-gin-gorm/internal/pkg/apollo"
	"go-starter-gin-gorm/internal/pkg/nacos"
	//"go-starter-gin/internal/pkg/rdssentinels"
	//"time"
)

// @Tags Test API
// @Summary List apollo some config
// @Description get apollo config
// @Accept  json
// @Produce  json
// @Success 200 {array} Response
// @Header 200 {string} Response
// @Failure 400,404 {object} string "Bad Request"
// @Router /v1/test1 [get]
func test(c *gin.Context) {
	c.JSON(200, Response{
		Code: 200,
		Data: map[string]interface{}{
			"MysqlHost":     nacos.Config.MysqlHost,
			"MysqlUserName": nacos.Config.MysqlUserName,
			"MysqlPort":     nacos.Config.MysqlPort,
		},
		Msg: "success",
	})
}

func testInsertSql(c *gin.Context) {
	ticket := RobotTicket{Id: 8888, MessageId: "8888", Status: "ing", ProcessingPerson: "8888", Desc: "test insert"}
	err := mysql_gorm.Insert(&ticket)
	if err != nil {
		log.Error("Insert data: ", err)
	}

	data := make([]RobotTicket, 0)
	err = mysql_gorm.ShowSome(&data, "=", "message_id", "8888")
	if err != nil {
		log.Error("Get data: ", err)
	}
	c.JSON(200, Response{
		Code: 200,
		Data: data,
		Msg:  "success",
	})
}

func testShowAllSql(c *gin.Context) {
	mysql_gorm.Engine.AutoMigrate(&RobotTicket{})
	//robotTicket := []*RobotTicket{
	//	&RobotTicket{Id: 123456, MessageId: "123456", Status: "close", ProcessingPerson: "3721", Desc: "test1"},
	//	&RobotTicket{Id: 23456, MessageId: "23456", Status: "open", ProcessingPerson: "1234", Desc: "test2"},
	//}
	//result := mysql_gorm.Engine.Create(robotTicket)
	//fmt.Sprintln(result)

	ticketData := make([]RobotTicket, 0)
	err := mysql_gorm.ShowAll(&ticketData)

	//result := mysql_gorm.Engine.Find(&ticketData)
	//if result.Error != nil {
	//	log.Error("Get data: ", result.Error)
	//	fmt.Sprintln("get data error: ", result.Error)
	//}

	//if err := mysql_gorm.Engine.Find(&ticketData); err != nil {
	if err != nil {
		log.Error("Get data: ", err)
		fmt.Sprintln("get data error: ", err)
	}
	log.Info("Get data: ", ticketData)
	c.JSON(200, gin.H{
		"ticket_data": ticketData,
	})
}

func testShowSql(c *gin.Context) {
	ticketData := make([]RobotTicket, 0)
	err := mysql_gorm.ShowSome(&ticketData, "LIKE", "message_id", "2345")

	if err != nil {
		log.Error("Get data: ", err)
		fmt.Sprintln("get data error: ", err)
	}
	log.Info("Get data: ", ticketData)
	c.JSON(200, gin.H{
		"ticket_data": ticketData,
	})
}

//func testRedis(c *gin.Context) {
//	//rds := rdssentinels.NewRedis(nil)
//	//rds.SetKey("testKey", "this is test demo", 3600 * time.Second)
//	//result := rds.GetKey("testKey")
//
//	rdssentinels.RedisConfig.SetKey("testKey", "this is test demo", 3600*time.Second)
//	result := rdssentinels.RedisConfig.GetKey("testKey")
//
//	c.JSON(200, gin.H{
//		"redis_testKey": result.Val(),
//	})
//}
