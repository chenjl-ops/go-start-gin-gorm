package mysql_gorm

import (
	"fmt"
	"go-starter-gin-gorm/internal/pkg/nacos"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Engine *gorm.DB

func NewDB() (*DB, error) {
	fmt.Sprintln("start mysql new config: ======")
	fmt.Sprintln("nacos config: ", nacos.Config)
	result := &DB{
		UserName: nacos.Config.MysqlUserName,
		Password: nacos.Config.MysqlPassword,
		Port:     nacos.Config.MysqlPort,
		DBName:   nacos.Config.MysqlDBName,
		Host:     nacos.Config.MysqlHost,
		Charset:  nacos.Config.MysqlCharset,
	}

	fmt.Sprintln("end get mysql config: ", result)
	//if input != nil {
	//	if input.Port != 0 {
	//		result.Port = input.Port
	//	}
	//	if input.DBName != "" {
	//		result.DBName = input.DBName
	//	}
	//	if input.Host != "" {
	//		result.Host = input.Host
	//	}
	//	if input.UserName != "" {
	//		result.UserName = input.UserName
	//	}
	//	if input.Password != "" {
	//		result.Password = input.Password
	//	}
	//	if input.Charset != "" {
	//		result.Charset = input.Charset
	//	}
	//	if input.MaxIdleConn != 0 {
	//		result.MaxIdleConn = input.MaxIdleConn
	//	}
	//	if input.MaxOpenConn != 0 {
	//		result.MaxOpenConn = input.MaxOpenConn
	//	}
	//}
	return result, nil
}

func (db *DB) InitMysql() error {
	fmt.Sprintln("start init mysql: =========")
	newEngine, err := db.NewConnect()
	if err != nil {
		return err
	}
	Engine = newEngine
	return nil
}

func (db *DB) NewConnect() (*gorm.DB, error) {
	fmt.Sprintln("db config: ", db)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", db.UserName, db.Password, db.Host, db.Port, db.DBName, db.Charset)
	fmt.Sprintln("mysql connect str: ", dsn)
	gormEngine, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	sqlDB, err := gormEngine.DB()
	if err != nil {
		return nil, err
	}
	if db.MaxOpenConn > 0 {
		sqlDB.SetMaxOpenConns(db.MaxOpenConn)
	}
	if db.MaxIdleConn > 0 {
		sqlDB.SetMaxIdleConns(db.MaxIdleConn)
	}
	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	return gormEngine, nil
}
