package mysql_gorm

type DB struct {
	Host        string `json:"host"`                        // db连接地址
	Port        int    `json:"port"`                        // 端口
	UserName    string `json:"username"`                    // 用户名
	Password    string `json:"password"`                    // 密码
	DBName      string `json:"dbname"`                      // 库名
	Charset     string `json:"charset" default:"utf8mb4"`   // 字符集
	MaxIdleConn int    `json:"max_idle_conn" default:"10"`  // 最大空闲连接
	MaxOpenConn int    `json:"max_open_conn" default:"100"` // 最大连接
}
