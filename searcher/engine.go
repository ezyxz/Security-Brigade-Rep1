package searcher

import (
	model2 "SearchEngineV2/model"
	"SearchEngineV2/searcher/model"
	"errors"
	"fmt"
	"github.com/wangbin/jiebago"
	"gorm.io/gorm"
	"sort"
)

type Engine struct {
	name           string
	seg            jiebago.Segmenter
	dbname         string
	dictionaryPath string
	MySqlDB        *gorm.DB
}

type Word struct {
	ID        int32  `gorm:"primaryKey;autoIncrement:true"`
	SplitWord string `gorm:"not null"`
	UrlID     int32  `gorm:"not null"`
	Loc       int    `gorm:"not null"`
}

type Url struct {
	ID  int32  `gorm:"primaryKey;autoIncrement:true"`
	Url string `gorm:"not null"`
	Dec string `gorm:"not null"`
}

type User struct {
	ID       int64  `gorm:"primaryKey;autoIncrement:true"`
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
}

func (c *Engine) Query(wordsSlices []string) *model.SearchResult {
	var result = &model.SearchResult{
		Words: wordsSlices,
	}

	//db, err := gorm.Open(mysql.Open(c.dbname), &gorm.Config{})
	c.MySqlDB.AutoMigrate(&Word{})
	n := len(wordsSlices)
	temp := []Word{}
	urlId := []int32{}
	var url2info map[int32][]string
	url2info = make(map[int32][]string)
	for i := 0; i < n; i++ {
		c.MySqlDB.Where("split_word = ? ", wordsSlices[i]).Find(&temp)
		for j := 0; j < len(temp); j++ {
			url2info[temp[j].UrlID] = append(url2info[temp[j].UrlID], wordsSlices[i])
			urlId = append(urlId, temp[j].UrlID)
		}
	}
	n = len(urlId)
	c.MySqlDB.AutoMigrate(&Url{})
	urls := []Url{}
	for i := 0; i < n; i++ {
		c.MySqlDB.Find(&urls, urlId)
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

func (c *Engine) InsertUser(u *model2.RegisterUserRequest) bool {
	err := c.MySqlDB.AutoMigrate(&User{})
	if err != nil {
		fmt.Println("数据库表获取失败")
		return false
	}
	if err = c.MySqlDB.Create(&User{Username: u.Username, Password: u.Password}).Error; err != nil {
		fmt.Println("插入新用户失败")
		return false
	}
	return true
}

func (c *Engine) QueryIDAndPwdByName(name string) (bool, int64, string) {
	err := c.MySqlDB.AutoMigrate(&User{})
	if err != nil {
		fmt.Println("数据库表获取失败")
		return false, -1, "服务器错误"
	}
	var queryResult User
	err = c.MySqlDB.Select("id", "password").Where("username=?", name).First(&queryResult).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, -1, "用户不存在"
	} else {
		return true, queryResult.ID, queryResult.Password
	}
}
