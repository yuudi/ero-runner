package model

type Container struct {
	ID         string `gorm:"primaryKey"`
	UserID     string
	User       User
	CreatedAt  int64 `gorm:"autoCreateTime"`
	StartedAt  int64 `gorm:"autoCreateTime"`
	LastUsedAt int64
}
