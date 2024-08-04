package db

import (
	"fmt"
	"server/common"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"github.com/redis/go-redis/v9"
)

// 鉴权服务器需要使用redis
var (
	mmodb 			*gorm.DB
	mmodbOnce  		sync.Once

	mmoRedis		*redis.Client
	mmoRedisOnce 	sync.Once
)

func createMMODBConnection(dbname, host, user, password string, port int) *gorm.DB {
	// data source name user:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=local 

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=local", user, password, host, port, dbname)
	
	var err error 
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{
		PrepareStmt: true,	// SQL预编译，提高查询效率
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,	// 单数表名
		},
	})

	if err != nil {
		panic("connect to mysql use dsn failed") // TODO
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("get sql db failed")
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(20)


	return db
}

func GetMMODBConnection() {
	mmodbOnce.Do(func() {
		if mmodb == nil {
			dbName := "mmo"
			viper := common.CreateConfig("mysql")
			host := viper.GetString(dbName + ".host")
			port := viper.GetInt(dbName + ".port")
			user := viper.GetString(dbName + ".user")
			password := viper.GetString(dbName + ".password")
			mmodb = createMMODBConnection(dbName, host, user, password, port)
		}
	})
}


func createRedisClient(address, password string, db int) *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr: address,
		Password: password,
		DB: 	db,
	})

	return cli
}

func GetRedisClient() *redis.Client {
	mmoRedisOnce.Do(func() {
		if mmoRedis == nil {
			viper := common.CreateConfig("redis")
			addr := viper.GetString("addr")
			pass := viper.GetString("password") // 没设置该配置项时，viper会默认赋零值，不会报错
			db := viper.GetInt("db")
			mmoRedis = createRedisClient(addr, pass, db)
		}
	})

	return mmoRedis
}