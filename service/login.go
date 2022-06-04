package service

import (
	"SearchEngineV2/global"
	"SearchEngineV2/middleware"
	"SearchEngineV2/model"
	"SearchEngineV2/searcher"
)

type Login struct {
	Container *searcher.Container
}

func NewLogin() *Login {
	return &Login{
		Container: global.Container,
	}
}

func (w *Login) LoginUser(request *model.LoginRequest) (bool, string) {
	isOK, ID, password := w.Container.Engine.QueryIDAndPwdByName(request.Username)

	if !isOK {
		return false, password
	} else if password != request.Password {
		return false, "密码错误"
	}

	token := middleware.CreateToken(ID)
	return true, token
}
