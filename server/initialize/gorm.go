package initialize

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/ve-weiyi/ve-blog-golang/server/config/properties"
	"github.com/ve-weiyi/ve-blog-golang/server/global"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn
func Gorm() {
	var cfg properties.Mysql

	cfg = global.CONFIG.Mysql
	global.DB = Open(cfg)

	global.LOG.Infof("Mysql 数据库连接成功！%s", cfg.Dsn())
}

func Open(cfg properties.Mysql) *gorm.DB {
	dsn := cfg.Dsn()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//PrepareStmt:            true, // 缓存预编译语句
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 表前缀
			TablePrefix: cfg.Prefix,
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
		// gorm日志模式
		//Logger: logger.Default.LogMode(logger.Info),
		Logger: logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false, // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,  // 彩色打印
			ParameterizedQueries:      false, // 使用参数化查询 (true时，会将参数值替换为?)
		}),
	})

	if err != nil {
		log.Fatalf("GORM 数据库连接失败: %v", err)
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("SQL 数据库连接失败: %v", err)
		return nil
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}

// RegisterTables 注册数据库表专用
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
	// 系统模块表

	// 自动化模块表
	// Code generated by github.com/ve-weiyi/ve-blog-golang/server Begin; DO NOT EDIT.

	// Code generated by github.com/ve-weiyi/ve-blog-golang/server End; DO NOT EDIT.
	)
	if err != nil {
		global.LOG.Errorf("register table failed %v", err)
		os.Exit(0)
	}
	global.LOG.Println("register table success")
}
