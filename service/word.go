package service

import (
	"SearchEngineV2/global"
	"SearchEngineV2/searcher"
)

type Word struct {
	Container *searcher.Container
}

func NewWord() *Word {
	return &Word{
		Container: global.Container,
	}
}

// WordCut 分词
func (w *Word) WordCut(keyword string) []string {
	return w.Container.Tokenizer.Cut(keyword)
}
