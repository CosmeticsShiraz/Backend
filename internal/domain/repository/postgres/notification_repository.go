package postgres

import (
	"github.com/CosmeticsShiraz/Backend/internal/domain/entity"
	"github.com/CosmeticsShiraz/Backend/internal/domain/enum"
	"github.com/CosmeticsShiraz/Backend/internal/infrastructure/database"
)

type NotificationRepository interface {
	GetNotificationByID(db database.Database, notificationID uint) (*entity.Notification, error)
	GetNotificationsByTypesAndUserID(db database.Database, userID uint, types []uint, opts ...QueryModifier) ([]*entity.Notification, error)
	GetNotificationSettingByID(db database.Database, settingID uint) (*entity.NotificationSetting, error)
	GetNotificationSettingByUserAndType(db database.Database, userID, typeID uint) (*entity.NotificationSetting, error)
	GetNotificationSettingByUserID(db database.Database, userID uint) ([]*entity.NotificationSetting, error)
	GetNotificationTypeByID(db database.Database, typeID uint) (*entity.NotificationType, error)
	GetNotificationTypes(db database.Database) ([]*entity.NotificationType, error)
	GetNotificationTypeByName(db database.Database, name enum.NotificationType) (*entity.NotificationType, error)
	CreateNotification(db database.Database, notification *entity.Notification) error
	UpdateNotification(db database.Database, notification *entity.Notification) error
	CreateNotificationSetting(db database.Database, setting *entity.NotificationSetting) error
	UpdateNotificationSetting(db database.Database, setting *entity.NotificationSetting) error
	CreateNotificationType(db database.Database, notificationType *entity.NotificationType) error
}
