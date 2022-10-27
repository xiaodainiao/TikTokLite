package config

import (
	"TikTokLite/util"
	"github.com/spf13/viper"
)

type MysqlConfig struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

type MinioConfig struct {
	Host            string
	Port            string
	AccessKeyID     string
	SecretAccessKey string
	Videobuckets    string
	Picbuckets      string
}

type RedisConfig struct {
	Host    string
	Port    string
	Network string
	Auth    string
}

type PathConfig struct {
	Videofile string
	Logfile   string
	Picfile   string
}

type Configs struct {
	Mysql MysqlConfig
	Minio MinioConfig
	Redis RedisConfig
	Path  PathConfig
	Level string
}

var Config Configs

func LoadConfig() {
	viper.SetConfigFile("./config.yaml")
	viper.ReadInConfig()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	path := PathConfig{
		Videofile: viper.GetString("videofile"),
		Logfile:   viper.GetString("logfile"),
		Picfile:   viper.GetString("picfile"),
	}

	mysql := MysqlConfig{
		Host:     viper.GetString("mysql.host"),
		Port:     viper.GetString("mysql.port"),
		Database: viper.GetString("mysql.database"),
		Username: viper.GetString("mysql.username"),
		Password: viper.GetString("mysql.password"),
	}

	minio := MinioConfig{
		Host:            viper.GetString("minio.host"),
		Port:            viper.GetString("minio.port"),
		AccessKeyID:     viper.GetString("minio.accessKeyID"),
		SecretAccessKey: viper.GetString("minio.secretAccessKey"),
		Videobuckets:    viper.GetString("minio.videobuckets"),
		Picbuckets:      viper.GetString("minio.picbuckets"),
	}

	redis := RedisConfig{
		Host:    viper.GetString("redis.address"),
		Port:    viper.GetString("redis.port"),
		Network: viper.GetString("redis.network"),
		Auth:    viper.GetString("redis.auth"),
	}

	Config = Configs{
		Mysql: mysql,
		Minio: minio,
		Redis: redis,
		Path:  path,
		Level: viper.GetString("level"),
	}

	err = util.Mkdir(path.Videofile)
	if err != nil {
		panic("mkdir videofile error")
	}
	err = util.Mkdir(path.Picfile)
	if err != nil {
		panic("mkdir picfile error")
	}
}

func GetConfig() Configs {
	return Config
}
