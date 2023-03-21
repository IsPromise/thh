package dbconnect

import (
	"fmt"
	"thh/bundles/bootstrap"

	"github.com/glebarez/sqlite"

	//"gorm.io/driver/sqlite"
	"log"
	"os"
	"path/filepath"
	"thh/arms"
	"thh/arms/logger"
	"thh/bundles/config"
	"time"

	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	// GORM 的 MYSQL 数据库驱动导入
	"gorm.io/driver/mysql"
)

func init() {
	bootstrap.AddDInit(connectDB)
}

// DB gorm.DB 对象
var dbIns *gorm.DB

func Std() *gorm.DB {
	return dbIns
}

// NewMysql dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
// NewMysql
func NewMysql(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

// ConnectDB 初始化模型
func connectDB() {
	fmt.Println("init connectDB")
	var err error
	switch config.GetString("DB_CONNECTION", "sqlite") {
	case "sqlite":
		logger.Info("use sqlite")
		dbIns, err = connectSqlLiteDB(logger.NewGormLogger())
		break
	case "mysql":
		logger.Info("use mysql")
		dbIns, err = connectMysqlDB(logger.NewGormLogger())
		break
	default:
		logger.Info("use sqlite because unselect db")
		dbIns, err = connectSqlLiteDB(logger.NewGormLogger())
		break
	}

	if err != nil {
		log.Println(err)
		panic(err)
	}

	// 获取底层的 sqlDB
	sqlDB, _ := dbIns.DB()
	var (
		maxOpenCoons = config.GetInt("DB_MAX_IDLE_CONNECTIONS", 20)
		maxIdleCoons = config.GetInt("DB_MAX_OPEN_CONNECTIONS", 20)
		maxLifetime  = time.Duration(config.GetInt("DB_MAX_LIFE_SECONDS", 300))
	)
	// 设置最大连接数
	sqlDB.SetMaxOpenConns(maxOpenCoons)
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(maxIdleCoons)
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(maxLifetime * time.Second)
}

func connectMysqlDB(_logger gormlogger.Interface) (*gorm.DB, error) {
	// 初始化 MySQL 连接信息
	var (
		dbUrl = config.GetString("DATABASE_URL", "db_user:db_pass@tcp(db_host:3306)/db_name?charset=utf8mb4&parseTime=True&loc=Local")
		debug = config.GetBool("APP_DEBUG")
	)

	gormConfig := mysql.New(mysql.Config{
		DSN: dbUrl,
	})

	// 准备数据库连接池
	db, err := gorm.Open(gormConfig, &gorm.Config{
		Logger: _logger,
	})
	if debug && err == nil {
		fmt.Println("开启debug")
		db = db.Debug()
	}
	return db, err
}

func connectSqlLiteDB(_logger gormlogger.Interface) (*gorm.DB, error) {
	var (
		dbPath = config.Get("DB_PATH", ":memory:")
		debug  = config.GetBool("APP_DEBUG")
	)

	dbDir := filepath.Dir(dbPath)
	if dbPath == ":memory:" {

	} else if !arms.IsExist(dbPath) {
		if err := os.MkdirAll(dbDir, os.ModePerm); err != nil {

		}
		arms.PutContent(dbPath, "")
	}
	// ":memory:"
	db, err := gorm.Open(sqlite.Open(dbPath+"?_pragma=busy_timeout(5000)"), &gorm.Config{Logger: _logger})
	if debug && err == nil {
		fmt.Println("开启debug")
		db = db.Debug()
	}
	return db, err
}
