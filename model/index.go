package model

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
