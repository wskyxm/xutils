package xdao

import (
	"database/sql"
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"time"
	"xutils/src/xlog"
)

var conn *gorm.DB

func Initialize(dsn DsnParams, models []any) (err error) {
	// 初始化GORM配置
	config := &gorm.Config{
		Logger: logger.New(
			xlog.Logger(xlog.WarningLog, 0), // 日志输出的目标
			logger.Config{
				SlowThreshold: time.Millisecond * 200, // 慢SQL阈值
				LogLevel: logger.Warn, // 日志级别
				IgnoreRecordNotFoundError: true, // 忽略ErrRecordNotFound
				Colorful: false, // 禁用彩色打印
			},
		),
	}

	// 连接数据库
	switch dsn.Type {
	case Sqlite: conn, err = gorm.Open(sqlite.Open(dsn.String()), config)
	case Pgsql: conn, err = gorm.Open(postgres.Open(dsn.String()), config)
	case Mysql: conn, err = gorm.Open(mysql.Open(dsn.String()), config)
	default: err = errors.New("unsupported database type")
	}

	// 返回结果
	return conn.AutoMigrate(models...)
}

func Exist(model, query interface{}, args ...interface{}) bool {
	return conn.Model(model).Where(query, args...).First(model).RowsAffected == 1
}

func Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return conn.Transaction(fc, opts...)
}

func Create(value interface{}) *gorm.DB {
	return conn.Create(value)
}

func Clauses(conds ...clause.Expression) *gorm.DB {
	return conn.Clauses(conds...)
}

func Table(name string, args ...interface{}) *gorm.DB {
	return conn.Table(name, args...)
}

func Where(query interface{}, args ...interface{}) *gorm.DB {
	return conn.Where(query, args...)
}

func Model(value interface{}) *gorm.DB {

	return conn.Model(value)
}

func Select(query interface{}, args ...interface{}) *gorm.DB {
	return conn.Select(query, args...)
}

func Raw(sql string, values ...interface{}) *gorm.DB {
	return conn.Raw(sql, values...)
}

func Count(m interface{}, query interface{}, args ...interface{}) (count int64) {
	conn.Model(m).Where(query, args...).Count(&count)
	return count
}
