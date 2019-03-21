package authService

import "gin-blog/models"

type Auth struct {
	UserName string
	Password string
}

func (a *Auth) Check() (bool, error) {
	return models.CheckAuth(a.UserName, a.Password)
}
