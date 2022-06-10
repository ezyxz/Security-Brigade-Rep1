package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wangbin/jiebago"
	"github.com/yanyiwu/gojieba"
	"log"
	"os"
	"path"
	"time"
)

const (
	user           = "root"
	password       = "123456"
	ip             = "localhost:"
	mysqlPort      = "3306"
	redisPort      = "6379"
	db_name        = "wukong"
	stopWordPath   = "./data/stop_words.utf8"
	dataSourcePath = "root:123456@tcp(localhost:3306)/wukong?charset=utf8"
)

type Kv struct {
	Url     string
	Caption string
}

type Tokenizer struct {
	seg jiebago.Segmenter
}

var tokenizer = &Tokenizer{}
var dictionaryPath = "./searcher/words/data/dictionary.txt"

//var jieba = gojieba.NewJieba("E:\\bytedance\\Security-Brigade-Rep1\\data\\jieba.dict.utf8", "E:\\bytedance\\Security-Brigade-Rep1\\data\\hmm_model.utf8")

func main() {
	start := time.Now()
	// 数据库路径
	dsn := dataSourcePath
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// 打开数据库
	db, err := sql.Open("mysql", dsn)

	if err = db.Ping(); err != nil {
		// do something about db error
		fmt.Printf("error")
	}

	// 加载停用词，并将停用词放在map中，之后分词时进行stop word
	stopWords, _ := readFile(stopWordPath)
	//fmt.Println(stopWords)
	type void struct{}
	var member void
	set := make(map[string]void)
	for i := 0; i < len(stopWords); i++ {
		set[stopWords[i]] = member
	}

	// 连接redis数据库
	rdb := redis.NewClient(&redis.Options{
		Addr:     ip + redisPort,
		Password: "",
		DB:       1,
	})
	ctx := context.Background()

	// 设定字典的相关路径
	dictDir := path.Join("E:\\bytedance\\Security-Brigade-Rep1", "data")
	jiebaPath := path.Join(dictDir, "jieba.dict.utf8")
	hmmPath := path.Join(dictDir, "hmm_model.utf8")
	userPath := path.Join(dictDir, "user.dict.utf8")
	idfPath := path.Join(dictDir, "idf.utf8")
	stopPath := path.Join(dictDir, "stop_words.utf8")
	jieba := gojieba.NewJieba(jiebaPath, hmmPath, userPath, idfPath, stopPath)
	var count_rows int = 0
	var temp_string string

	// 获取数据库行数的总大小
	rows, err := db.Query("SELECT COUNT(*) FROM kv")
	if err != nil {
		log.Fatal(err)
	}
	rows.Next()
	rows.Scan(&count_rows)
	fmt.Println(count_rows)

	// 分成4次分别进行分词
	offset := count_rows / 4
	//offset:=4
	//jh()
	var temp_participle []string
	cnt := 0.0
	for i := 0; i <= 4; i++ {
		// 获取caption Note: 通过limit(i*offset, offset)获取部分的row
		row, _ := db.Query("select caption from kv limit ?, ?", i*offset, offset)
		for row.Next() {
			err := row.Scan(&temp_string)

			if err != nil {
				return
			}
			//temp_participle := Cut(temp_string)
			// 进行搜索引擎模式分词
			temp_participle = jieba.CutForSearch(temp_string, true)

			count := make(map[string]int)
			//var virtualParticiple []string
			// 遍历分词结果，统计词频，同时进行stop word
			for i := 0; i < len(temp_participle); i++ {
				_, ok := set[temp_participle[i]]
				if !ok {
					count[temp_participle[i]]++
				}
			}
			// 将停用词放进redis中，以有序set存储，包含了对应的id，和id中某词出现的频率
			for k, v := range count {
				rdb.ZAdd(ctx, k, &redis.Z{Score: cnt, Member: v})
			}
			cnt++
		}
	}
	fmt.Printf("cost=[%s]\n", time.Since(start))
}

func readFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf := bufio.NewScanner(file)
	var lineStr []string
	for {
		if !buf.Scan() {
			break
		}
		line := buf.Text()
		lineStr = append(lineStr, line)
	}
	return lineStr, nil
}

//func Cut(text string) []string {
//	text = strings.ToLower(text)
//	//移除所有的标点符号
//	text = utils.RemovePunctuation(text)
//	//移除所有的空格
//	text = utils.RemoveSpace(text)
//
//	var wordMap = make(map[string]int)
//
//	resultChan := tokenizer.seg.CutForSearch(text, true)
//	for {
//		w, ok := <-resultChan
//		if !ok {
//			break
//		}
//		_, found := wordMap[w]
//		if !found {
//			//去除重复的词
//			wordMap[w] = 1
//		}
//	}
//
//	var wordsSlice []string
//	for k := range wordMap {
//		wordsSlice = append(wordsSlice, k)
//	}
//
//	return wordsSlice
//}
