package testpkg

import (
	"fmt"
	"testing"
	"zhu/myrest/model"
)

func TestGetAllUser(t *testing.T) {
	model.InitDb()
	users, err := model.GetAllUsers()
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		fmt.Println(user)
	}
}
