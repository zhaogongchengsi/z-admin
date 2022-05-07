package model

import "time"

type PublicModel struct {
	Id        string    `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
}

func CreateTables() error {
	err := CreateUserTable()
	if err != nil {
		return err
	}
	err = CreateMenuTable()
	if err != nil {
		return err
	}
	if err = CreateRoleTable(); err != nil {
		return err
	}
	return nil
}
