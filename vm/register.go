package vm

import (
	"log"

	"github.com/jonwzh/blog-sys/model"
)

// RegisterViewModel struct
type RegisterViewModel struct {
	BaseViewModel
	Errs []string
}

// AddError func
func (v *RegisterViewModel) AddError(errs ...string) {
	v.Errs = append(v.Errs, errs...)
}

// RegisterViewModelOp struct
type RegisterViewModelOp struct {}

// GetVM func
func (RegisterViewModelOp) GetVM() RegisterViewModel {
	v := RegisterViewModel{}
	v.SetTitle("Register")
	return v
}

// CheckUserExists : returns true if the username exists in the db
func CheckUserExists(username string) bool {
	_, err := model.GetUserByUsername(username)
	if err != nil {
		log.Println("Cannot find user with username: ", username)
		return true
	}
	return false
}

// AddUser : adds a user in the db
func AddUser(username, passowrd, email string) error {
	return model.AddUser(username, passowrd, email)
}
