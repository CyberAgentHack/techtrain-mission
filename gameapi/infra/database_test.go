package infra

import (
	"os"

	"github.com/task4233/techtrain-mission/log"
)

var (
	logger = log.MyLogger
)

func ExampleNewDB() {
	db, err := NewDB()
	if err != nil {
		logger.Println(err)
		os.Exit(1)
	}
	// 何があっても強い気持ちでCloseする
	// クローズに失敗した場合はそれもしっかりロギングする
	defer func() {
		cerr := db.Close()
		if err != nil {
			logger.Println(cerr)
		}
	}()
}
