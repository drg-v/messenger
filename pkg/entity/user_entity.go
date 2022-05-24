package entity

type User struct {
	ID       uint      `gorm:"primaryKey;autoIncrement"`
	Name     string    `gorm:"type:varchar(100);not null"`
	Password string    `gorm:"type:varchar(100);not null"`
	Messages []Message `gorm:"foreignKey:Receiver"`
}
