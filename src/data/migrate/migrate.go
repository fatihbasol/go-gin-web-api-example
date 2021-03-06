package migrate

import (
	"github.com/fatihbasol/GoGinExample/src/model"
	"gorm.io/gorm"
)

func MigrateModels(g *gorm.DB) {
	err := g.AutoMigrate(&model.User{}, &model.Phone{})
	if err != nil {
		panic("migrate error")
	}
	initSeedData(g)
}

func initSeedData(db *gorm.DB) {

	hasTable := db.Migrator().HasTable(&model.User{})
	if hasTable {
		var rowCount int64
		db.Model(&model.User{}).Count(&rowCount)

		if rowCount == 0 {
			users := []model.User{
				{
					Name: "robert", Email: "robert@robert.com",
					Phone: &model.Phone{
						Number: "5457778899", CountryCode: "+90",
					},
				},
				{
					Name: "zoe", Email: "zoe@zoe.com",
					Phone: &model.Phone{
						Number: "5454445566", CountryCode: "+90",
					}},
				{
					Name: "anna", Email: "anna@anna.com",
					Phone: &model.Phone{
						Number: "5451112233", CountryCode: "+90",
					}},
			}
			db.Create(&users)
		}

	}
}
