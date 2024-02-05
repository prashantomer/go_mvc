package models

import "gorm.io/gorm"

type Post struct {
  gorm.Model
  Title string
	Body 	string
}

// // equals
// type Post struct {
//   ID        uint           `gorm:"primaryKey"`
//   CreatedAt time.Time
//   UpdatedAt time.Time
//   DeletedAt gorm.DeletedAt `gorm:"index"`
//   Title string
//   Body string
// }