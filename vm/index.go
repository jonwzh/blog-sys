package vm

import "github.com/jonwzh/blog-sys/model"

// IndexViewModel struct
type IndexViewModel struct {
	BaseViewModel
	model.User
	Posts []model.Post
}

// IndexViewModelOp struct
type IndexViewModelOp struct{}

// GetVM func
func (IndexViewModelOp) GetVM() IndexViewModel {
	u1 := model.User{Username: "Jonathan"}
	u2 := model.User{Username: "Wang"}

	posts := []model.Post{
		model.Post{User: u1, Body: "Beautiful day in San Francisco!"},
		model.Post{User: u2, Body: "The last trip was so freaking awesome!"},
	}

	v := IndexViewModel{BaseViewModel{Title: "Homepage"}, u1, posts}
	return v
}