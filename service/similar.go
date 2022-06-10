package service

import (
	"SearchEngineV2/global"
	"SearchEngineV2/searcher"
	"SearchEngineV2/searcher/similarity"
)

type Similar struct {
	Container *searcher.Container
}

func NewSimilar() *Similar {
	return &Similar{
		Container: global.Container,
	}
}

// similar 相似
func (w *Similar) WordSimilar(keyword string) []string {

	return similarity.Similar(w.Container.Tokenizer.Cut(keyword))
}
