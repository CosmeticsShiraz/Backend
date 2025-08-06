package seed

import (
	"github.com/CosmeticsShiraz/Backend/internal/domain/entity"
	"github.com/CosmeticsShiraz/Backend/internal/domain/enum"
	repository "github.com/CosmeticsShiraz/Backend/internal/domain/repository/postgres"
	"github.com/CosmeticsShiraz/Backend/internal/infrastructure/database"
)

type NotificationTypeSeeder struct {
	userRepository         repository.UserRepository
	notificationRepository repository.NotificationRepository
	db                     database.Database
}

func NewNotificationTypeSeeder(
	userRepository repository.UserRepository,
	notificationRepository repository.NotificationRepository,
	db database.Database,
) *NotificationTypeSeeder {
	return &NotificationTypeSeeder{
		userRepository:         userRepository,
		notificationRepository: notificationRepository,
		db:                     db,
	}
}

func (seeder *NotificationTypeSeeder) SeedNotificationTypes() {
	for _, notification := range enum.GetAllNotificationTypes() {
		notificationType, err := seeder.notificationRepository.GetNotificationTypeByName(seeder.db, notification)
		if err != nil {
			panic(err)
		}
		if notificationType == nil {
			notificationType := &entity.NotificationType{
				Name:              notification,
				Description:       notification.Description(),
				SupportsEmail:     notification.SupportsEmail(),
				SupportsPush:      notification.SupportsPush(),
				EmailTemplatePath: notification.EmailTemplatePath(),
			}
			err := seeder.notificationRepository.CreateNotificationType(seeder.db, notificationType)
			if err != nil {
				panic(err)
			}
			seeder.syncNewNotificationTypesForUsers(notificationType)
		}
	}
}

func (seeder *NotificationTypeSeeder) syncNewNotificationTypesForUsers(newType *entity.NotificationType) error {
	users, err := seeder.userRepository.FindUsers(seeder.db)
	if err != nil {
		return err
	}
	for _, user := range users {
		setting := &entity.NotificationSetting{
			UserID:         user.ID,
			TypeID:         newType.ID,
			IsEmailEnabled: newType.SupportsEmail,
			IsPushEnabled:  newType.SupportsPush,
		}
		err := seeder.notificationRepository.CreateNotificationSetting(seeder.db, setting)
		if err != nil {
			panic(err)
		}
	}
	return nil
}
