package initialize

import (
	"time"

	"github.com/beego/beego/v2/core/logs"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	initMySql()
	initLogs()
}

// initMySql 初始化mysql
func initMySql() {
	dbHost, _ := beego.AppConfig.String("DB_HOST")
	dbPort, _ := beego.AppConfig.String("DB_PORT")
	dbUser, _ := beego.AppConfig.String("DB_USER")
	dbPassword, _ := beego.AppConfig.String("DB_PASSWORD")
	dbName, _ := beego.AppConfig.String("DB_NAME")

	if dbPort == "" {
		dbPort = "3306"
	}

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&loc=Asia%2FShanghai"

	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		panic("mysql数据库驱动异常")
	}
	// 参数1        数据库的别名，用来在ORM中切换数据库使用
	// 参数2        driverName
	// 参数3        对应的链接字符串
	// 参数4(可选)  设置最大空闲连接  orm.MaxIdleConnections(maxIdle)
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2) username:password@tcp(127.0.0.1:3306)/db_name   orm.MaxOpenConnections(maxConn)
	//maxIdle := 30
	//maxConn := 30
	// set default database
	//fmt.Println(dsn)
	err = orm.RegisterDataBase("default", "mysql", dsn)
	if err != nil {
		panic("mysql数据库连接异常")
	}

	// 自动建表
	// orm.RunSyncdb("default", false, true)
	// 设置为 UTC 时间(default：本地时区)
	orm.DefaultTimeLoc = time.UTC
}

// initLogs 初始化日志
func initLogs() {
	// 日志：会保存手动输出的日志和系统异常日志
	// 如： logs.Error和panic
	logs.Async()

	//设置生产环境日志输出级别，此处设置输出 信息级别，不输出debug日志
	level, _ := beego.AppConfig.Int("logLevel")
	logs.SetLevel(level)

	// level 日志保存的时候的级别，默认是 Trace 级别，level值越高，记录的日志范围越广
	// filename 保存的文件名
	// maxlines 每个文件保存的最大行数，默认值 1000000
	// maxsize 每个文件保存的最大尺寸，默认值是 1 << 28, //256 MB
	// daily 是否按照每天 logrotate，默认是 true
	// maxdays 文件最多保存多少天，默认保存 7 天
	// rotate 是否开启 logrotate，默认是 true
	// perm 日志文件权限
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"./logs/logger.log",
	"level":6,"maxlines":0,"maxsize":0,"daily":true,"maxdays":30,
	"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info"]}`)
}
