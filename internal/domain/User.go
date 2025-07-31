package domain

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"null"`
	Email     string    `json:"email" gorm:"not null;unique;index"`
	Password  string    `json:"password"`
	Role      string    `json:"role" gorm:"default:buyer"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	Status    string    `json:"status"`
	Verified  bool      `json:"verified" gorm:"default:false"`
	ExpiredAt time.Time `json:"expired_at" gorm:"null"`
	CreatedAt time.Time `json:"created_at" gorm:"current_timestamp"`
	UpdatedAt time.Time `json:"updated_at"	gorm:"current_timestamp;onupdate:current_timestamp"`
}
