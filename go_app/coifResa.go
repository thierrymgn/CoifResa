package coifResa

// import "embed"

// //go:embed templates/*

// var EmbedTemplates embed.FS

type UserType string

const (
	Admin      UserType = "admin"
	Client     UserType = "client"
	SalonOwner UserType = "salon_owner"
)

type UserItem struct {
	ID       int64    `json:"id"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Email    string   `json:"email"`
	UserType UserType `json:"user_type"`
}

type UserStoreInterface interface {
	CreateUser(user *UserItem) error
	GetUser(id int64) (*UserItem, error)
	GetUserByUsername(username string) (*UserItem, error)
	GetUserByEmail(email string) (*UserItem, error)
	UpdateUser(user *UserItem) error
	DeleteUser(id int64) error
}

type SalonItem struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Cp          string `json:"cp"`
	Description string `json:"description"`
	UserId      int64  `json:"user_id"`
}

type SalonStoreInterface interface {
	CreateSalon(salon *SalonItem) error
}
