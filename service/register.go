package service

import (
	"SearchEngineV2/global"
	"SearchEngineV2/model"
	"SearchEngineV2/searcher"
)

type Register struct {
	Container *searcher.Container
}

func NewRegister() *Register {
	return &Register{
		Container: global.Container,
	}
}

//func (w *Register) CheckUserExisted(u *controller.RegisterUserData) int32 {
//
//}

func (w *Register) RegisterUser(request *model.RegisterUserRequest) bool {
	return w.Container.Engine.InsertUser(request)
}
