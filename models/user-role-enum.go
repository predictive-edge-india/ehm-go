package models

type Role struct {
	Number int16
	Text   string
}

type userRoleRegistry struct {
	SuperAdministrator    Role
	CustomerAdministrator Role
	CustomerManager       Role
}

var UserRoleEnum = newUserRoleRegistry()

func newUserRoleRegistry() *userRoleRegistry {
	return &userRoleRegistry{
		SuperAdministrator: Role{
			Number: 1,
			Text:   "Super Administrator",
		},
		CustomerAdministrator: Role{
			Number: 2,
			Text:   "Customer Administrator",
		},
		CustomerManager: Role{
			Number: 3,
			Text:   "Customer Manager",
		},
	}
}

func (r *userRoleRegistry) CustomerRoles() [][]interface{} {
	return [][]interface{}{
		{r.CustomerAdministrator.Number, r.CustomerAdministrator.Text},
		{r.CustomerManager.Number, r.CustomerManager.Text},
	}
}

func (r *userRoleRegistry) AppRoles() [][]interface{} {
	appRoles := [][]interface{}{
		{r.SuperAdministrator.Number, r.SuperAdministrator.Text},
	}
	return append(appRoles, r.CustomerRoles()...)
}
