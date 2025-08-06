package enum

type PermissionType uint
type PermissionCategory uint

const (
	PermissionAll PermissionType = iota + 1
	PermissionGeneral

	// Admin Role Permissions
	// User Management
	UserViewAll
	UserBanUnban
	UserChangeRole
	UserViewRoles
	UserManageRolePermissions
	UserRemoveRole
	UserCreateRole

	// Ticket Management
	TicketViewAll
	TicketRespond
	TicketClose
	TicketComment

	// Report Management
	ReportViewAll
	ReportRespond

	// News Management
	NewsViewAll
	NewsCreate
	NewsEdit
	NewsDelete

	// Maintenance Management
	MaintenanceViewAll
	MaintenanceAcceptRequest
	MaintenanceCreateRecord
	MaintenanceUpdateRecord

	// Guarantee Management
	GuaranteeViewAll
	GuaranteeCreate
	GuaranteeArchiveUnarchive

	// Profile Management
	ProfileViewPrivate
	ProfileUpdate
)

const (
	CategoryGeneral PermissionCategory = iota + 1
	CategoryUser
	CategoryTicket
	CategoryReport
	CategoryNews
	CategoryPanel
	CategoryMaintenance
	CategoryGuarantee
	CategoryProfile
)

var permissionNames = map[PermissionType]string{
	PermissionAll:     "general.all",
	PermissionGeneral: "general.general",

	// User Management
	UserViewAll:               "user.viewAll",
	UserBanUnban:              "user.banUnban",
	UserChangeRole:            "user.changeRole",
	UserViewRoles:             "user.viewRoles",
	UserManageRolePermissions: "user.manageRolePermissions",
	UserRemoveRole:            "user.removeRole",
	UserCreateRole:            "user.createRole",

	// Ticket Management
	TicketViewAll: "ticket.viewAll",
	TicketRespond: "ticket.respond",
	TicketClose:   "ticket.close",
	TicketComment: "ticket.comment",

	// Report Management
	ReportViewAll: "report.viewAll",
	ReportRespond: "report.respond",

	// News Management
	NewsViewAll: "news.viewAll",
	NewsCreate:  "news.create",
	NewsEdit:    "news.edit",
	NewsDelete:  "news.delete",

	// Panel Management
	PanelViewAll: "panel.viewAll",
	PanelCreate:  "panel.create",

	// Maintenance Management
	MaintenanceViewAll:       "maintenance.viewAll",
	MaintenanceAcceptRequest: "maintenance.acceptRequest",
	MaintenanceCreateRecord:  "maintenance.createRecord",
	MaintenanceUpdateRecord:  "maintenance.updateRecord",

	// Guarantee Management
	GuaranteeViewAll:          "guarantee.viewAll",
	GuaranteeCreate:           "guarantee.create",
	GuaranteeArchiveUnarchive: "guarantee.archiveUnarchive",

	// Profile Management
	ProfileViewPrivate: "profile.viewPrivate",
	ProfileUpdate:      "profile.update",
}

var permissionDescriptions = map[PermissionType]string{
	PermissionAll:     "دسترسی کامل به سیستم",
	PermissionGeneral: "دسترسی عمومی",

	// User Management
	UserViewAll:               "مشاهده لیست کاربران وب‌سایت",
	UserBanUnban:              "مسدود/رفع مسدودیت کاربران",
	UserChangeRole:            "تغییر نقش کاربر",
	UserViewRoles:             "مشاهده لیست نقش‌ها و مجوزها",
	UserManageRolePermissions: "تغییر مجوزهای نقش",
	UserRemoveRole:            "حذف نقش",
	UserCreateRole:            "ایجاد نقش جدید",

	// Ticket Management
	TicketViewAll: "مشاهده لیست تیکت‌ها",
	TicketRespond: "پاسخ به تیکت",
	TicketClose:   "بستن تیکت",
	TicketComment: "افزودن نظر به تیکت",

	// Report Management
	ReportViewAll: "مشاهده لیست گزارش‌ها",
	ReportRespond: "پاسخ به گزارش",

	// News Management
	NewsViewAll: "مشاهده لیست اخبار",
	NewsCreate:  "ایجاد خبر جدید",
	NewsEdit:    "ویرایش خبر",
	NewsDelete:  "حذف خبر",

	// Panel Management
	PanelViewAll: "مشاهده لیست پنل‌ها",
	PanelCreate:  "ایجاد پنل جدید",          "لغو پیشنهاد",

	// Maintenance Management
	MaintenanceViewAll:       "مشاهده لیست درخواست‌های تعمیر",
	MaintenanceAcceptRequest: "پذیرش/لغو درخواست تعمیر",
	MaintenanceCreateRecord:  "ایجاد سابقه تعمیر",
	MaintenanceUpdateRecord:  "به‌روزرسانی سابقه تعمیر",

	// Guarantee Management
	GuaranteeViewAll:          "مشاهده لیست ضمانت‌ها",
	GuaranteeCreate:           "ایجاد ضمانت جدید",
	GuaranteeArchiveUnarchive: "آرشیو/رفع آرشیو ضمانت",

	// Profile Management
	ProfileViewPrivate: "مشاهده اطلاعات خصوصی پروفایل",
	ProfileUpdate:      "به‌روزرسانی پروفایل",
}

var permissionCategories = map[PermissionType]PermissionCategory{
	PermissionAll:     CategoryGeneral,
	PermissionGeneral: CategoryGeneral,

	// User Management
	UserViewAll:               CategoryUser,
	UserBanUnban:              CategoryUser,
	UserChangeRole:            CategoryUser,
	UserViewRoles:             CategoryUser,
	UserManageRolePermissions: CategoryUser,
	UserRemoveRole:            CategoryUser,
	UserCreateRole:            CategoryUser,

	// Ticket Management
	TicketViewAll: CategoryTicket,
	TicketRespond: CategoryTicket,
	TicketClose:   CategoryTicket,
	TicketComment: CategoryTicket,

	// Report Management
	ReportViewAll: CategoryReport,
	ReportRespond: CategoryReport,

	// News Management
	NewsViewAll: CategoryNews,
	NewsCreate:  CategoryNews,
	NewsEdit:    CategoryNews,
	NewsDelete:  CategoryNews,

	// Panel Management
	PanelViewAll: CategoryPanel,
	PanelCreate:  CategoryPanel,

	// Maintenance Management
	MaintenanceViewAll:       CategoryMaintenance,
	MaintenanceAcceptRequest: CategoryMaintenance,
	MaintenanceCreateRecord:  CategoryMaintenance,
	MaintenanceUpdateRecord:  CategoryMaintenance,

	// Guarantee Management
	GuaranteeViewAll:          CategoryGuarantee,
	GuaranteeCreate:           CategoryGuarantee,
	GuaranteeArchiveUnarchive: CategoryGuarantee,

	// Profile Management
	ProfileViewPrivate: CategoryProfile,
	ProfileUpdate:      CategoryProfile,
}

func (perm PermissionType) String() string {
	if description, ok := permissionNames[perm]; ok {
		return description
	}
	return ""
}

func (perm PermissionType) Description() string {
	if description, ok := permissionDescriptions[perm]; ok {
		return description
	}
	return ""
}

func (perm PermissionType) Category() PermissionCategory {
	if category, ok := permissionCategories[perm]; ok {
		return category
	}
	return CategoryGeneral
}

func (category PermissionCategory) String() string {
	switch category {
	case CategoryGeneral:
		return "عمومی"
	case CategoryUser:
		return "مدیریت کاربران"
	case CategoryTicket:
		return "مدیریت تیکت"
	case CategoryReport:
		return "مدیریت گزارشات"
	case CategoryNews:
		return "مدیریت اخبار"
	case CategoryPanel:
		return "مدیریت پنل‌ها"
	case CategoryMaintenance:
		return "مدیریت تعمیرات"
	case CategoryGuarantee:
		return "مدیریت ضمانت‌ها"
	case CategoryProfile:
		return "مدیریت پروفایل"
	}
	return "unknown"
}

func GetAllPermissionTypes() []PermissionType {
	return []PermissionType{
		PermissionAll, PermissionGeneral,

		// User Management
		UserViewAll, UserBanUnban, UserChangeRole, UserViewRoles,
		UserManageRolePermissions, UserRemoveRole, UserCreateRole,

		// Ticket Management
		TicketViewAll, TicketRespond, TicketClose, TicketComment,

		// Report Management
		ReportViewAll, ReportRespond,

		// News Management
		NewsViewAll, NewsCreate, NewsEdit, NewsDelete,

		// Panel Management
		PanelViewAll, PanelCreate,

		// Maintenance Management
		MaintenanceViewAll, MaintenanceAcceptRequest, MaintenanceCreateRecord, MaintenanceUpdateRecord,

		// Guarantee Management
		GuaranteeViewAll, GuaranteeCreate, GuaranteeArchiveUnarchive,

		// Profile Management
		ProfileViewPrivate, ProfileUpdate,
	}
}