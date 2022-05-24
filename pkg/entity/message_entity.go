package entity

type Message struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Sender   uint
	Receiver uint
	Text     string `gorm:"type:varchar(500);not null"`
}
