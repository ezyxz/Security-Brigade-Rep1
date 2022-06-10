package searcher

import (
	"SearchEngineV2/searcher/words"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Container struct {
	Dir        string                //文件夹
	Engine     *Engine               //引擎
	Debug      bool                  //调试
	Tokenizer  *words.Tokenizer      //分词器
	Shard      int                   //分片
	Timeout    int64                 //超时关闭数
}

func (c *Container) Init() error {
	c.Engine = c.NewEngine()
	return nil
}

func (c *Container) NewEngine() *Engine {
	var engine = &Engine{
		name: "my searcher",
		//dbname:         "root:root@tcp(localhost:3306)/searchdb2?charset=utf8mb4&parseTime=True&loc=Local",
		dbname:         "root:123456@tcp(localhost:3306)/searchdb2?charset=utf8mb4&parseTime=True&loc=Local",
		dictionaryPath: "./searcher/words/data/dictionary.txt",
	}
	var err error
	engine.MySqlDB, err = gorm.Open(mysql.Open(engine.dbname), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败")
		panic(c)
		return nil
	}
	return engine
}
