package main
/*
import (
	"SearchEngineV2/searcher/utils"
	"errors"
	"fmt"
	"strings"

	"github.com/wangbin/jiebago"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Word struct {
	ID        int32  `gorm:"primaryKey;autoIncrement:true"`
	SplitWord string `gorm:"not null""`
	UrlID     int32  `gorm:"not null""`
	Loc       int    `gorm:"not null""`
}

type Tokenizer struct {
	seg jiebago.Segmenter
}

//Url
type Url struct {
	ID  int32  `gorm:"primaryKey;autoIncrement:true"`
	Url string `gorm:"not null""`
	Dec string `gorm:"not null""`
}

func main() {
	start := 0
	dsn := "root:123456@tcp(localhost:3306)/searchdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		errors.As(err, "err")
	}
	db.AutoMigrate(&Url{})
	urls := []Url{}
	db.Find(&urls) //将查询数据传入到对象中
	db.AutoMigrate(&Word{})
	for i := 2; i < 8; i++ {
		go threads_insert(urls, i, start+1000*(i-1), db)
	}
	for i := start; i < start+1000; i++ {

		s := urls[i].Dec
		ss := Cut(s)
		fmt.Println("Thread ", 1, " ", urls[i].ID, ss)

		count := 0
		for j := 0; j < len(ss); j++ {
			db.Create(&Word{SplitWord: ss[j], UrlID: urls[i].ID, Loc: count})
			count += len(ss[j])
		}
	}

}
func threads_insert(urls []Url, id int, start int, db *gorm.DB) {
	for i := start; i < start+1000; i++ {

		s := urls[i].Dec
		ss := Cut(s)
		fmt.Println("Thread ", id, " ", urls[i].ID, ss)

		count := 0
		for j := 0; j < len(ss); j++ {
			db.Create(&Word{SplitWord: ss[j], UrlID: urls[i].ID, Loc: count})
			count += len(ss[j])
		}
	}

}

//1662
func Cut(text string) []string {
	tokenizer := &Tokenizer{}
	dictionaryPath := "E:/query/study/Security-Brigade-Rep1/Security-Brigade-Rep1/data/dictionary.txt"
	err := tokenizer.seg.LoadDictionary(dictionaryPath) //不区分大小写
	if err != nil {
		panic(err)
	}
	text = strings.ToLower(text)
	//移除所有的标点符号
	text = utils.RemovePunctuation(text)
	//移除所有的空格
	text = utils.RemoveSpace(text)

	var wordMap = make(map[string]int)

	resultChan := tokenizer.seg.CutForSearch(text, true)
	for {
		w, ok := <-resultChan
		if !ok {
			break
		}
		_, found := wordMap[w]
		if !found {
			//去除重复的词
			wordMap[w] = 1
		}
	}

	var wordsSlice []string
	for k := range wordMap {
		wordsSlice = append(wordsSlice, k)
	}

	return wordsSlice
}
*/