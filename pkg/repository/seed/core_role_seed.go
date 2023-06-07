package seed

import (
	"errors"
	"gorm.io/gorm"
	"hospital-api/pkg/repository/model"
)

func SeedRole(gormDB *gorm.DB) {
	if gormDB.Migrator().HasTable(&model.CoreRole{}) {
		if err := gormDB.First(&model.CoreRole{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			permissions := []model.CoreRole{
				{
					Name: "ADMIN",
					Permission: []model.CorePermission{
						{ID: 1, Name: "manage-user"},
						{ID: 2, Name: "manage-role"},
					},
				},
			}
			gormDB.Create(&permissions)
		}
	}
}
