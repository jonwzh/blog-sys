package vm

import (
	"log"

	"github.com/jonwzh/blog-sys/model"
)

// LoginViewModel struct
type LoginViewModel struct {
	BaseViewModel
	Errs []string
}

// AddError func
func (v *LoginViewModel) AddError(errs ...string) {
	v.Errs = append(v.Errs, errs...)
}

// LoginViewModelOp struct
type LoginViewModelOp struct {}

// GetVM func
func (LoginViewModelOp) GetVM() LoginViewModel {
	v := LoginViewModel{}
	v.SetTitle("Login")
	return v
}

// CheckLogin : returns true if the login details match
func CheckLogin(username, password string) bool {
	user, err := model.GetUserByUsername(username)
	if err != nil {
		log.Println("Failed to find username: ", username)
		log.Println("Error: ", err.Error())
		return false;
	}
	return user.CheckPassword(password)
}