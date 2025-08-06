package enum

type UserType uint

const (
	UserTypeGuest UserType = iota + 1
	UserTypeCustomer
	UserTypeAdmin
)

func (userType UserType) String() string {
	switch userType {
	case UserTypeGuest:
		return "guest"
	case UserTypeCustomer:
		return "customer"
	case UserTypeAdmin:
		return "admin"
	}
	return ""
}

func GetAllUserTypes() []UserType {
	return []UserType{
		UserTypeGuest,
		UserTypeCustomer,
		UserTypeAdmin,
	}
}
