package models

//struct for author database schema
type AuthorDetail struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Name string
	Address string
	Phone string
}