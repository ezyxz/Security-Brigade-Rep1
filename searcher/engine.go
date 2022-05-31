package searcher

import (
	"SearchEngineV2/searcher/model"
	"errors"
	"github.com/wangbin/jiebago"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sort"
)

type Engine struct {
	name           string
	seg            jiebago.Segmenter
	dbname         string
	dictionaryPath string
}
type Word struct {
	ID        int32  `gorm:"primaryKey;autoIncrement:true"`
	SplitWord string `gorm:"not null""`
	UrlID     int32  `gorm:"not null""`
	Loc       int    `gorm:"not null""`
}

type Url struct {
	ID  int32  `gorm:"primaryKey;autoIncrement:true"`
	Url string `gorm:"not null""`
	Dec string `gorm:"not null""`
}

func (c *Engine) Query(wordsSlices []string) *model.SearchResult {
	var result = &model.SearchResult{
		Words: wordsSlices,
	}

	db, err := gorm.Open(mysql.Open(c.dbname), &gorm.Config{})
	if err != nil {
		errors.As(err, "err")
	}
	db.AutoMigrate(&Word{})
	n := len(wordsSlices)
	temp := []Word{}
	urlId := []int32{}
	var url2info map[int32][]string
	url2info = make(map[int32][]string)
	for i := 0; i < n; i++ {
		db.Where("split_word = ? ", wordsSlices[i]).Find(&temp)
		for j := 0; j < len(temp); j++ {
			url2info[temp[j].UrlID] = append(url2info[temp[j].UrlID], wordsSlices[i])
			urlId = append(urlId, temp[j].UrlID)
		}
	}
	n = len(urlId)
	db.AutoMigrate(&Url{})
	urls := []Url{}
	for i := 0; i < n; i++ {
		db.Find(&urls, urlId)
	}
	for i := 0; i < len(urls); i++ {
		temp := model.ResponseDoc{
			Url:          urls[i].Url,
			OriginalText: urls[i].Dec,
			Keys:         url2info[urls[i].ID],
			Score:        len(url2info[urls[i].ID]),
		}
		result.Documents = append(result.Documents, temp)
	}
	result.Total = len(urls)
	sort.Slice(result.Documents, func(i, j int) bool {
		return result.Documents[i].Score > result.Documents[j].Score
	})

	return result
}
