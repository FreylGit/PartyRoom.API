package constants

type Role struct {
	Name      string
	UpperName string
	LowerName string
}

var UserRole Role
var ModeratorRole Role
var AdminRole Role

func init() {
	UserRole = Role{
		Name:      "User",
		UpperName: "USER",
		LowerName: "user",
	}
	ModeratorRole = Role{
		Name:      "Moderator",
		UpperName: "MODERATOR",
		LowerName: "moderator",
	}

	AdminRole = Role{
		Name:      "Admin",
		UpperName: "ADMIN",
		LowerName: "admin",
	}
}
