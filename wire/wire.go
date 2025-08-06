//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/CosmeticsShiraz/Backend/bootstrap"
	"github.com/CosmeticsShiraz/Backend/internal/application/service"
	"github.com/CosmeticsShiraz/Backend/internal/application/usecase"
	"github.com/CosmeticsShiraz/Backend/internal/domain/communication"
	domainLogger "github.com/CosmeticsShiraz/Backend/internal/domain/logger"
	"github.com/CosmeticsShiraz/Backend/internal/domain/message"
	domainMetrics "github.com/CosmeticsShiraz/Backend/internal/domain/metrics"
	domainPostgres "github.com/CosmeticsShiraz/Backend/internal/domain/repository/postgres"
	domainRedis "github.com/CosmeticsShiraz/Backend/internal/domain/repository/redis"
	"github.com/CosmeticsShiraz/Backend/internal/domain/s3"
	"github.com/CosmeticsShiraz/Backend/internal/infrastructure/communication/email"
	"github.com/CosmeticsShiraz/Backend/internal/infrastructure/communication/sms"
	"github.com/CosmeticsShiraz/Backend/internal/infrastructure/database"
	infraJWT "github.com/CosmeticsShiraz/Backend/internal/infrastructure/jwt"
	infraLocalization "github.com/CosmeticsShiraz/Backend/internal/infrastructure/localization"
	infraLogger "github.com/CosmeticsShiraz/Backend/internal/infrastructure/logger"
	infraMetrics "github.com/CosmeticsShiraz/Backend/internal/infrastructure/metrics"
	infraRabbitMQ "github.com/CosmeticsShiraz/Backend/internal/infrastructure/rabbitmq"
	"github.com/CosmeticsShiraz/Backend/internal/infrastructure/rabbitmq/consumer"
	infraPostgres "github.com/CosmeticsShiraz/Backend/internal/infrastructure/repository/postgres"
	infraRedis "github.com/CosmeticsShiraz/Backend/internal/infrastructure/repository/redis"
	"github.com/CosmeticsShiraz/Backend/internal/infrastructure/seed"
	infraStorage "github.com/CosmeticsShiraz/Backend/internal/infrastructure/storage"
	"github.com/CosmeticsShiraz/Backend/internal/infrastructure/websocket"
	"github.com/CosmeticsShiraz/Backend/internal/presentation/controller/v1/address"
	"github.com/CosmeticsShiraz/Backend/internal/presentation/controller/v1/chat"
	"github.com/CosmeticsShiraz/Backend/internal/presentation/controller/v1/corporation"
	"github.com/CosmeticsShiraz/Backend/internal/presentation/controller/v1/guarantee"
	"github.com/CosmeticsShiraz/Backend/internal/presentation/controller/v1/installation"
	"github.com/CosmeticsShiraz/Backend/internal/presentation/controller/v1/maintenance"
	"github.com/CosmeticsShiraz/Backend/internal/presentation/controller/v1/news"
	"github.com/CosmeticsShiraz/Backend/internal/presentation/controller/v1/notification"
	"github.com/CosmeticsShiraz/Backend/internal/presentation/controller/v1/payment"
	"github.com/CosmeticsShiraz/Backend/internal/presentation/controller/v1/report"
	"github.com/CosmeticsShiraz/Backend/internal/presentation/controller/v1/ticket"
	"github.com/CosmeticsShiraz/Backend/internal/presentation/controller/v1/user"
	"github.com/CosmeticsShiraz/Backend/internal/presentation/middleware"
	"github.com/google/wire"
)

var DatabaseProviderSet = wire.NewSet(
	database.NewPostgresDatabase,
	database.NewRedisDatabase,
	wire.Bind(new(database.Database), new(*database.PostgresDatabase)),
	wire.Bind(new(database.Cache), new(*database.RedisDatabase)),
	wire.Struct(new(Database), "*"),
)

var RepositoryProviderSet = wire.NewSet(
	infraPostgres.NewUserRepository,
	infraPostgres.NewInstallationRepository,
	infraPostgres.NewAddressRepository,
	infraRedis.NewUserCacheRepository,
	infraPostgres.NewCorporationRepository,
	infraPostgres.NewChatRepository,
	infraPostgres.NewNotificationRepository,
	infraPostgres.NewMaintenanceRepository,
	infraPostgres.NewTicketRepository,
	infraPostgres.NewReportRepository,
	infraPostgres.NewGuaranteeRepository,
	infraPostgres.NewPaymentRepository,
	infraPostgres.NewNewsRepository,
	wire.Bind(new(domainPostgres.UserRepository), new(*infraPostgres.UserRepository)),
	wire.Bind(new(domainPostgres.InstallationRepository), new(*infraPostgres.InstallationRepository)),
	wire.Bind(new(domainPostgres.AddressRepository), new(*infraPostgres.AddressRepository)),
	wire.Bind(new(domainRedis.UserCacheRepository), new(*infraRedis.UserCacheRepository)),
	wire.Bind(new(domainPostgres.CorporationRepository), new(*infraPostgres.CorporationRepository)),
	wire.Bind(new(domainPostgres.ChatRepository), new(*infraPostgres.ChatRepository)),
	wire.Bind(new(domainPostgres.NotificationRepository), new(*infraPostgres.NotificationRepository)),
	wire.Bind(new(domainPostgres.MaintenanceRepository), new(*infraPostgres.MaintenanceRepository)),
	wire.Bind(new(domainPostgres.TicketRepository), new(*infraPostgres.TicketRepository)),
	wire.Bind(new(domainPostgres.ReportRepository), new(*infraPostgres.ReportRepository)),
	wire.Bind(new(domainPostgres.GuaranteeRepository), new(*infraPostgres.GuaranteeRepository)),
	wire.Bind(new(domainPostgres.PaymentRepository), new(*infraPostgres.PaymentRepository)),
	wire.Bind(new(domainPostgres.NewsRepository), new(*infraPostgres.NewsRepository)),
)

var ServiceProviderSet = wire.NewSet(
	wire.Struct(new(service.UserServiceDeps), "*"),
	wire.Struct(new(service.NotificationServiceDeps), "*"),
	wire.Struct(new(service.InstallationServiceDeps), "*"),
	service.NewUserService,
	service.NewOTPService,
	sms.NewSMSService,
	email.NewEmailService,
	service.NewJWTService,
	service.NewInstallationService,
	service.NewAddressService,
	service.NewCorporationService,
	service.NewChatService,
	service.NewNotificationService,
	service.NewMaintenanceService,
	service.NewTicketService,
	service.NewReportService,
	service.NewGuaranteeService,
	service.NewPaymentService,
	service.NewNewsService,
	wire.Bind(new(usecase.UserService), new(*service.UserService)),
	wire.Bind(new(usecase.OTPService), new(*service.OTPService)),
	wire.Bind(new(communication.SMSService), new(*sms.SMSService)),
	wire.Bind(new(communication.EmailService), new(*email.EmailService)),
	wire.Bind(new(usecase.JWTService), new(*service.JWTService)),
	wire.Bind(new(usecase.InstallationService), new(*service.InstallationService)),
	wire.Bind(new(usecase.AddressService), new(*service.AddressService)),
	wire.Bind(new(usecase.CorporationService), new(*service.CorporationService)),
	wire.Bind(new(usecase.ChatService), new(*service.ChatService)),
	wire.Bind(new(usecase.NotificationService), new(*service.NotificationService)),
	wire.Bind(new(usecase.MaintenanceService), new(*service.MaintenanceService)),
	wire.Bind(new(usecase.TicketService), new(*service.TicketService)),
	wire.Bind(new(usecase.ReportService), new(*service.ReportService)),
	wire.Bind(new(usecase.GuaranteeService), new(*service.GuaranteeService)),
	wire.Bind(new(usecase.PaymentService), new(*service.PaymentService)),
	wire.Bind(new(usecase.NewsService), new(*service.NewsService)),
)

var AdapterProviderSet = wire.NewSet(
	infraLocalization.NewTranslationService,
	infraLogger.NewLogger,
	infraJWT.NewJWTKeyManager,
	infraMetrics.NewPrometheusMetrics,
	infraStorage.NewS3Storage,
	infraRabbitMQ.NewRabbitMQ,
	wire.Bind(new(domainLogger.Logger), new(*infraLogger.Logger)),
	wire.Bind(new(domainMetrics.MetricsClient), new(*infraMetrics.PrometheusMetrics)),
	wire.Bind(new(s3.S3Storage), new(*infraStorage.S3Storage)),
	wire.Bind(new(message.Broker), new(*infraRabbitMQ.RabbitMQ)),
)

var GeneralControllerProviderSet = wire.NewSet(
	user.NewGeneralUserController,
	address.NewGeneralAddressController,
	corporation.NewGeneralCorporationController,
	notification.NewGeneralNotificationController,
	installation.NewGeneralInstallationController,
	news.NewGeneralNewsController,
	payment.NewGeneralPaymentController,
	ticket.NewGeneralTicketController,
	wire.Struct(new(GeneralControllers), "*"),
)

var CustomerControllerProviderSet = wire.NewSet(
	user.NewCustomerUserController,
	installation.NewCustomerInstallationController,
	address.NewCustomerAddressController,
	corporation.NewCustomerCorporationController,
	chat.NewCustomerChatController,
	notification.NewCustomerNotificationController,
	maintenance.NewCustomerMaintenanceController,
	ticket.NewCustomerTicketController,
	report.NewCustomerReportController,
	wire.Struct(new(CustomerControllers), "*"),
)

var CorporationControllerProviderSet = wire.NewSet(
	corporation.NewCorporationCorporationController,
	installation.NewCorporationInstallationController,
	chat.NewCorporationChatController,
	maintenance.NewCorporationMaintenanceController,
	guarantee.NewCorporationGuaranteeController,
	wire.Struct(new(CorporationControllers), "*"),
)

var AdminControllerProviderSet = wire.NewSet(
	ticket.NewAdminTicketController,
	user.NewAdminUserController,
	report.NewAdminReportController,
	news.NewAdminNewsController,
	corporation.NewAdminCorporationController,
	installation.NewAdminInstallationController,
	wire.Struct(new(AdminControllers), "*"),
)

var ControllersProviderSet = wire.NewSet(
	wire.Struct(new(Controllers), "*"),
)

var MiddlewareProviderSet = wire.NewSet(
	middleware.NewAuthMiddleware,
	middleware.NewCorsMiddleware,
	middleware.NewRecovery,
	middleware.NewLocalization,
	middleware.NewRateLimit,
	middleware.NewLoggerMiddleware,
	middleware.NewPrometheusMiddleware,
	middleware.NewWebsocketMiddleware,
	wire.Struct(new(Middlewares), "*"),
)

var SeederProviderSet = wire.NewSet(
	seed.NewAddressSeeder,
	seed.NewNotificationTypeSeeder,
	seed.NewRoleSeeder,
	seed.NewContactTypeSeeder,
	wire.Struct(new(Seeds), "*"),
)

var ConsumerProviderSet = wire.NewSet(
	consumer.NewRegisterConsumer,
	consumer.NewPushConsumer,
	consumer.NewEmailConsumer,
	consumer.NewSendNotificationConsumer,
	wire.Struct(new(Consumers), "*"),
)

func ProvideConstants(container *bootstrap.Config) *bootstrap.Constants {
	return container.Constants
}

func ProvideLoggerConfig(container *bootstrap.Config) *bootstrap.Logger {
	return &container.Env.Logger
}

func ProvideRateLimitConfig(container *bootstrap.Config) *bootstrap.RateLimit {
	return &container.Env.RateLimit
}

func ProvideDBConfig(container *bootstrap.Config) *bootstrap.Database {
	return &container.Env.PrimaryDB
}

func ProvideRDBConfig(container *bootstrap.Config) *bootstrap.Redis {
	return &container.Env.PrimaryRedis
}

func ProvideOTPConfig(container *bootstrap.Config) *bootstrap.OTP {
	return &container.Env.OTP
}

func ProvideSMSGatewayConfig(container *bootstrap.Config) *bootstrap.SMSGateway {
	return &container.Env.SMSGateway
}

func ProvideSMSTemplates(container *bootstrap.Config) *bootstrap.SMSTemplates {
	return &container.Constants.SMSTemplates
}

func ProvideJWTKeysPath(container *bootstrap.Config) *bootstrap.JWTKeysPath {
	return &container.Constants.JWTKeysPath
}

func ProvideEmailTemplates(container *bootstrap.Config) *bootstrap.EmailTemplates {
	return &container.Constants.EmailTemplates
}

func ProvideMetrics(container *bootstrap.Config) *bootstrap.Metrics {
	return &container.Constants.Metrics
}

func ProvidePaginationConfig(container *bootstrap.Config) *bootstrap.Pagination {
	return &container.Env.Pagination
}

func ProvideStorageConfig(container *bootstrap.Config) *bootstrap.S3 {
	return &container.Env.Storage
}

func ProvideWebsocketSetting(container *bootstrap.Config) *bootstrap.WebsocketSetting {
	return &container.Env.WebsocketSetting
}

func ProvideEmailSenderAccount(container *bootstrap.Config) *bootstrap.EmailAccount {
	return &container.Env.EmailSenderAccount
}

func ProvideSuperAdminCredential(container *bootstrap.Config) *bootstrap.AdminCredentials {
	return &container.Env.SuperAdmin
}

func ProvideRabbitMQConfig(container *bootstrap.Config) *bootstrap.RabbitMQ {
	return &container.Env.RabbitMQ
}

func ProvideRabbitMQConstants(container *bootstrap.Config) *bootstrap.RabbitMQConstants {
	return &container.Constants.RabbitMQ
}

var ProviderSet = wire.NewSet(
	DatabaseProviderSet,
	RepositoryProviderSet,
	ServiceProviderSet,
	AdapterProviderSet,
	GeneralControllerProviderSet,
	CustomerControllerProviderSet,
	CorporationControllerProviderSet,
	AdminControllerProviderSet,
	ControllersProviderSet,
	MiddlewareProviderSet,
	SeederProviderSet,
	ConsumerProviderSet,
	ProvideConstants,
	ProvideLoggerConfig,
	ProvideRateLimitConfig,
	ProvideDBConfig,
	ProvideRDBConfig,
	ProvideOTPConfig,
	ProvideSMSGatewayConfig,
	ProvideSMSTemplates,
	ProvideEmailTemplates,
	ProvideJWTKeysPath,
	ProvideMetrics,
	ProvidePaginationConfig,
	ProvideStorageConfig,
	ProvideWebsocketSetting,
	ProvideEmailSenderAccount,
	ProvideSuperAdminCredential,
	ProvideRabbitMQConfig,
	ProvideRabbitMQConstants,
)

type Database struct {
	DB  database.Database
	RDB database.Cache
}

type GeneralControllers struct {
	UserController         *user.GeneralUserController
	AddressController      *address.GeneralAddressController
	CorporationController  *corporation.GeneralCorporationController
	NotificationController *notification.GeneralNotificationController
	InstallationController *installation.GeneralInstallationController
	NewsController         *news.GeneralNewsController
	PaymentController      *payment.GeneralPaymentController
	TicketController       *ticket.GeneralTicketController
}

type CustomerControllers struct {
	UserController         *user.CustomerUserController
	InstallationController *installation.CustomerInstallationController
	AddressController      *address.CustomerAddressController
	CorporationController  *corporation.CustomerCorporationController
	ChatController         *chat.CustomerChatController
	NotificationController *notification.CustomerNotificationController
	MaintenanceController  *maintenance.CustomerMaintenanceController
	TicketController       *ticket.CustomerTicketController
	ReportController       *report.CustomerReportController
}

type CorporationControllers struct {
	CorporationController  *corporation.CorporationCorporationController
	InstallationController *installation.CorporationInstallationController
	ChatController         *chat.CorporationChatController
	MaintenanceController  *maintenance.CorporationMaintenanceController
	GuaranteeController    *guarantee.CorporationGuaranteeController
}

type AdminControllers struct {
	TicketController       *ticket.AdminTicketController
	UserController         *user.AdminUserController
	ReportController       *report.AdminReportController
	NewsController         *news.AdminNewsController
	CorporationController  *corporation.AdminCorporationController
	InstallationController *installation.AdminInstallationController
}

type Controllers struct {
	General     *GeneralControllers
	Customer    *CustomerControllers
	Corporation *CorporationControllers
	Admin       *AdminControllers
}

type Middlewares struct {
	Authentication      *middleware.AuthMiddleware
	CORS                *middleware.CORSMiddleware
	Recovery            *middleware.RecoveryMiddleware
	Localization        *middleware.LocalizationMiddleware
	RateLimit           *middleware.RateLimitMiddleware
	Logger              *middleware.LoggerMiddleware
	Prometheus          *middleware.PrometheusMiddleware
	WebsocketMiddleware *middleware.WebsocketMiddleware
}

type Seeds struct {
	AddressSeeder          *seed.AddressSeeder
	NotificationTypeSeeder *seed.NotificationTypeSeeder
	RoleSeeder             *seed.RoleSeeder
	ContactType            *seed.ContactTypeSeeder
}

type Consumers struct {
	Register     *consumer.RegisterConsumer
	Push         *consumer.PushConsumer
	Email        *consumer.EmailConsumer
	Notification *consumer.SendNotificationConsumer
}

type Application struct {
	Database    *Database
	Controllers *Controllers
	Middlewares *Middlewares
	Seeds       *Seeds
	Consumers   *Consumers
}

func NewApplication(
	database *Database,
	controllers *Controllers,
	middlewares *Middlewares,
	seeds *Seeds,
	consumers *Consumers,
) *Application {
	return &Application{
		Database:    database,
		Controllers: controllers,
		Middlewares: middlewares,
		Seeds:       seeds,
		Consumers:   consumers,
	}
}

func InitializeApplication(container *bootstrap.Config, hub *websocket.Hub) (*Application, error) {
	wire.Build(
		ProviderSet,
		NewApplication,
	)
	return &Application{}, nil
}
