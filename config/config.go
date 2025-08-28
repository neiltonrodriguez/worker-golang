package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Database DatabaseConfig
	AWS      AWSConfig
	SMTP     SMTPConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type AWSConfig struct {
	AccessKey string
	SecretKey string
	Region    string
	QueueURL  string
}

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

var GlobalConfig AppConfig

func (cfg *AppConfig) LoadVariables(envPath ...string) error {
	err := godotenv.Load(envPath...)
	if err != nil {
		log.Println(".env file not found. Loading from system environment", err)
	}

	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	cfg.Database = DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     dbPort,
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Database: os.Getenv("DB_NAME"),
	}

	cfg.AWS = AWSConfig{
		AccessKey: os.Getenv("AWS_ACCESS_KEY"),
		SecretKey: os.Getenv("AWS_SECRET_KEY"),
		Region:    os.Getenv("AWS_REGION"),
		QueueURL:  os.Getenv("SQS_QUEUE_URL"),
	}

	smtpPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	cfg.SMTP = SMTPConfig{
		Host:     os.Getenv("SMTP_HOST"),
		Port:     smtpPort,
		Username: os.Getenv("SMTP_USER"),
		Password: os.Getenv("SMTP_PASS"),
		From:     os.Getenv("SMTP_FROM"),
	}

	return nil
}
