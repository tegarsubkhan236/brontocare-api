package seed

import (
	"errors"
	"gorm.io/gorm"
	"hospital-api/pkg/repository/model"
)

func SeedUser(gormDB *gorm.DB) {
	if gormDB.Migrator().HasTable(&model.CoreUser{}) {
		if err := gormDB.First(&model.CoreUser{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			users := []model.CoreUser{
				{
					Name:     "admin",
					Username: "admin",
					Sex:      "male",
					Email:    "admin@admin.admin",
					// admin123
					Password: "$2a$14$1Fo5fY0klE/c1Iog1klCNOI/8DorxfgNqfSD2RQeO4dQkIUVawy6m",
					Status:   1,
					Permission: []model.CorePermission{
						{ID: 1, Name: "manage-user"},
						{ID: 2, Name: "manage-role"},
					},
					Role: []model.CoreRole{
						{ID: 1, Name: "ADMIN"},
					},
				},
			}
			gormDB.Create(&users)
		}
	}
}
