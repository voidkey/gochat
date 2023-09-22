package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	Rdb *redis.Client
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app", viper.Get("app"))
	fmt.Println("config mysql", viper.Get("mysql"))
}

func InitMySQL() {
	//自定义日志模版，打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢sql阈值
			LogLevel:      logger.Info, //级别
			Colorful:      true,        //彩色
		},
	)

	var err error
	DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dsn")), &gorm.Config{Logger: newLogger})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
}

func InitRedis(ctx context.Context) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"), // 密码
		DB:           viper.GetInt("redis.DB"),          // 数据库
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConns"), // 连接池大小
	})
	pong, err := Rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Init redis error: ", err)
	} else {
		fmt.Println("Init redis successed: ", pong)
	}
}

const (
	PublishKey = "websocket"
)

// 发布消息到Redis
func Publish(ctx context.Context, channel string, msg string) error {
	var err error
	fmt.Println("Publish...", msg)
	err = Rdb.Publish(ctx, channel, msg).Err()
	return err
}

// 订阅Redis消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := Rdb.Subscribe(ctx, channel)
	fmt.Println("Subscribe.", ctx)
	msg, err := sub.ReceiveMessage(ctx)
	fmt.Println("Subscribe...", msg.Payload)
	return msg.Payload, err
}
