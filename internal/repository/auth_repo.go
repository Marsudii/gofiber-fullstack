package repository

type AuthRepo interface {
	Login(email string, password string) (string, error)
}

type authRepo struct {
}

func NewAuthRepo() AuthRepo {
	return &authRepo{}
}

func (a *authRepo) Login(email string, password string) (string, error) {
	var dataALL string
	dataALL = "email: " + email + " password: " + password
	return dataALL, nil
}

func (a *authRepo) Register(email string, password string) (string, error) {
	var dataALL string
	dataALL = "email: " + email + " password: " + password
	return dataALL, nil
}
