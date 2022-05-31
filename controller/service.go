package controller

import "SearchEngineV2/service"

var srv *Services

type Services struct {
	Word     *service.Word
	Searcher *service.Search
}

func NewServices() {
	srv = &Services{
		Word:     service.NewWord(),
		Searcher: service.NewSearch(),
	}
}
