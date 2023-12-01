package config

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

//config package is to load configurations from .env file

type Config struct {
	PORT               string `mapstructure:"PORT"`
	DBHost             string `mapstructure:"DB_HOST"`
	DBName             string `mapstructure:"DB_NAME"`
	DBUser             string `mapstructure:"DB_USER"`
	DBPort             string `mapstructure:"DB_PORT"`
	DBPassword         string `mapstructure:"DB_PASSWORD"`
	AdminUsername      string `mapstructure:"ADMIN"`
	AdminPassword      string `mapstructure:"ADMINPASS"`
	JwtSecret          string `mapstructure:"JWT_SECRET"`
	TwilioAccountSid   string `mapstructure:"TWILIO_ACCOUNT_SID"`
	TwilioAuthToken    string `mapstructure:"TWILIO_AUTH_TOKEN"`
	TwilioServiceSid   string `mapstructure:"VERIFY_SERVICE_SID"`
	RazorPayKeyId      string `mapstructure:"RAZORPAY_KEY_ID"`
	RazorPayKeySecret  string `mapstructure:"RAZORPAY_KEY_SECRET"`
	AWSRegion          string `mapstructure:"AWS_REGION"`
	AWSAccessKeyID     string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AWSSecretAccessKey string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	S3BucketName       string `mapstructure:"S3_BUCKET_NAME"`
	S3BucketMediaPath  string `mapstructure:"S3_BUCKET_MEDIA_PATH"`
}

type AdminCredentials struct {
	AdminUsername string
	AdminPassword string
}

var (
	envs = []string{
		"DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", "ADMIN",

		"ADMINPASS", "JWT_SECRET", "TWILIO_ACCOUNT_SID", "TWILIO_AUTH_TOKEN", "VERIFY_SERVICE_SID",

		"RAZORPAY_KEY_ID", "RAZORPAY_KEY_SECRET", "AWS_REGION", "AWS_ACCESS_KEY_ID",

		"AWS_SECRET_ACCESS_KEY", "S3_BUCKET_CODENATION", "S3_BUCKET_CHAT_MEDIA_PATH",
	}

	config Config
)

func LoadConfig() (Config, error) {

	viper.AddConfigPath("./")   //set the config path
	viper.SetConfigFile(".env") //set config file
	err := viper.ReadInConfig() //read the config file
	if err != nil {             //handle the error returning from ReadInConfig
		log.Fatal(" Error while reading the config file ", err)
	}

	for _, env := range envs {
		if err = viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}
	return config, nil

}

func GetAdminCredentials() AdminCredentials {
	return AdminCredentials{config.AdminUsername, config.AdminPassword}
}

func GetConfig() Config {
	return config
}
