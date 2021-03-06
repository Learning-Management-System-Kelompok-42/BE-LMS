package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type AppConfig struct {
	App struct {
		Name      string `toml:"name"`
		Port      int    `toml:"port"`
		Env       string `toml:"env"`
		Timeout   int    `toml:"timeout"`
		SecretKey string `toml:"secretkey"`
		Crt       string `toml:"crt"`
		Key       string `toml:"key"`
	} `toml:"app"`
	Database struct {
		Driver   string `toml:"driver"`
		Name     string `toml:"name"` // name of the database
		Address  string `toml:"address"`
		Port     int    `toml:"port"`
		Username string `toml:"username"`
		Password string `toml:"password"`
	} `toml:"database"`
	Storage struct {
		CloudName     string `toml:"cloudname"`
		APIKey        string `toml:"apikey"`
		APISecret     string `toml:"apisecret"`
		UploadStorage string `toml:"uploadfolder"`
		AwsKey        string `toml:"awsKey"`
	} `toml:"storage"`
	AwsS3 struct {
		AwsId  string `toml:"awsId"`
		AwsKey string `toml:"awsKey"`
		Region string `toml:"region"`
		Bucket string `toml:"bucket"`
	} `toml:"awss3"`
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.App.Name = "LMS-API"
	defaultConfig.App.Env = "dev"
	defaultConfig.App.Port = 4001

	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./config/")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	var finalConfig AppConfig
	err := viper.Unmarshal(&finalConfig)
	if err != nil {
		log.Fatalf("Error unmarshalling config file, %s", err)
	}

	return &finalConfig
}
