package coifResa

import "time"

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
	PostalCode  string `json:"postal_code"`
	Description string `json:"description"`
	UserId      int64  `json:"user_id"`
}

type SalonStoreInterface interface {
	CreateSalon(salon *SalonItem) error
	GetSalon(id int64) (*SalonItem, error)
	GetSalonsByUserId(userId int64) ([]*SalonItem, error)
	UpdateSalon(salon *SalonItem) error
	DeleteSalon(id int64) error
}

type HairdresserItem struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	SalonId int64  `json:"salon_id"`
}

type HairdresserStoreInterface interface {
	CreateHairdresser(hairdresser *HairdresserItem) error
	GetHairdresser(id int64) (*HairdresserItem, error)
	GetHairdressersBySalonId(salonId int64) ([]*HairdresserItem, error)
	UpdateHairdresser(hairdresser *HairdresserItem) error
	DeleteHairdresser(id int64) error
}

type SlotItem struct {
	ID            int64     `json:"id"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	HairdresserId int64     `json:"hairdresser_id"`
}

type SlotStoreInterface interface {
	CreateSlot(slot *SlotItem) error
	GetSlot(id int64) (*SlotItem, error)
	GetSlotsByHairdresserId(hairdresserId int64) ([]*SlotItem, error)
	UpdateSlot(slot *SlotItem) error
	DeleteSlot(id int64) error
}

type ReservationItem struct {
	ID     int64    `json:"id"`
	UserId int64    `json:"user_id"`
	SlotId int64    `json:"slot_id"`
	Slot   SlotItem `json:"slot"`
}

type ReservationStoreInterface interface {
	CreateReservation(reservation *ReservationItem) error
	GetReservation(id int64) (*ReservationItem, error)
	GetReservationsByUserId(userId int64) ([]*ReservationItem, error)
	DeleteReservation(id int64) error
}
