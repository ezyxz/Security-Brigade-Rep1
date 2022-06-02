package service

import (
	"SearchEngineV2/global"
	"SearchEngineV2/register"
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
func (w *Register) CheckUserExisted(user *register.LogUser) int32 {
	return w.Container.Register.CheckUser(user)
}

func (w *Register) RegisterNewUser(user *register.User) int {
	return w.Container.Register.RegisterUser(user)
}
