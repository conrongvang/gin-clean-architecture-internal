package models

import "clean_architecture/internal/models"

type User struct {
	models.BaseModel
	FullName string
	NickName string
	BirthDay int64
}
