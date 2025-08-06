package enum

type RoleName uint

const (
	SuperAdmin RoleName = iota + 1
	Customer
	Technician
	SupportAgent
	ContentManager
	Moderator
)

var rolePermissions = map[RoleName][]PermissionType{
	SuperAdmin: {
		PermissionAll,
	},
	Customer: {},
	ContentManager: {
		NewsViewAll, NewsCreate, NewsEdit, NewsDelete,
	},
}

func (role RoleName) Permissions() []PermissionType {
	if permissions, ok := rolePermissions[role]; ok {
		return permissions
	}
	return nil
}

func (role RoleName) String() string {
	switch role {
	case SuperAdmin:
		return "سوپر ادمین"
	case Customer:
		return "مشتری"
	case Technician:
		return "تکنسین"
	case SupportAgent:
		return "پشتیبان"
	case ContentManager:
		return "مدیر محتوا"
	case Moderator:
		return "ناظر"
	}
	return "unknown"
}

func GetAllRoleNames() []RoleName {
	return []RoleName{
		SuperAdmin,
		Customer,
		Technician,
		SupportAgent,
		ContentManager,
		Moderator,
	}
}
