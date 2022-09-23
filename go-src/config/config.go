package config

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	API_VERSION       = "API_VERSION"
	EXCLUDE_REGIONS   = "EXCLUDE_REGIONS"
	SUMMARY_END_POINT = "SUMMARY_END_POINT"
	SUMMARY_HOST      = "SUMMARY_HOST"
	SUMMARY_KEY       = "SUMMARY_KEY"
	HISTORY_END_POINT = "HISTORY_END_POINT"
	HISTORY_HOST      = "HISTORY_HOST"
	HISTORY_KEY       = "HISTORY_KEY"
	AWS_REGION        = "AWS_REGION"
	DB_TABLE_NAME     = "DB_TABLE_NAME"
)

// The init function initialises the logger only once
var logger *log.Logger

func init() {
	logger = &log.Logger{
		Out:       os.Stdout,
		Level:     log.InfoLevel,
		Formatter: &log.JSONFormatter{},
	}
}

func Logger() *log.Logger {
	return logger
}

// The environment variables dictate end-points and configuration values -- all are mandatory
// If one of the environment variables cannot be found a PANIC will be initiated
func variableMustExist(key string) {

	if os.Getenv(key) == "" {
		str := fmt.Sprintf("Environment variable '%s' cannot be empty. Exiting..", key)
		panic(str)
	}
}

func GetApiVersion() string {
	variableMustExist(API_VERSION)
	return os.Getenv(API_VERSION)
}

func GetExcludeRegion() string {
	variableMustExist(EXCLUDE_REGIONS)
	return os.Getenv(EXCLUDE_REGIONS)
}

func GetSummaryEndPoint() string {
	variableMustExist(SUMMARY_END_POINT)
	return os.Getenv(SUMMARY_END_POINT)
}

func GetSummaryHost() string {
	variableMustExist(SUMMARY_HOST)
	return os.Getenv(SUMMARY_HOST)
}

func GetSummaryKey() string {
	variableMustExist(SUMMARY_KEY)
	return os.Getenv(SUMMARY_KEY)
}

func GetHistoryEndPoint() string {
	variableMustExist(HISTORY_END_POINT)
	return os.Getenv(HISTORY_END_POINT)
}

func GetHistoryHost() string {
	variableMustExist(HISTORY_HOST)
	return os.Getenv(HISTORY_HOST)
}

func GetHistoryKey() string {
	variableMustExist(HISTORY_KEY)
	return os.Getenv(HISTORY_KEY)
}

func GetAwsRegion() string {
	variableMustExist(AWS_REGION)
	return os.Getenv(AWS_REGION)
}

func GetDbTableName() string {
	variableMustExist(DB_TABLE_NAME)
	return os.Getenv(DB_TABLE_NAME)
}
