package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	SecretKey   string
	DatabaseURL string
	MinioConfig MinioConfig
}

type MinioConfig struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
	BucketName      string
	Location        string
}

var AppConfig *Config

// LoadConfig loads environment variables and initializes the configuration
func LoadConfig() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Parse MinIO SSL setting
	minioSSL, _ := strconv.ParseBool(getEnv("MINIO_USE_SSL", "false"))

	AppConfig = &Config{
		Port:        getEnv("PORT", "8080"),
		SecretKey:   getEnv("SECRET", ""),
		DatabaseURL: getEnv("DB", ""),
		MinioConfig: MinioConfig{
			Endpoint:        getEnv("ENDPOINT", ""),
			AccessKeyID:     getEnv("ACCESSKEYID", ""),
			SecretAccessKey: getEnv("SECRETACCESSKEY", ""),
			UseSSL:          minioSSL,
			BucketName:      getEnv("BUCKETNAME", ""),
			Location:        getEnv("LOCATION", ""),
		},
	}

	// Validate required environment variables
	if AppConfig.SecretKey == "" {
		log.Fatal("SECRET environment variable is required")
	}

	if AppConfig.DatabaseURL == "" {
		log.Fatal("DB environment variable is required")
	}
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// GetConfig returns the application configuration
func GetConfig() *Config {
	if AppConfig == nil {
		LoadConfig()
	}
	return AppConfig
}
