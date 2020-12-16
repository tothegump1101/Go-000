package main

import (
	"github.com/tothegump/project/configs"
	"github.com/tothegump/project/internal/model"
)

func main() {
	c := configs.Conf{}
	u := model.User{
		Id:   0,
		Name: "",
		Age:  0,
	}
	_ = c
	_ = u
}
