package user

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	newUuid := uuid.NewV4()
	return scope.SetColumn("Id", newUuid.String())
}
