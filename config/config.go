package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var GlobalServerConfig ServerConfig

type MysqlConfig struct {
	Host         string `mapstructure:"host" json:"host"`
	Port         int    `mapstructure:"port" json:"port"`
	Name         string `mapstructure:"db" json:"db"`
	User         string `mapstructure:"user" json:"user"`
	Password     string `mapstructure:"password" json:"password"`
	DbName       string `mapstructure:"db_name" json:"db_name"`
	MaxOpenConns int    `mapstructure:"max_open_conns" json:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"json:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Password string `mapstructure:"password" json:"password"`
}

type RabbitMqConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Exchange string `mapstructure:"exchange" json:"exchange"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

type EtcdConfig struct {
	Endpoints []string `mapstructure:"endpoints" json:"endpoints"`
	Key       string   `mapstructure:"key" json:"key"`
}

type OtelConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
}

type ServerConfig struct {
	Name         string         `mapstructure:"name" json:"name"`
	Host         string         `mapstructure:"host" json:"host"`
	Port         string         `mapstructure:"port" json:"port"`
	WsAddr       string         `mapstructure:"wsAddr" json:"wsAddr"`
	MysqlInfo    MysqlConfig    `mapstructure:"mysql" json:"mysql"`
	RedisInfo    RedisConfig    `mapstructure:"redis" json:"redis"`
	RabbitMqInfo RabbitMqConfig `mapstructure:"rabbitmq" json:"rabbitmq"`
	OtelInfo     OtelConfig     `mapstructure:"otel" json:"otel"`
	TripSrvInfo  TripSrvConfig  `mapstructure:"trip_srv" json:"trip_srv"`
	EtcdInfo     EtcdConfig     `mapstructure:"etcd" json:"etcd"`
}

type TripSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

func Init(fileName string) (err error) {

	viper.SetConfigFile(fileName)

	err = viper.ReadInConfig()
	if err != nil {
		// 读取配置信息失败
		log.Printf("viper.ReadInconfig() failed,err:%v", err)
		return
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	if err := viper.Unmarshal(&GlobalServerConfig); err != nil {
		fmt.Printf("viper.Unmarshal failed,err:%v\n", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("配置文件修改了...")
		if viper.Unmarshal(&GlobalServerConfig); err != nil {
			log.Printf("viper.Unmarshal failed,err:%v \n", err)
		}
	})
	return
}
