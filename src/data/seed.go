package data

import (
	"github.com/fatihbasol/GoGinExample/src/model"
)

type Repo struct {
	Users  []model.User
	Phones []model.Phone
}

func InitSeedData() (repo *Repo) {
	repo = &Repo{
		Users: []model.User{
			{
				Id:            1,
				Name:          "robert",
				Email:         "robert@robert.com",
				PhoneNumberId: 1,
			},
			{
				Id:            2,
				Name:          "zoe",
				Email:         "zoe@zoe.com",
				PhoneNumberId: 2,
			},
			{
				Id:            3,
				Name:          "anna",
				Email:         "anna@anna.com",
				PhoneNumberId: 3,
			},
		},
		Phones: []model.Phone{
			{Id: 1, CountryCode: "+90", Number: "5457778899"},
			{Id: 2, CountryCode: "+90", Number: "5454445566"},
			{Id: 3, CountryCode: "+90", Number: "5451112233"},
		},
	}
	return
}
