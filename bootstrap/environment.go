package bootstrap

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Env struct {
	Server             Server
	Logger             Logger
	RateLimit          RateLimit
	PrimaryDB          Database
	PrimaryRedis       Redis
	Storage            S3
	OTP                OTP
	SMSGateway         SMSGateway
	Pagination         Pagination
	EmailSenderAccount EmailAccount
	SuperAdmin         AdminCredentials
}

type Server struct {
	Port string
	Mode string
}

type Logger struct {
	LogLevel      string
	LogFile       string
	ConsoleOutput string
}

type RateLimit struct {
	Limit string
	Burst string
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type Redis struct {
	Port      string
	Address   string
	Password  string
	RDBNumber string
}

type S3 struct {
	Buckets   BucketName
	Region    string
	AccessKey string
	SecretKey string
	Endpoint  string
}

type BucketName struct {
	ProfilePic             string
	NewsMedia              string
}

type OTP struct {
	Length       int
	ExpiryMinute int
	MaxAttempts  int
}

type SMSGateway struct {
	APIKey string
}

type Pagination struct {
	DefaultPage     int
	DefaultPageSize int
}

type EmailAccount struct {
	EmailFrom     string
	EmailPassword string
	SMTPHost      string
	SMTPPort      string
}

type AdminCredentials struct {
	FirstName    string
	LastName     string
	Phone        string
	Password     string
	Email        string
	NationalCode string
}

func NewEnvironments() *Env {
	// godotenv.Load("../../.env")
	godotenv.Load(".env")
	return &Env{
		Server: Server{
			Port: os.Getenv("SERVER_PORT"),
			Mode: os.Getenv("SERVER_MODE"),
		},
		Logger: Logger{
			LogLevel:      os.Getenv("LOG_LEVEL"),
			LogFile:       os.Getenv("LOG_FILE"),
			ConsoleOutput: os.Getenv("CONSOLE_OUTPUT"),
		},
		RateLimit: RateLimit{
			Limit: os.Getenv("RATE_LIMIT"),
			Burst: os.Getenv("RATE_lIMIT_BURST"),
		},
		PrimaryDB: Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		PrimaryRedis: Redis{
			Port:      os.Getenv("RDB_PORT"),
			Address:   os.Getenv("RDB_ADDRESS"),
			Password:  os.Getenv("RDB_PASSWORD"),
			RDBNumber: os.Getenv("RDB_NUMBER"),
		},
		Storage: S3{
			Buckets: BucketName{
				ProfilePic:             os.Getenv("PROFILE_PIC_BUCKET_NAME"),
				NewsMedia:              os.Getenv("NEWS_MEDIA_BUCKET_NAME"),
			},
			Region:    os.Getenv("BUCKET_REGION"),
			AccessKey: os.Getenv("BUCKET_ACCESS_key"),
			SecretKey: os.Getenv("BUCKET_SECRET_key"),
			Endpoint:  os.Getenv("BUCKET_ENDPOINT"),
		},
		OTP: OTP{
			Length:       getEnvInt("OTP_LENGTH", 6),
			ExpiryMinute: getEnvInt("OTP_EXPIRY_MINUTES", 2),
			MaxAttempts:  getEnvInt("OTP_MAX_ATTEMPTS", 3),
		},
		SMSGateway: SMSGateway{
			APIKey: os.Getenv("SMS_GATEWAY_API_KEY"),
		},
		Pagination: Pagination{
			DefaultPage:     getEnvInt("DEFAULT_PAGE", 6),
			DefaultPageSize: getEnvInt("DEFAULT_PAGE_SIZE", 2),
		},
		EmailSenderAccount: EmailAccount{
			EmailFrom:     os.Getenv("EMAIL_FROM"),
			EmailPassword: os.Getenv("EMAIL_PASSWORD"),
			SMTPHost:      os.Getenv("SMTP_HOST"),
			SMTPPort:      os.Getenv("SMTP_PORT"),
		},
		SuperAdmin: AdminCredentials{
			FirstName:    os.Getenv("SUPER_ADMIN_FIRST_NAME"),
			LastName:     os.Getenv("SUPER_ADMIN_LAST_NAME"),
			Phone:        os.Getenv("SUPER_ADMIN_PHONE"),
			Password:     os.Getenv("SUPER_ADMIN_PASSWORD"),
			Email:        os.Getenv("SUPER_ADMIN_EMAIL"),
			NationalCode: os.Getenv("SUPER_ADMIN_NATIONAL_CODE"),
		},
	}
}

func getEnvDuration(key string, defaultVal time.Duration) time.Duration {
	if val := os.Getenv(key); val != "" {
		if parsed, err := time.ParseDuration(val); err == nil {
			return parsed
		}
	}
	return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
	if val := os.Getenv(key); val != "" {
		if parsed, err := strconv.Atoi(val); err == nil {
			return parsed
		}
	}
	return defaultVal
}
