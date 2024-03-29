package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type AppConfig struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOSTNAME string
	DB_PORT     int
	DB_NAME     string
	JWT_SECRET  string
}

func InitConfig() *AppConfig {
	return ReadEnv()
}

var (
	MIDTRANS_SERVERKEY    string
	CLOUDINARY_KEY        string
	CLOUDINARY_SECRET     string
	CLOUDINARY_CLOUD_NAME string
)

func ReadEnv() *AppConfig {
	app := AppConfig{}
	isRead := true
	if val, found := os.LookupEnv("JWTSECRET"); found {
		app.JWT_SECRET = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBUSER"); found {
		app.DB_USERNAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPASS"); found {
		app.DB_PASSWORD = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBHOST"); found {
		app.DB_HOSTNAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPORT"); found {
		conv, _ := strconv.Atoi(val)
		app.DB_PORT = conv
		isRead = false
	}
	if val, found := os.LookupEnv("DBNAME"); found {
		app.DB_NAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("MIDTRANS_SERVERKEY"); found {
		MIDTRANS_SERVERKEY = val
		isRead = false
	}
	if val, found := os.LookupEnv("CLOUDINARY_KEY"); found {
		CLOUDINARY_KEY = val
		isRead = false
	}
	if val, found := os.LookupEnv("CLOUDINARY_SECRET"); found {
		CLOUDINARY_SECRET = val
		isRead = false
	}
	if val, found := os.LookupEnv("CLOUDINARY_CLOUD_NAME"); found {
		CLOUDINARY_CLOUD_NAME = val
		isRead = false
	}
	//akan membaca file local.env
	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("local")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config: ", err.Error())
			return nil
		}
		app.DB_USERNAME = viper.Get("DBUSER").(string)
		app.DB_PASSWORD = viper.Get("DBPASS").(string)
		app.DB_HOSTNAME = viper.Get("DBHOST").(string)
		app.DB_PORT, _ = strconv.Atoi(viper.Get("DBPORT").(string))
		app.DB_NAME = viper.Get("DBNAME").(string)
		app.JWT_SECRET = viper.Get("JWTSECRET").(string)
		MIDTRANS_SERVERKEY = viper.GetString("MIDTRANS_SERVERKEY")
		CLOUDINARY_KEY = viper.Get("CLOUDINARY_KEY").(string)
		CLOUDINARY_SECRET = viper.Get("CLOUDINARY_SECRET").(string)
		CLOUDINARY_CLOUD_NAME = viper.Get("CLOUDINARY_CLOUD_NAME").(string)
	}
	return &app
}
