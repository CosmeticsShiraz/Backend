package enum

type NotificationType uint

const (
	ChatNotificationType NotificationType = iota + 1
	PanelReportCreated
)

var notificationDescriptions = map[NotificationType]string{
	ChatNotificationType:        "پیام جدید",
	PanelReportCreated:          "گزارش جدید از پنل",
}

var supportsEmail = map[NotificationType]bool{
	ChatNotificationType:        false,
	PanelReportCreated:          true,
}

var supportsPush = map[NotificationType]bool{
	ChatNotificationType:        true,
	PanelReportCreated:          true,
}

var notificationEmailTemplate = map[NotificationType]string{
	ChatNotificationType:        "",
	PanelReportCreated:          "/panel_report/fa.html",
}

func (notificationType NotificationType) String() string {
	switch notificationType {
	case ChatNotificationType:
		return "پیام جدید"
	case PanelReportCreated:
		return "گزارشات جدید از پنل"
	}
	return ""
}

func (notificationType NotificationType) Description() string {
	if description, ok := notificationDescriptions[notificationType]; ok {
		return description
	}
	return ""
}

func (notificationType NotificationType) EmailTemplatePath() string {
	if template, ok := notificationEmailTemplate[notificationType]; ok {
		return template
	}
	return ""
}

func (notificationType NotificationType) SupportsEmail() bool {
	if support, ok := supportsEmail[notificationType]; ok {
		return support
	}
	return false
}

func (notificationType NotificationType) SupportsPush() bool {
	if support, ok := supportsPush[notificationType]; ok {
		return support
	}
	return true
}

func GetAllNotificationTypes() []NotificationType {
	return []NotificationType{
		ChatNotificationType,
		PanelReportCreated,
	}
}
