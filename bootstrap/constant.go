package bootstrap

import "fmt"

type Constants struct {
	Context             Context
	LogLevel            LogLevel
	RedisKey            RedisKey
	S3BucketPath        BucketPath
	Field               ErrorField
	Tag                 ErrorTag
	SMSTemplates        SMSTemplates
	EmailTemplates      EmailTemplates
	JWTKeysPath         JWTKeysPath
	Metrics             Metrics
	AddressOwners       AddressOwners
}

type Context struct {
	Translator                   string
	IsLoadedValidationTranslator string
	ID                           string
}

type LogLevel struct {
	Debug string
	Info  string
	Warn  string
	Error string
	Fatal string
}

type RedisKey struct {
}

type BucketPath struct {
}

type ErrorField struct {
	User                string
	Phone               string
	Email               string
	Password            string
	OTP                 string
	NationalID          string
	RegistrationNumber  string
	IBAN                string
	Address             string
	Name                string
	Province            string
	City                string
	Page                string
	Role                string
	Permission          string
	News                string
	Media               string
	Post                string
	Like                string

	Picture             string
}

type ErrorTag struct {
	AlreadyRegistered      string
	MinimumLength          string
	ContainsLowercase      string
	ContainsUppercase      string
	ContainsNumber         string
	ContainsSpecialChar    string
	Expired                string
	Invalid                string
	NotRegistered          string
	NotVerified            string
	NotActive              string
	InvalidAuthCredentials string
	ExpiredAuthToken       string
	InvalidAuthToken       string
	Unauthorized           string
	AwaitingApproval       string
	Rejected               string
	NotExist               string
	AlreadyExist           string
	ForbiddenStatus        string
	Pending                string
	AlreadyBlocked         string
	AlreadyActive          string
	AlreadyResolved        string
	AlreadyArchived        string
	StatusNotChange        string
	AlreadyCanceled        string
	AlreadyRejected        string
	AlreadyAccepted        string
	AlreadyDraft           string
}

type SMSTemplates struct {
	OTP string
}

type EmailTemplates struct {
	Path            string
	PersianFileName string
	EnglishFileName string
}

type JWTKeysPath struct {
	PublicKey  string
	PrivateKey string
}

type Metrics struct {
	HTTPRequestsTotal   Options
	HTTPRequestDuration Options
}

type Options struct {
	Name string
	Help string
}

type AddressOwners struct {
	User                string
}

type Queues struct {
	DLQ string
}

type Headers struct {
	RetryCount string
	LastError  string
	DeadLetter string
}

func NewConstants() *Constants {
	return &Constants{
		Context: Context{
			Translator:                   "translator",
			IsLoadedValidationTranslator: "isLoadedValidationTranslator",
			ID:                           "ID",
		},
		LogLevel: LogLevel{
			Debug: "debug",
			Info:  "info",
			Warn:  "warn",
			Error: "error",
			Fatal: "fatal",
		},
		Field: ErrorField{
			User:                "user",
			Phone:               "phone",
			Email:               "email",
			Password:            "password",
			OTP:                 "otp",
			NationalID:          "nationalID",
			RegistrationNumber:  "registrationNumber",
			IBAN:                "iban",
			Address:             "address",
			Name:                "name",
			Province:            "province",
			City:                "city",
			Page:                "page",
			Role:                "role",
			Permission:          "permission",
			News:                "news",
			Media:               "media",
			Post:                "post",
			Like:                "like",

			Picture:             "picture",
		},
		Tag: ErrorTag{
			AlreadyRegistered:      "alreadyRegistered",
			MinimumLength:          "minimumLength",
			ContainsLowercase:      "containsLowercase",
			ContainsUppercase:      "containsUppercase",
			ContainsNumber:         "containsNumber",
			ContainsSpecialChar:    "containsSpecialChar",
			Expired:                "Expired",
			Invalid:                "invalid",
			NotRegistered:          "notRegistered",
			NotVerified:            "notVerified",
			NotActive:              "notActive",
			InvalidAuthCredentials: "invalidAuthCredentials",
			ExpiredAuthToken:       "expiredAuthToken",
			InvalidAuthToken:       "invalidAuthToken",
			Unauthorized:           "unauthorized",
			AwaitingApproval:       "awaitingApproval",
			Rejected:               "rejected",
			NotExist:               "notExist",
			AlreadyExist:           "alreadyExist",
			ForbiddenStatus:        "forbiddenStatus",
			Pending:                "pending",
			AlreadyBlocked:         "alreadyBlocked",
			AlreadyActive:          "alreadyActive",
			AlreadyResolved:        "alreadyResolved",
			AlreadyArchived:        "alreadyArchived",
			StatusNotChange:        "statusNotChange",
			AlreadyCanceled:        "alreadyCanceled",
			AlreadyRejected:        "alreadyRejected",
			AlreadyAccepted:        "alreadyAccepted",
			AlreadyDraft:           "alreadyDraft",
		},
		SMSTemplates: SMSTemplates{
			OTP: "sendOTPTemplate",
		},
		JWTKeysPath: JWTKeysPath{
			PublicKey:  "./internal/infrastructure/jwt/publicKey.pem",
			PrivateKey: "./internal/infrastructure/jwt/privateKey.pem",
		},
		EmailTemplates: EmailTemplates{
			Path:            "./internal/infrastructure/communication/email/templates/",
			PersianFileName: "fa.html",
			EnglishFileName: "en.html",
		},
		Metrics: Metrics{
			HTTPRequestsTotal: Options{
				Name: "http_requests_total",
				Help: "Total number of HTTP requests",
			},
			HTTPRequestDuration: Options{
				Name: "http_request_duration_seconds",
				Help: "HTTP request duration in seconds",
			},
		},
		AddressOwners: AddressOwners{
			User:                "users",
		},
	}
}

func (r *RedisKey) GenerateOTPKey(value string) string {
	return fmt.Sprintf("otp:%s", value)
}

func (path *BucketPath) GetUserProfilePath(userID uint, pictureFileName string) string {
	return fmt.Sprintf("user/%d/profile/%s", userID, pictureFileName)
}

func (path *BucketPath) GetNewsMediaPath(newsID uint, MediaFileName string) string {
	return fmt.Sprintf("news/%d/media/%s", newsID, MediaFileName)
}

func (path *BucketPath) GetNewsCoverImagePath(newsID uint, mediaFileName string) string {
	return fmt.Sprintf("news/%d/cover-image/%s", newsID, mediaFileName)
}