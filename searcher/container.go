package searcher

import (
	"SearchEngineV2/searcher/words"
)

type Container struct {
	Dir       string           //文件夹
	Engine    *Engine          //引擎
	Debug     bool             //调试
	Tokenizer *words.Tokenizer //分词器
	Shard     int              //分片
	Timeout   int64            //超时关闭数据库
}

func (c *Container) Init() error {

	c.Engine = c.NewEngine()
	return nil
}

func (c *Container) NewEngine() *Engine {
	var engine = &Engine{
		name:           "my searcher",
		dbname:         "root:root@tcp(localhost:3306)/searchdb2?charset=utf8mb4&parseTime=True&loc=Local",
		dictionaryPath: "./searcher/words/data/dictionary.txt",
	}
	return engine
}
