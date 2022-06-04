package controller

import (
	"SearchEngineV2/service"
)

var srv *Services

type Services struct {
	Word     *service.Word
	Searcher *service.Search
	Register *service.Register
	Login    *service.Login
}

func NewServices() {
	srv = &Services{
		Word:     service.NewWord(),
		Searcher: service.NewSearch(),
		Register: service.NewRegister(),
		Login:    service.NewLogin(),
	}
}
